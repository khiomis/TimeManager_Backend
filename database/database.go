package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"strings"
)

var Db *sqlx.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	pass := os.Getenv("PASSWORD")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)
	db, errSql := sqlx.Connect("postgres", psqlSetup)
	if errSql != nil {
		log.Fatalln("There is an error while connecting to the database ", errSql)
	} else {
		Db = db
		log.Println("Successfully connected to database!")
	}

	content, err := os.ReadFile("sql/schema_1.sql")
	schema := string(content)

	if strings.Contains(schema, "-- Already executed: false") {
		Db.MustExec(schema)
		log.Println("Schema created successfully!")
		schema = strings.Replace(schema, "-- Already executed: false", "-- Already executed: true", 1)
		errWrite := os.WriteFile("sql/schema_1.sql", []byte(schema), 0644)
		if errWrite != nil {
			log.Fatalln("There is an error while writing the file", err)
		}
	}

}
