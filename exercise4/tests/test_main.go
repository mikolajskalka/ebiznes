package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/mikolajskalka/ebiznes/exercise4/database"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting test suite...")

	// Initialize the database
	database.Initialize()

	// Run the tests
	exitCode := m.Run()

	fmt.Println("All tests completed.")

	// Exit with the test result code
	os.Exit(exitCode)
}
