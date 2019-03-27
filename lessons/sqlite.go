package lessons

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	DbConnection, _ = sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := `CREATE TABLE IF NOT EXISTS person(
		  	name STRING,
			age INT);`

	_, error := DbConnection.Exec(cmd)
	if error != nil {
		log.Fatalln("error", error)
	}

	// INSERT
	//cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	//_, error = DbConnection.Exec(cmd, "Mike", 24)
	//if error != nil {
	//	log.Fatalln("error", error)
	//}

	// UPDATE
	//cmd = "UPDATE person SET age = ? WHERE name = ?"
	//_, err := DbConnection.Exec(cmd, 25, "Mike")
	//if err != nil {
	//
	//	log.Fatalln("error", error)
	//}

	// MULTIPLE SELECT
	//cmd = "SELECT * FROM person"
	//rows, _ := DbConnection.Query(cmd)
	//defer rows.Close()
	//var pp []Person
	//for rows.Next() {
	//
	//	var p Person
	//	err := rows.Scan(&p.Name, &p.Age)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	pp = append(pp, p)
	//}
	//for _, p := range pp {
	//	fmt.Println(p.Name, p.Age)
	//}

	// SINGLE SELECTION
	//cmd = "SELECT * FROM person where age = ?"
	//row := DbConnection.QueryRow(cmd, 20)
	//var p Person
	//err := row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row!")
	//	} else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	cmd = "DELETE FROM person WHERE name = ?"

	_, err := DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	tableName := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next() {

		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}