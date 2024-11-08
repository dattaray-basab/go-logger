package main

import (
	logger "github.com/dattaray-basab/gologger" // Importing with alias 'logger'
)

func main() {
	// Now using the alias 'logger'
	logger.LogInfo("This is an info message")
	logger.LogWarning("This is a warning message")
	logger.LogError("This is an error message")
}
