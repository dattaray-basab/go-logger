package main

// import "github.com/dattaray-basab/gologger/producer/gologger/v2"
import "github.com/dattaray-basab/gologger/producer/v2/gologger"


func main() {
	// Using the default package name
	gologger.LogInfo("This is an info message")
	gologger.LogWarning("This is a warning message")
	gologger.LogError("This is an error message")
}
