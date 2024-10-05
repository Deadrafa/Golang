package main

import (
	"fmt"
)

const (
	STACK_SIZE = 20
)

type Stack_message struct {
	Data  [STACK_SIZE]string
	size  uint32
	Error string
}

func peek(st Stack_message) string {
	if st.size == 0 {
		return "STACK_IS_EMPTY"
	}

	return st.Data[st.size-1]

}

func pop(st *Stack_message) string {
	if st.size == 0 {
		return "STACK_IS_EMPTY"
	}
	st.size--
	return st.Data[st.size]
}

func Print_stack(s *Stack_message) {
	fmt.Printf("Stack size: %d\n", s.size)
	fmt.Print("Stack data: [")
	for i := uint32(0); i < s.size; i++ {
		fmt.Printf("%s", s.Data[i])
		if i < s.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func push(st *Stack_message, value string) string {
	if st.size >= STACK_SIZE {
		return "STACK_IS_FULL"
	}
	st.Data[st.size] = value
	st.size++
	return "Successes"
}

func init_stack(st *Stack_message) {
	st.Data = [STACK_SIZE]string{}
	st.size = 0
	st.Error = ""

}
