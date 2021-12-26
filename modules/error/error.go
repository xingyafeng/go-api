package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {

	if arg == 43 {

		return -1, errors.New("can't work with 43 ...")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {

	return fmt.Sprintf("%d ---- %s", e.arg, e.prob)

}

func f2(arg int) (int, error) {

	if arg == 43 {

		return -1, &argError{arg, "can't work with it ..."}
		// return -1, &argError{arg: arg, prob: "can't work with it ..."}

	} else {

		return arg + 3, nil

	}
}

func main() {

	// fmt.Println("hello go")

	for _, i := range []int{7, 43} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("work:", r)
		}
	}

	for _, i := range []int{7, 43} {

		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("work: ", r)
		}
	}

	_, e := f2(43)
	if ae, yes := e.(*argError); yes {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
