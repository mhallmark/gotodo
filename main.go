package main

import (
	"github.com/mhallmark/gotodo/data"
	"github.com/mhallmark/gotodo/cmd"
)

func main() {
	data.Open()
	defer data.Close()
	
	cmd.Execute()
}
