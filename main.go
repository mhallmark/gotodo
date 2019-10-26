package main

import (
	"fmt"

	"github.com/mhallmark/gotodo/cmd"
	"github.com/mhallmark/gotodo/data/todoitems"
)

func main() {
	err := todoitems.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer todoitems.Close()
	cmd.Execute()
}
