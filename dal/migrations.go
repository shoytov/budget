package dal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

func ApplyMigration(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
}

func ApplyMigrations() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := pwd + "/dal/migrations"

	db, err := sql.Open("sqlite3", "db.sqlite3")

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			filename := file.Name()
			migration, err := os.ReadFile(dir + "/" + filename)
			if err != nil {
				panic(err)
			}
			ApplyMigration(db, string(migration))
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
