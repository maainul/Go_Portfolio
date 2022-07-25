package postgres

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"fmt"
)

func TestMain(m *testing.M) {
	const dbConnEnv = "DATABASE_CONNECTION"
	ddlConnStr := os.Getenv(dbConnEnv)
	fmt.Println("###################################")
	fmt.Println("###################################")
	fmt.Println("###################################")
	fmt.Println(ddlConnStr)
	fmt.Println("###################################")
	fmt.Println("###################################")
	fmt.Println("###################################")
	if ddlConnStr == "" {
		log.Printf("%s is not set, skipping", dbConnEnv)
		return
	}

	var teardown func()
	_testStorage, teardown = NewTestStorage(ddlConnStr, filepath.Join("..", "..", "migrations", "sql"))

	exitCode := m.Run()

	if teardown != nil {
		teardown()
	}

	os.Exit(exitCode)
}
