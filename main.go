package main

import (
	"fmt"

	"github.com/PimentaJoao/DFA-go/AFD"
)

func main() {

	afd := AFD.New("./examples/example1.txt")

	// states, success := AFD.Process(afd)

	// fmt.Println("Estados percorridos: ", states)
	// fmt.Println("Palavra foi aceita?: ", success)

	for i := 0; i < len(afd.Estados); i++ {
		for j := 0; j < len(afd.Estados); j++ {
			bla := fmt.Sprint(afd.Transicoes[i][j])
			fmt.Print(bla)
		}
		fmt.Println("\n-------------")
	}

}
