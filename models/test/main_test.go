package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// This file runs before ALL tests.
// This gives us the opportunity to run setup() and shutdown() functions... 
// ...before and after m.Run()
// http://cs-guy.com/blog/2015/01/test-main/

var ROOT_CONNECTION_STRING string = "mysql://root:admin@tcp(localhost:3306)/lake"
var MIGRATIONS_PATH string = "file://../../db/migration"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	runMigrationsDown()
	runMigrationsUp()
}

func runMigrationsUp() {
	m, err := migrate.New(
		MIGRATIONS_PATH,
		ROOT_CONNECTION_STRING)

	if err != nil {
		fmt.Println("ERROR: Could not init migrate for UP: ", err)
	}
	err = m.Up()
	if err != nil {
		fmt.Println("ERROR: Could not run migrations UP: ", err)
	}
}

func runMigrationsDown() {
	m, err := migrate.New(
		MIGRATIONS_PATH,
		ROOT_CONNECTION_STRING)

	if err != nil {
		fmt.Println("ERROR: Could not init migrate for DOWN: ", err)
	}
	err = m.Down()
	if err != nil {
		fmt.Println("ERROR: Could not run migrations DOWN: ", err)
	}
}
