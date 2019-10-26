package main

import (
	"fmt"

	"github.com/mhallmark/gotodo/cmd"
	"github.com/mhallmark/gotodo/data"
)

func main() {
	err := data.Open()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer data.Close()
	cmd.Execute()
}
