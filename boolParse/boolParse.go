package boolParse

import (
	"fmt"
)

/*
CreateParser returns a function which will return the next rune of a string each time
it is called
*/
func CreateParser(input string) func() byte {
	index := -1
	if input[len(input)-1] != '$' {
		input = string(append([]byte(input), '$'))
	}
	return func() byte {
		index++
		for input[index] == ' ' {
			index++
		}
		// fmt.Println("Analyzing: ", string(input[index]))
		// if input[index] == '$' {
		// 	fmt.Println("Reached end of string")
		// }
		return input[index]
	}

}

func IT_TAIL(lex *byte, get func() byte) bool {
	fmt.Println("In IT_TAIL")
	if *lex == '-' {
		*lex = get()
		if *lex == '>' {
			*lex = get()
			if OT(lex, get) {
				if IT_TAIL(lex, get) {
					return true
				}
			}
		} else {
			fmt.Println("Was expecting -> but received", lex)
		}
	}
	return true
}

func A(lex *byte, get func() byte) bool {
	fmt.Println("In A")
	if *lex == 'T' {
		*lex = get()
		return true
	}
	if *lex == 'F' {
		*lex = get()
		return true
	}
	if *lex == '(' {
		*lex = get()
		if IT(lex, get) {
			if *lex == ')' {
				*lex = get()
				return true
			}
		}
	}
	return false
}

func AT_TAIL(lex *byte, get func() byte) bool {
	fmt.Println("In AT_TAIL")
	if *lex == '^' {
		*lex = get()
		if L(lex, get) {
			if AT_TAIL(lex, get) {
				return true
			}
		} else {
			return false
		}
	}
	return true
}
func L(lex *byte, get func() byte) bool {
	fmt.Println("In L")

	if A(lex, get) {
		return true
	}
	if *lex == '~' {
		*lex = get()
		if L(lex, get) {
			return true
		}
	}
	return false
}

func AT(lex *byte, get func() byte) bool {
	fmt.Println("In AT")
	if L(lex, get) {
		if AT_TAIL(lex, get) {
			return true
		}
	}
	return false
}

func OT_TAIL(lex *byte, get func() byte) bool {
	fmt.Println("In OT_TAIL")
	if *lex == 'v' {
		*lex = get()
		if AT(lex, get) {
			if OT_TAIL(lex, get) {
				return true
			}
		}
	}
	return true
}

func OT(lex *byte, get func() byte) bool {
	fmt.Println("In OT")
	if AT(lex, get) {
		if OT_TAIL(lex, get) {
			return true
		}
	}
	return false
}

func IT(lex *byte, get func() byte) bool {
	fmt.Println("In IT")
	if OT(lex, get) {
		if IT_TAIL(lex, get) {
			return true
		}
	}
	return false
}
func B(lex *byte, get func() byte) bool {
	fmt.Println("In B")
	if IT(lex, get) {
		fmt.Println("Lex: ", string(*lex))
		if *lex == '.' {
			*lex = get()
			if *lex == '$' {
				return true
			}
		} else {
			fmt.Println("expected . but received", string(*lex))
		}
	}
	return false
}