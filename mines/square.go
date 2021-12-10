package mines

import "fmt"

var (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type square struct {
	isBomb    bool
	isShown   bool
	bombCount int
	isMarked  bool
}

func (this square) Print() {
	if this.isMarked {
		fmt.Printf("%s%sX%s", Bold, Red, Reset)
	} else if !this.isShown {
		fmt.Printf("%s%s?%s", Bold, Yellow, Reset)
	} else if this.bombCount > 0 {
		fmt.Printf("%s%s%d%s", Bold, Purple, this.bombCount, Reset)
	} else if this.isBomb {
		fmt.Printf("%s%sO%s", Bold, Red, Reset)
	} else {
		fmt.Printf("%s%s.%s", Bold, Green, Reset)
	}
}
