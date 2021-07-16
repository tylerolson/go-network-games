package main

import "fmt"

type TicTacToeGame struct {
	board          [3][3]*TicTacToeMark
	rowX           [3]int
	columnX        [3]int
	diagX          [3]int
	diagOppX       [3]int
	rowO           [3]int
	columnO        [3]int
	diagO          [3]int
	diagOppO       [3]int
	moveCount      int
	winner         rune
	isCurrentMoveX bool
	gs             *GameScreen
}

func NewTicTacToeGame(gs *GameScreen) *TicTacToeGame {
	game := &TicTacToeGame{
		isCurrentMoveX: true,
		gs:             gs,
		winner:         ' ',
	}
	game.ClearBoard()
	return game
}

func (game *TicTacToeGame) ClearBoard() {
	num := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			game.board[i][j] = NewTicTacToeMark(i, j, 0, 0, num, rune(num)+48)
			num++
		}
	}
}

func (game *TicTacToeGame) DoMove(i, j int) {
	if game.winner != ' ' {
		return
	}

	moveRune := 'X'

	if !game.isCurrentMoveX {
		moveRune = 'O'
	}
	if game.board[i][j].mark == 'X' || game.board[i][j].mark == 'O' {
		game.gs.infoText = fmt.Sprintf("Cannot place %c at cell #%d, there is already an %c", moveRune, game.board[i][j].num, game.board[i][j].mark)
		return
	}
	if game.isCurrentMoveX {
		game.board[i][j].mark = 'X'
		game.rowX[i]++
		game.columnX[j]++
		if i == j {
			game.diagX[i]++
		}
		if i+j+1 == 3 {
			game.diagOppX[i]++
		}

		moveRune = 'O'
	} else {
		game.board[i][j].mark = 'O'
		game.rowO[i]++
		game.columnO[j]++
		if i == j {
			game.diagO[i]++
		}
		if i+j+1 == 3 {
			game.diagOppO[i]++
		}

		moveRune = 'X'
	}
	game.moveCount++
	game.CheckWinner(i, j)

	game.isCurrentMoveX = !game.isCurrentMoveX

	game.gs.infoText = fmt.Sprintf("Waiting for %c to place", moveRune)
}

func (game *TicTacToeGame) DoMoveKeyboard(moveNum int) {
	if game.winner != ' ' {
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.board[i][j].num == moveNum {
				game.DoMove(i, j)
				return
			}
		}
	}
}

func (game *TicTacToeGame) DoMoveMouse(x, y int) {
	if game.winner != ' ' {
		return
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.board[i][j].x == x && game.board[i][j].y == y {
				game.DoMove(i, j)
				return
			}
		}
	}
}

func (game *TicTacToeGame) PrintBoard() {
	line := ""

	drawBox(game.gs.s, (game.gs.width/2)-11, 1, (game.gs.width/2)+11, 7, game.gs.textStyle, line) // Draw outline box

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			line += string(game.board[i][j].mark) // Add row to one line to print

			if j != 2 {
				line += "  |  " // Add seperator
			}

		}
		drawText(game.gs.s, (game.gs.width/2)-(len(line)/2), 2+i*2, (game.gs.width/2)+(len(line)/2)+1, 2+i*2, game.gs.textStyle, line) // Draw board value
		if i > 0 {
			line = "---------------------"
			drawText(game.gs.s, (game.gs.width/2)-(len(line)/2), 1+i*2, (game.gs.width/2)+(len(line)/2)+1, 1+i*2, game.gs.textStyle, line) // Draw divider
		}
		line = ""
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x := (game.gs.width / 2) - (len(line) / 2) - (6*-j + 6)
			y := 2 + (2 * i)
			game.board[i][j].x = x
			game.board[i][j].y = y
		}
	}
}

func (game *TicTacToeGame) CheckWinner(row, col int) {
	xCount, xOppCount, oCount, oOppCount := 0, 0, 0, 0
	for i := 0; i < 3; i++ { // Check sum of diag lists
		xCount += game.diagX[i]
		xOppCount += game.diagOppX[i]

		oCount += game.diagO[i]
		oOppCount += game.diagOppO[i]
	}

	if game.rowX[row] == 3 || game.columnX[col] == 3 || xCount == 3 || xOppCount == 3 {
		game.winner = 'X'
	}
	if game.rowO[row] == 3 || game.columnO[col] == 3 || oCount == 3 || oOppCount == 3 {
		game.winner = 'O'
	}

	if game.moveCount >= 9 {
		game.winner = 'n'
	}
}

type TicTacToeMark struct {
	i, j, x, y, num int
	mark            rune
}

func NewTicTacToeMark(i, j, x, y, num int, mark rune) *TicTacToeMark {
	return &TicTacToeMark{
		i:    i,
		j:    j,
		x:    x,
		y:    y,
		num:  num,
		mark: mark,
	}
}
