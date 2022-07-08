package main

import (
	"fmt"
)

func main() {

	err := initApp()
	if err != nil {
		fmt.Println("[main][initApp] error when initializing app: ", err)
	}
	
}