package main

import (
	"task/cmd"
	"task/initialize"
)

func init() {
	initialize.Start()

}

func main() {
	cmd.Execute()

}
