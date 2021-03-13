package main

import (
	"fmt"

	"github.com/rusldv/kit/fileutil"
	"github.com/rusldv/kit/flagutil"
)

// User data for user
type User struct {
	ID     int    `flag:"testid"`
	Name   string `flag:"testname"`
	Status bool   `flag:"status"`
}

func main() {
	var u User
	flagutil.Parse(&u)
	fmt.Println(u)
	/*
		if err := flagutil.Parse(123); err != nil {
			fmt.Println(err)
		}
	*/
	if err := fileutil.WriteFileString("test.txt", "Hello, File!"); err != nil {
		fmt.Println(err)
	}

	str, err := fileutil.ReadFileString("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

}
