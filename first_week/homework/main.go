package main

import (
	"errors"
	"fmt"
	"net/http"
)

var ERRNOACK error = errors.New("no ack")

func searchMem()  {
	
}

func main() {
	_, err := InitSql("root:123456@tcp(127.0.0.1:3306)/member")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := InitApiServer()
	http.ListenAndServe(":10000", server)
	//test()
}

func test() (error) {
	err := searchAge()
	if err != nil {
		if errors.Is(err, ERRNOACK) {
			fmt.Println("test")
		}
	}
	fmt.Println("err test")
	return nil
}

func searchAge() (error) {
	return fmt.Errorf("%w", ERRNOACK)
}
