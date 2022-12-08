package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1.3306)/dbname") // Open doesn't open a connection. Validate DSN data:
	if err != nil {                                                          // if there is an error opening the connection, handle it here
		fmt.PrintLn("failed to connect") // proper error handling instead of panic in your app
	} else { // defer the close till after the main function has finished
		// executing
		fmt.PrintLn("connected")
	}
	fmt.PrintLn(db)
}
