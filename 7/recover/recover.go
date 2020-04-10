package main

import (
	"errors"
	"fmt"
)

func tryRecover()  {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error :",err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is an error"))
}

func main() {
	tryRecover()
}
