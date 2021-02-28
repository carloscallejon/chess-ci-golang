
import {ChessGame} from './js/ChessGame.js'

function startBoard() {
    document.getElementById("newGameWhite").onclick = newGameAsWhite;
    document.getElementById("newGameBlack").onclick = newGameAsBlack;
    newGameAsWhite();
}

function newGameAsWhite() {
    return newGame('white')
}

function newGameAsBlack() {
    return newGame('black')
}

function newGame(color) {
    const startingFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1";
    const chessGame = new ChessGame({
        fen: startingFen, 
        playerColor: color
    });
} 
export {startBoard};

