package mines

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func readChar() string {
	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b []byte = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(ch)

	for {
		select {
		case stdin, _ := <-ch:
			fmt.Println("Keys pressed:", stdin)
		default:
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func readKey() int {
	buffer := []byte{0}
	i, _ := os.Stdin.Read(buffer)
	fmt.Println(buffer[0])
	return i
}

var ttyinit = false

func readKeyPress() byte {
	if !ttyinit {
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		ttyinit = true
	}
	var b []byte = make([]byte, 2)
	for {
		os.Stdin.Read(b)
		//		fmt.Println("I got the byte", b, "("+string(b)+")")
		return b[0]
	}
}
