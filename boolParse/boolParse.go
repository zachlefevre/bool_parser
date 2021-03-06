package boolParse

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

/*
CreateParser returns a function which will return the next rune of a string each time the function is called,
	and a stack.
*/
func CreateParser(input string) (func() byte, *stack.Stack) {
	index := -1
	if input[len(input)-1] != '$' {
		input = string(append([]byte(input), '$'))
	}
	return func() byte {
			index++
			for input[index] == ' ' {
				index++
			}
			return input[index]
		},
		stack.New()

}

func IT_TAIL(lex *byte, get func() byte, s *stack.Stack) bool {
	if *lex == '-' {
		*lex = get()
		if *lex == '>' {
			*lex = get()
			if OT(lex, get, s) {
				if IT_TAIL(lex, get, s) {
					p := s.Pop()
					q := s.Pop()
					fmt.Println(p, q)
					if q.(bool) {
						if !p.(bool) {
							s.Push(false)
						} else {
							s.Push(true)
						}
					} else {
						s.Push(true)
					}
					return true
				}
			}
		} else {
		}
	}
	return true
}

func A(lex *byte, get func() byte, s *stack.Stack) bool {
	if *lex == 'T' {
		s.Push(true)
		*lex = get()
		return true
	}
	if *lex == 'F' {
		s.Push(false)
		*lex = get()
		return true
	}
	if *lex == '(' {
		*lex = get()
		if IT(lex, get, s) {
			if *lex == ')' {
				*lex = get()
				return true
			}
		}
	}
	return false
}

func AT_TAIL(lex *byte, get func() byte, s *stack.Stack) bool {
	if *lex == '^' {
		*lex = get()
		if L(lex, get, s) {
			p := s.Pop()
			q := s.Pop()
			s.Push(q.(bool) && p.(bool))
			if AT_TAIL(lex, get, s) {
				return true
			}
		} else {
			return false
		}
	}
	return true
}
func L(lex *byte, get func() byte, s *stack.Stack) bool {
	if A(lex, get, s) {
		return true
	}
	if *lex == '~' {
		*lex = get()
		if L(lex, get, s) {
			p := s.Pop()
			if p.(bool) {
				s.Push(false)
			} else {
				s.Push(true)
			}
			return true
		}
	}
	return false
}

func AT(lex *byte, get func() byte, s *stack.Stack) bool {
	if L(lex, get, s) {
		if AT_TAIL(lex, get, s) {
			return true
		}
	}
	return false
}

func OT_TAIL(lex *byte, get func() byte, s *stack.Stack) bool {
	if *lex == 'v' {
		*lex = get()
		if AT(lex, get, s) {
			p := s.Pop()
			q := s.Pop()

			s.Push(p.(bool) || q.(bool))
			if OT_TAIL(lex, get, s) {
				return true
			}
		}
	}
	return true
}

func OT(lex *byte, get func() byte, s *stack.Stack) bool {
	if AT(lex, get, s) {
		if OT_TAIL(lex, get, s) {
			return true
		}
	}
	return false
}

func IT(lex *byte, get func() byte, s *stack.Stack) bool {
	if OT(lex, get, s) {
		if IT_TAIL(lex, get, s) {
			return true
		}
	}
	return false
}
func B(lex *byte, get func() byte, s *stack.Stack) bool {
	if IT(lex, get, s) {
		if *lex == '.' {
			*lex = get()
			if *lex == '$' {
				return true
			}
		} else {
		}
	}
	return false
}
