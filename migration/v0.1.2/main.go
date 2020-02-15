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
	command := flag.String("command", "up", "Migration action. (up|down)")

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
	case "up":
		log.Println("Upgrating...")
		migrations := &migrate.FileMigrationSource{
			Dir: "./schema",
		}
		n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
		if err != nil {
			log.Printf("Failed to migrate database! ERR : %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	case "down":
		log.Println("Downgrating...")
		migrations := &migrate.FileMigrationSource{
			Dir: "./schema",
		}
		n, err := migrate.Exec(db, "mysql", migrations, migrate.Down)
		if err != nil {
			log.Printf("Failed to migrate database! ERR : %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Applied %d migrations!\n", n)
	default:
		log.Printf("Unsupported command. command : %s\n", *command)
	}
}
