package main

import (
	"log"
	"os"

	//这个改为相对路径 为了可以直接编译运行
	//This is the relative path
	//for can go build and run
	_ "./matchers"
	"./search"
	//这个是原始代码 需要目录
	//This is the origin source
	//_ "github.com/goinaction/code/chapter2/sample/matchers"
	//"github.com/goinaction/code/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
