# brackets

This project requires Golang v1.17+ to run.

Install the dependencies

cd main
go get


# Start The project

go run main/main.go

This project requires entering a text consisting of parentheses. and then checks the symmetry of the parentheses.

Returns #YES if the text entered by the user is symmetrical, if it is not symmetrical return #NO.

Asks for new text after every return.


# Run test

cd test
go test


# Examples

Example 1 returns #YES

Enter text : ()()
YES

Example 2 returns #NO

Enter text : ((()]]
NO




