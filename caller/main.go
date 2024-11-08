package main

import "github.com/dattaray.basab/go-logger"

func main() {
	logger.LogInfo("This is an information message")
	logger.LogWarning("This is a warning message")
	logger.LogError("This is an error message")
}