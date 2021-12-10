package mines

import (
	"fmt"
	"time"
)

func printBomb() {
	fmt.Println(Red)
	for i := 0; i < 10; i++ {
		//		fmt.Println("(•_•)   ( •_•)  >⌐■-■    (⌐■_■)")
		for _, str := range bomb {
			fmt.Print(str)
			fmt.Print("\r")
			time.Sleep(100 * time.Millisecond)
		}

	}
}

var deal = []string{
	"  (•_•)   ⌐■-■",
	"  (•_•)   ⌐■-■",
	"  (•_•)   ⌐■-■",
	"  ( •_•)  ⌐■-■",
	"  ( •_•)  ⌐■-■",
	"  ( •_•)  ⌐■-■",
	"  ( •_•) ⌐■-■ ",
	"  ( •_•)⌐■-■  ",
	"  ( •_•⌐■-■   ",
	"  ( •_⌐■-■    ",
	"  ( •⌐■-■     ",
	"  ( ⌐■-■      ",
	"  (⌐■_■)      ",
	"  (⌐■_■)      ",
	"  (⌐■_■)      ",
	"  (⌐■_■)     C",
	"  (⌐■_■)    CO",
	"  (⌐■_■)   COO",
	"  (⌐■_■)  COOL",
	"  (⌐■_■) COOL!",
	"  (⌐■_■) COOL!",
	"  (⌐■_■) COOL!",
	"  (⌐■_■) COOL!",
}

var bomb = []string{
	"     ====---------",
	"     ====--------*",
	"     ====-------* ",
	"     ====------*  ",
	"     ====-----*   ",
	"     ====----*    ",
	"     ====---*     ",
	"     ====--*      ",
	"     ====-*       ",
	"     ====*        ",
	"       (BO)       ",
	"      (BOOM)      ",
	"     ((BOOM))     ",
	"    (((BOOM)))    ",
	"   ( ((BOOM)) )   ",
	"  ((  (BOOM)  ))  ",
	" ( (  (BOOM)  ) ) ",
	"( ((  (BOOM) )) ) ",
	"                  ",
	"( ((  (BOOM) )) ) ",
	"                  ",
	"( ((  (BOOM) )) ) ",
}

func printDealWithIt() {
	fmt.Println(Green)
	for i := 0; i < 10; i++ {
		//		fmt.Println("(•_•)   ( •_•)  >⌐■-■    (⌐■_■)")
		for _, str := range deal {
			fmt.Print(str)
			fmt.Print("\r")
			time.Sleep(200 * time.Millisecond)
		}

	}
}
