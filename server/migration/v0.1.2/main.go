package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rubenv/sql-migrate"
	"log"
	"os"
)

func main() {
	log.Println("START Migration...")
	command := flag.String("command", "new", "Migration action. (\"new\" | \"up\" | \"down\")")

	MySQLHost := flag.String("host", "127.0.0.1", "MySQL Host destination")
	MySQLPort := flag.Uint("port", 3306, "MySQL Database Port")
	MySQLUser := flag.String("user", "user", "MySQL User ID")
	MySQLPassword := flag.String("password", "password", "MySQL User password")
	MySQLDB := flag.String("db", "get5", "MySQL Dtabase name")

	flag.Parse()

	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", *MySQLUser, *MySQLPassword, *MySQLHost, *MySQLPort, *MySQLDB)
	db, err := sql.Open("mysql", sqloption)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	switch *command {
	case "new":
		log.Println("Initializing Database...")
		// OR: Read migrations from a folder:
		migrations := &migrate.FileMigrationSource{
			Dir: "./schema",
		}
		n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
		if err != nil {
			log.Printf("Failed to migrate database! ERR : %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Applied %d migrations!\n", n)

	case "up":
		log.Println("Upgrating...")
		// Hardcoded strings in memory:
		migrations := &migrate.MemoryMigrationSource{
			Migrations: []*migrate.Migration{
				&migrate.Migration{
					Id: "123",
					// Up: []string{"CREATE TABLE people (id int)"}, // TODO...
					// Down: []string{"DROP TABLE people"}, // TODO...
				},
			},
		}
		n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
		if err != nil {
			log.Printf("Failed to migrate database! ERR : %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	case "down":
		log.Println("Downgrating...")
		// Hardcoded strings in memory:
		migrations := &migrate.MemoryMigrationSource{
			Migrations: []*migrate.Migration{
				&migrate.Migration{
					Id: "123",
					// Up: []string{"CREATE TABLE people (id int)"}, // TODO...
					// Down: []string{"DROP TABLE people"}, // TODO...
				},
			},
		}
		n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
		if err != nil {
			log.Printf("Failed to migrate database! ERR : %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	default:
		log.Printf("Unsupported command. command : %s\n", *command)
	}
}
