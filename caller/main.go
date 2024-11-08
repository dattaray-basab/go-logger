package main

import gologger "github.com/dattaray-basab/gologger"


func main() {
	// Using the default package name
	gologger.LogInfo("This is an info message")
	gologger.LogWarning("This is a warning message")
	gologger.LogError("This is an error message")
}
