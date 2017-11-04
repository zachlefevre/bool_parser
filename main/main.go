package main

import "../boolParse"
import "bufio"
import "os"
import "fmt"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Text: ")
	for scanner.Scan() {
		if scanner.Text() == `` {
			fmt.Println("Exiting")
			os.Exit(0)
		}
		get, stk := boolParse.CreateParser(scanner.Text())
		inp := get()
		syntax := boolParse.B(&inp, get, stk)
		value := stk.Pop()

		fmt.Println("syntactally correct: ", syntax)
		fmt.Println("value: ", value)
		fmt.Println("Enter Text: ")

	}
}
