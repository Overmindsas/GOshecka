package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	fmt.Printf("%T %v %T %v", err, err, nil, nil)
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

// выведется *main.customError <nil> <nil> <nil>error
// *main.customError - тип переменной err, а у nil тип nil
