package main

import (
	"fmt"
	"math/rand"
	"time"
	//"bufio"
	//"os"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func main() {

	fmt.Print("How many cups?")

	var numberofCups int

	_, error := fmt.Scanln(&numberofCups)
	if error != nil {
		fmt.Print("We have a problem")
		return
	}

	fmt.Println(numberofCups)
	cupNumber := random(1, numberofCups)
	
	var cupSelect int

	/*reader := bufio.NewReader(os.Stdin)
	text,_:=reader.ReadString('\n')

	fmt.Println(text)*/

	
	for{ 
		fmt.Println("Which cup?")
		fmt.Scanln(&cupSelect)
		println(cupSelect)

		if cupNumber == cupSelect {
			fmt.Println("Congrats!")
			break
		} else {
			fmt.Println("try again")
		}
	}

}
