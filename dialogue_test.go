package dialogue

import (
	"fmt"
	"testing"
	"time"
)

func TestFlush(t *testing.T) {
	fmt.Println("long operation...")
	time.Sleep(5 * time.Second)
	fmt.Print("Y/N? ")

	flush()

	var s string
	fmt.Scanln(&s)
	fmt.Printf("%q\n", s)
}
