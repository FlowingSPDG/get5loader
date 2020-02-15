package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

const (
	target = "20000101000000-create_tables.sql"
)

func main() {
	log.Println("START Migration...")

	MySQLHost := flag.String("host", "127.0.0.1", "MySQL Host destination")
	MySQLPort := flag.Uint("port", 3306, "MySQL Database Port")
	MySQLUser := flag.String("user", "user", "MySQL User ID")
	MySQLPassword := flag.String("password", "password", "MySQL User password")
	MySQLDB := flag.String("db", "get5", "MySQL Dtabase name")

	flag.Parse()

	sqloption := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", *MySQLUser, *MySQLPassword, *MySQLHost, *MySQLPort, *MySQLDB)
	db, err := sql.Open("mysql", sqloption)
	if err != nil {
		log.Printf("Failed to initialize db... ERR : %v\n", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	log.Println("Upgrating...")

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS `gorp_migrations` ( `id` varchar(255) NOT NULL, `applied_at` datetime DEFAULT NULL ) ENGINE=InnoDB DEFAULT CHARSET=utf8;"); err != nil {
		log.Printf("Failed to initialize db... ERR : %v\n", err.Error())
		os.Exit(1)
	}
	db.Exec("ALTER TABLE `gorp_migrations` ADD PRIMARY KEY (`id`);")

	Now := time.Now()
	AppliedAt := Now.Format("2006-01-2 15:04:05")

	if _, err := db.Exec("INSERT INTO `gorp_migrations` (id, applied_at) VALUE (?,?)", target, AppliedAt); err != nil {
		log.Printf("Failed to initialize db... ERR : %v\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Database %s is ready to begin migaration...\n", *MySQLDB)
}
