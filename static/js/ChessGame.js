
export class ChessGame {
    constructor(args) {
        if (args) {
            this.fen = args.fen;
            this.playerColor = args.playerColor;
        }
        this.clearNotifications();
        this.gameOver = false;
        this._selectedPieceMoves = [];

        this.chessBoard = ChessBoard('board', {
            position: 'start',
            draggable: true,
            onDragStart: this.onDragStart.bind(this),
            onDrop: this.onDrop.bind(this)
        });
        this.chessBoard.start(false);
        this.chessBoard.orientation(this.playerColor);
        if (this.playerColor == "black") {
            this.getMove().then((firstMove) => {
                this.makeMove(...firstMove.split('-'));
            })
        }
    }

    clearNotifications() {
        const notifyElement = document.querySelector('.game-result');
        if (notifyElement) {
            notifyElement.innerText = "";
        }
    }

    async makeMove(source, target) {
        return fetch('/makeMove', {
            headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify({
                "fen": this.fen,
                "from": source,
                "to": target
            })
        }).then(async (res) => {
            return res.json();
        }).then((res) => {
            this.chessBoard.position(res.fen, true)
            this.fen = res.fen;
            if (res.checkMate) {
                this.notifyCheckMate();
                this.gameOver == true;
            } else if (res.staleMate) {
                this.notifyStaleMate();
                this.gameOver = true;
            }
        }); 
    }

    async getMove() {
        const depth = 8;
        return fetch('/getMove', {
            headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify({
                "fen": this.fen,
                "depth": depth
            })
        }).then(async (res) => {
            return res.json();
        }).then((res) => {
            return res.nextMove; 
        }); 
    }

    async onDragStart(source) {
        this._selectedPieceMoves = await this.getAllowedSquares(source);
        if (this._selectedPieceMoves.length != 0) {
            this._selectedPieceMoves.forEach((square) => {
                this._makeSquareGrey(square)
            });
        }
    }

    async getAllowedSquares(square) {
        return fetch('/getAllowedMoves', {
            headers: {'Content-Type': 'application/json'},
            method: 'POST',
            body: JSON.stringify({
                "from": square,
                "fen": this.fen
            })
        }).then(function (res) {
            return res.json();
        }).then((res) => {
            let allowedMoves = [];
            if (res.moves.length != 0) {
                res.moves.forEach((move) => {
                    allowedMoves.push(move.split('-')[1]);
                })
            }
            return allowedMoves;
        });
    }

    onDrop(source, target) {
        this._removeGreySquares();
        if (this.gameOver || !this._selectedPieceMoves.includes(target)) {
            return 'snapback';
        }
        this.makeMove(source, target).then(async () => {
            if (!this.gameOver) {
                console.log(this.fen)
                await new Promise((res)=>{setTimeout(()=>{res()}, 500)})
                this.getMove().then((nextMove) => {
                    console.log(nextMove)
                    this.makeMove(...nextMove.split('-'));
                });
            }
        })
    }

    notifyCheckMate() {
        const notifyElement = document.querySelector('.game-result');
        const numberOfMoves = this.fen.split(' ')[5];
        const winner = this.fen.split(' ')[1] == "w" ? "White" : "Black";
        notifyElement.innerText = "Checkmate! " + winner + " wins after " + numberOfMoves + " moves!";
    }

    notifyStaleMate() {
        const notifyElement = document.querySelector('.game-result');
        const numberOfMoves = this.fen.split(' ')[5];
        notifyElement.innerText = "Stalemate after " + numberOfMoves + " moves!"; 

    }

    _makeSquareGrey(square) {
        var $square = $("#board .square-" + square);
        var background = '#a9a9a9';
        if ($square.hasClass("black-3c85d")) {
            background = '#696969';
        }
        $square.css("background", background);
    }

    _removeGreySquares() {
        $('#board .square-55d63').css('background', '')
    }
}