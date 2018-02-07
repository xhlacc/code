// Sample program to show how to show you how to briefly work
// with the sql package.
package main

import (
	"database/sql"

	//this is the relative path
	_ "./postgres"
	//this is the origin source
	//_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

// main is the entry point for the application.
func main() {
	sql.Open("postgres", "mydb")
}
