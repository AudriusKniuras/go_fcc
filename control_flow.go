package main

import (
	"fmt"
	"net/http"
)

func defer_func() {
	// defer - function is only executed when the surrounding function returns
	// for example defer file.Close() to make sure the file is closed before the function ends
	fmt.Println("start - defer")
	defer fmt.Println("middle - defer")
	fmt.Println("end - defer")

	// defer is called last to first order
	defer fmt.Println("first - defer order")
	defer fmt.Println("last - defer order")

	// defer f-tion arguments are evaluated immediately
	a := "start - defer is evaluated immediately"
	defer fmt.Println(a)
	a = "end"
}

func panic_ex1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// go very rarely throws panic (exception), so we need to decide when an error is panic worthy
		panic(err.Error())
	}
}
