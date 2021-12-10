package mines

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	x       int
	y       int
	Matrix  [][]square
	cursorX int
	cursorY int
	Ended   bool
}

func NewGame(x int, y int, bombs int) Game {
	game := Game{x: x, y: y}

	game.Matrix = make([][]square, y)
	for i := 0; i < y; i++ {
		game.Matrix[i] = make([]square, x)
	}

	rand.Seed(time.Now().UnixNano())
	for bombs > 0 {
		posx := rand.Intn(x)
		posy := rand.Intn(y)

		if !game.Matrix[posy][posx].isBomb {
			game.Matrix[posy][posx].isBomb = true
			bombs--
		}

	}

	game.SetBombCounter()
	return game
}

func (this *Game) Print() {
	//fmt.Println("\033[2J")
	fmt.Println("\033[H")
	fmt.Print("\nArrows to move, space for select and X for mark\n\n")
	for y, line := range this.Matrix {
		for x, item := range line {
			if x == this.cursorX && y == this.cursorY {
				fmt.Print("[")
				item.Print()
				fmt.Print("]")
			} else {
				fmt.Print(" ")
				item.Print()
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (this *Game) Input() {

	key := readKeyPress()
	switch key {
	case 65:
		this.moveUp()
	case 66:
		this.moveDown()
	case 67:
		this.moveRight()
	case 68:
		this.moveLeft()
	case 32:
		this.Press()
	case 120:
		this.Mark()
	}
}

func (this *Game) moveUp() {
	if this.cursorY > 0 {
		this.cursorY--
	}
}

func (this *Game) moveLeft() {
	if this.cursorX > 0 {
		this.cursorX--
	}
}

func (this *Game) moveDown() {
	if this.cursorY < this.y-1 {
		this.cursorY++
	}
}

func (this *Game) moveRight() {
	if this.cursorX < this.x-1 {
		this.cursorX++
	}
}

func (this Game) PrintPos() {
	fmt.Printf("( %d %d )", this.cursorX, this.cursorY)
}

func (this *Game) Mark() {
	this.Matrix[this.cursorY][this.cursorX].isMarked = !this.Matrix[this.cursorY][this.cursorX].isMarked
}

func (this *Game) Press() {
	if !this.Matrix[this.cursorY][this.cursorX].isMarked {
		if this.Matrix[this.cursorY][this.cursorX].isBomb {
			this.Ended = true
		} else {
			this.expandNoBombsSelection(this.cursorX, this.cursorY)
			this.Ended = this.checkGameFinished()
		}
	}
}

func (this *Game) showAll() {
	for posy := 0; posy < this.y; posy++ {
		for posx := 0; posx < this.x; posx++ {
			this.Matrix[posy][posx].isShown = true
		}
	}
}

func (this *Game) SetBombCounter() {
	for posy := 0; posy < this.y; posy++ {
		for posx := 0; posx < this.x; posx++ {
			if !this.Matrix[posy][posx].isBomb {
				this.Matrix[posy][posx].bombCount = this.getBombCounter(posx, posy)
			}
		}
	}

}

func (this Game) getBombCounter(x int, y int) int {
	ret := this.contBomb(x+1, y+1)
	ret += this.contBomb(x+1, y)
	ret += this.contBomb(x+1, y-1)
	ret += this.contBomb(x, y+1)
	ret += this.contBomb(x, y-1)
	ret += this.contBomb(x-1, y+1)
	ret += this.contBomb(x-1, y)
	ret += this.contBomb(x-1, y-1)
	return ret
}

func (this Game) contBomb(posx int, posy int) int {
	if posx < 0 || posy < 0 || posx >= this.x || posy >= this.y {
		return 0
	} else if this.Matrix[posy][posx].isBomb {
		return 1
	} else {
		return 0
	}

}

func (this Game) expandNoBombsSelection(x int, y int) {
	//	fmt.Printf(" x:%d   y:%d", x, y)
	if x < 0 || y < 0 || x >= this.x || y >= this.y {
	} else if this.Matrix[y][x].isShown {
	} else if this.Matrix[y][x].isBomb {
	} else if this.Matrix[y][x].bombCount > 0 {
		this.Matrix[y][x].isShown = true
	} else {
		this.Matrix[y][x].isShown = true
		this.expandNoBombsSelection(x+1, y+1)
		this.expandNoBombsSelection(x+1, y)
		this.expandNoBombsSelection(x+1, y-1)
		this.expandNoBombsSelection(x, y+1)
		this.expandNoBombsSelection(x, y-1)
		this.expandNoBombsSelection(x-1, y+1)
		this.expandNoBombsSelection(x-1, y)
		this.expandNoBombsSelection(x-1, y-1)
	}
}

func (this *Game) checkGameFinished() bool {
	for posy := 0; posy < this.y; posy++ {
		for posx := 0; posx < this.x; posx++ {
			if !this.Matrix[posy][posx].isShown && !this.Matrix[posy][posx].isBomb {
				return false
			}
		}
	}
	return true
}

func (this Game) GoodBye() {
	if this.checkGameFinished() {
		this.showAll()
		this.Print()
		printDealWithIt()
	} else {
		this.showAll()
		this.Print()
		printBomb()
	}
}
