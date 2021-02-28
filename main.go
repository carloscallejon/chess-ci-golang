package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	board "github.com/carloscallejon/chess-ci-golang/board"
	_ "github.com/heroku/x/hmetrics/onload"
)

type getMoveRequest struct {
	Fen   string `json:"fen"`
	Depth int8   `json:"depth"`
}

type makeMoveRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
	Fen  string `json:"fen"`
}

func main() {
	/*
		// Uncomment for performance profiling
		start := time.Now()
		findDeepMove(board.Str2FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"), 5)
		duration := time.Since(start)
		fmt.Println(duration)
	*/
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/makeMove", makeMoveHandler)
	http.HandleFunc("/getAllowedMoves", getAllowedMovesHandler)
	http.HandleFunc("/getMove", getMoveHandler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	fmt.Printf("Starting server at port " + port + "\n")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/makeMove" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	// Read incoming request of format {"fen": ... , "from": ..., "to": ...}
	var moveRequest makeMoveRequest
	err := json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the response
	rawResponse := make(map[string]interface{})

	// Produce allowed moves of format {"e2-e4", "e2-e3"}
	var chessBoard board.Board
	chessBoard = board.Board{}
	initialFen := board.Str2FEN(moveRequest.Fen)
	chessBoard.Init(initialFen)
	chessBoard.GetOpponentVision()
	chessBoard.GetAllowedMoves()
	moveToMake := board.Move{
		From: board.Square2Pos(moveRequest.From),
		To:   board.Square2Pos(moveRequest.To),
	}
	newFen := chessBoard.Move(moveToMake)
	rawResponse["fen"] = board.Fen2Str(newFen)

	// Check for stalemate or checkmate
	boardAfterMove := board.Board{}
	boardAfterMove.Init(newFen)
	boardAfterMove.GetOpponentVision()
	boardAfterMove.GetAllowedMoves()
	rawResponse["checkMate"] = boardAfterMove.InCheckMate
	rawResponse["staleMate"] = boardAfterMove.InStaleMate

	// Marshal response and return
	res, err := json.Marshal(rawResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(res))
}

func getAllowedMovesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getAllowedMoves" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	// Read incoming request of format {"fen": ... , "from": ...}
	var position map[string]string
	err := json.NewDecoder(r.Body).Decode(&position)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Produce allowed moves of format {"e2-e4", "e2-e3"}
	var cb *board.Board
	cb = &board.Board{}
	initialFen := board.Str2FEN(position["fen"])
	cb.Init(initialFen)
	cb.GetOpponentVision()
	cb.GetAllowedMoves()
	fromSquare := board.Square2Pos(position["from"])
	var pieceMoves []string = cb.GetPieceMoves(fromSquare)
	rawResponse := make(map[string]interface{})
	rawResponse["moves"] = pieceMoves

	res, err := json.Marshal(rawResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(res))
}

func getMoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/getMove" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	// Read incoming request of format {"fen": ... , "from": ...}
	var moveRequest getMoveRequest
	err := json.NewDecoder(r.Body).Decode(&moveRequest)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rawResponse := make(map[string]interface{})
	preferredDepth := float64(moveRequest.Depth)

	// Produce allowed moves of format {"e2-e4", "e2-e3"}
	cb := board.Board{}
	initialFen := board.Str2FEN(moveRequest.Fen)
	cb.Init(initialFen)
	cb.GetOpponentVision()
	cb.GetAllowedMoves()
	depth := getDepth(preferredDepth, cb.Fen)
	fmt.Println("depth: ", depth)
	start := time.Now()
	rawResponse["nextMove"], rawResponse["evaluation"] = findDeepMove(initialFen, depth)
	duration := time.Since(start)
	fmt.Println("Time to completion:", duration)

	res, err := json.Marshal(rawResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, string(res))
}

