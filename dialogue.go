package dialog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	YesDefintions     = []string{"y", "Y", "yes", "YES", "Yes"}
	NoDefinitions     = []string{"n", "N", "no", "NO", "No"}
	YesOrNoSuffix     = "? [Y/N]: "
	ErrProvideYesOrNo = errors.New("please answer with 'yes' or 'no'")
)

type Dialog struct {
}

//YesOrNo ask a question that can be answered by yes or no
func YesOrNo(question string, def bool) bool {
	//ask
	fmt.Printf(question + YesOrNoSuffix)

	//answer
	s, _ := readline()
	if hasString(YesDefintions, s) {
		return true
	} else if hasString(NoDefinitions, s) {
		return false
	} else if s == "" {
		return def
	} else {
		fmt.Println(ErrProvideYesOrNo.Error())
		return YesOrNo(question, def)
	}
}

//AskString asks for a string
func AskString(question string) string {
	fmt.Printf("%s:\n\t", question)
	s, _ := readline()
	return s
}

//readline reads input from the user byte by byte
func readline() (value string, err error) {
	var valb []byte
	var n int
	b := make([]byte, 1)
	for {
		// read one byte at a time so we don't accidentally read extra bytes
		n, err = os.Stdin.Read(b)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 || b[0] == '\n' {
			break
		}
		valb = append(valb, b[0])
	}

	return strings.TrimSuffix(string(valb), "\r"), nil
}

//indexOfString returns the index a string in a string slice, returns -1
//if the given string is not found
func indexOfString(h []string, n string) int {
	for i, v := range h {
		if v == n {
			return i
		}
	}

	return -1
}

//hasString is a wrapper around indexOfString
func hasString(h []string, n string) bool {
	if indexOfString(h, n) != -1 {
		return true
	}
	return false
}
