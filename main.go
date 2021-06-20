package main

import (
	"fmt"

	"github.com/PimentaJoao/DFA-go/AFD"
)

func main() {

	afd := AFD.New("./examples/example1.txt")

	wordTest := AFD.Process(afd)

	for _, test := range wordTest {
		fmt.Print("Palavra: ", test.Word, "\n")
		fmt.Println("Estados percorridos: ")
		for _, w := range test.States {
			fmt.Print(" -> ", w)
		}
		fmt.Print("\nPalavra foi aceita?: ", test.AcceptedStatus, "\n\n")
	}

}