func findDeepMove(fen board.FEN, depth int8) (string, string) {
	cb := board.Board{}
	cb.Init(fen)
	cb.GetOpponentVision()
	cb.GetAllowedMoves()
	if cb.InCheckMate {
		return "", "-Inf"
	} else if cb.InStaleMate {
		return "", "0"
	}

	allowedMoves := cb.AllowedMoves
	var currentEval float64
	var moveToEvaluate board.FEN
	maxEval := math.Inf(-1)
	var maxIdx int = -1
	var evaluations [][2]float64
	evalChan := make(chan [2]float64, len(allowedMoves))
	var waitGroup sync.WaitGroup
	for moveNum := 0; moveNum < len(allowedMoves); moveNum++ {
		waitGroup.Add(1)
		go func(moveNum int, evalChan chan [2]float64, wg *sync.WaitGroup) {
			defer waitGroup.Done()
			cb := board.Board{}
			cb.Init(fen)
			moveToEvaluate = cb.Move(allowedMoves[moveNum])
			alpha := math.Inf(-1)
			beta := math.Inf(1)
			currentEval = -evalToDepth(moveToEvaluate, depth, alpha, beta)
			evalChan <- [2]float64{float64(moveNum), currentEval}
		}(moveNum, evalChan, &waitGroup)
	}

	for range allowedMoves {
		evaluations = append(evaluations, <-evalChan)
	}
	waitGroup.Wait()

	for i := 0; i < len(allowedMoves); i++ {
		if evaluations[i][1] > maxEval {
			maxIdx = int(evaluations[i][0])
			maxEval = evaluations[i][1]
		} else if currentEval == maxEval {
			chooseMove := rand.Intn(2)
			if chooseMove == 1 {
				maxIdx = int(evaluations[i][0])
				maxEval = evaluations[i][1]
			}
		}
	}
	return board.Pos2Square(allowedMoves[maxIdx].From) + "-" + board.Pos2Square(allowedMoves[maxIdx].To), strconv.Itoa(int(maxEval))
}

func evalToDepth(fen board.FEN, depth int8, alpha, beta float64) float64 {
	cb := &board.Board{}
	cb.Init(fen)
	if depth == 0 {
		// return cb.Evaluate()
		return evalUntilNoCaptures(fen, alpha, beta)
	}
	cb.GetOpponentVision()
	cb.GetAllowedMoves()
	if cb.InCheckMate {
		return float64(-100000000) - float64(depth)
	}
	if cb.InStaleMate {
		return 0
	}
	allowedMoves := cb.AllowedMoves
	cbNext := &board.Board{}
	for i := 0; i < len(allowedMoves); i++ {
		nextFen := cb.Move(allowedMoves[i])
		cbNext.Init(nextFen)
		currentEval := -evalToDepth(nextFen, depth-1, -1*beta, -1*alpha)
		if currentEval >= beta {
			return beta
		}
		if currentEval >= alpha {
			alpha = currentEval
		}
	}
	return alpha
}

func evalUntilNoCaptures(fen board.FEN, alpha, beta float64) float64 {
	cb := &board.Board{}
	cb.Init(fen)
	currentEval := cb.Evaluate()
	if currentEval >= beta {
		return beta
	}
	if currentEval >= alpha {
		alpha = currentEval
	}
	cb.GetOpponentVision()
	captureMoves := cb.GetCaptureMoves()
	cbNext := &board.Board{}
	for i := 0; i < len(captureMoves); i++ {
		nextFen := cb.Move(captureMoves[i])
		cbNext.Init(nextFen)
		currentEval = -evalUntilNoCaptures(nextFen, -1*beta, -1*alpha)
		if currentEval >= beta {
			return beta
		}
		if currentEval >= alpha {
			alpha = currentEval
		}
	}
	return alpha
}

func getDepth(preferredDepth float64, fen board.FEN) int8 {

	numBishops := float64(strings.Count(fen.Position, "b") + strings.Count(fen.Position, "B"))
	numQueens := float64(strings.Count(fen.Position, "q") + strings.Count(fen.Position, "Q"))
	numRooks := float64(strings.Count(fen.Position, "r") + strings.Count(fen.Position, "R"))
	numPawns := float64(strings.Count(fen.Position, "p") + strings.Count(fen.Position, "P"))
	numKnights := float64(strings.Count(fen.Position, "n") + strings.Count(fen.Position, "N"))

	// complexity for starting position: 20 + 48 + 8 +12 = 88
	complexityEstimate := 10.0*numQueens + 6.0*(numRooks+numBishops) + 0.5*numPawns + 3.0*numKnights
	fmt.Println("Position complexity: ", complexityEstimate)
	var calculatedDepth float64
	if fen.FullMove <= 4 {
		calculatedDepth = 4
	} else if complexityEstimate < 20 {
		calculatedDepth = 7
	} else if complexityEstimate < 40 {
		calculatedDepth = 6
	} else if complexityEstimate <= 75 {
		calculatedDepth = 5
	} else {
		calculatedDepth = 4
	}

	return int8(math.Min(calculatedDepth, preferredDepth))
}
