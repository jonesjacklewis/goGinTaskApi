package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Database_Connection is a global variable that holds the database connection
var Database_Connection *sql.DB

// getDbFolderPath returns the path to the database folder
// If the folder does not exist, it creates it
// It returns the path to the database folder
// It panics if it fails to get the current working directory or create the folder
func getDbFolderPath() string {

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	db_folder_path := filepath.Join(cwd, "db")

	if _, err := os.Stat(db_folder_path); os.IsNotExist(err) {
		if err := os.Mkdir(db_folder_path, os.ModePerm); err != nil {
			panic(err)
		}
		fmt.Println("Database folder created at:", db_folder_path)
	}

	return db_folder_path
}

// createTables creates the tables in the database if they do not already exist
// It returns an error if it fails to create the tables
func createTables() error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		DisplayName TEXT NOT NULL UNIQUE
	);`

	createTasksTable := `
	CREATE TABLE IF NOT EXISTS Tasks (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		TaskHeader TEXT NOT NULL,
		TaskDescription TEXT NOT NULL,
		Complete BOOL DEFAULT FALSE
	);`

	createUsersTasksTable := `
	CREATE TABLE IF NOT EXISTS UsersTasks (
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		UsersId INTEGER NOT NULL,
		TasksId INTEGER NOT NULL,
		FOREIGN KEY (UsersId) REFERENCES Users(Id),
		FOREIGN KEY (TasksId) REFERENCES Tasks(Id)
	);`

	// Execute the table creation queries
	_, err := Database_Connection.Exec(createUsersTable)
	if err != nil {
		return fmt.Errorf("failed to create Users table: %w", err)
	}

	_, err = Database_Connection.Exec(createTasksTable)
	if err != nil {
		return fmt.Errorf("failed to create Tasks table: %w", err)
	}

	_, err = Database_Connection.Exec(createUsersTasksTable)
	if err != nil {
		return fmt.Errorf("failed to create UsersTasks table: %w", err)
	}

	fmt.Println("Tables created (if not already present)")
	return nil
}

// InitDb initializes the database connection and creates the tables if they do not already exist
// It returns an error if it fails to initialize the database or create the tables
func InitDb() error {
	db_path := filepath.Join(getDbFolderPath(), "tasks.db")

	var err error                                           // Declare error separately
	Database_Connection, err = sql.Open("sqlite3", db_path) // Assign directly to the global db variable

	if err != nil {
		panic(err)
	}

	if err := Database_Connection.Ping(); err != nil {
		panic(err)
	}

	if err := createTables(); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

// CloseDb closes the database connection
// It panics if it fails to close the connection
func CloseDb() {
	if err := Database_Connection.Close(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection closed")
}
