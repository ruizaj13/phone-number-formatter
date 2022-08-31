package main

import (
	// "bytes"
	"Documents/go-practice/phone-number-formatter/db"
	"database/sql"
	"fmt"
	"regexp"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "juanruiz"
	password = "1234"
	dbname = "gophercises_phone"
)

func main() {
	//database info
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	must(db.Reset("postgres", psqlInfo, dbname))

	//We now include dbname per the docs to connect to our SQL database
	psqlInfo = fmt.Sprintf("%s dbname=%s", psqlInfo, dbname)
	must(db.Migrate("postgres", psqlInfo))
	
	db, err := db.Open("postgres", psqlInfo)
	must(err)
	defer db.Close()

	err = db.Seed()
	must(err)

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			fmt.Println("Updating or removing...", number)
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No changes required")
		}
	}
}

func insertPhone(db *sql.DB, phone string) (int, error) {
	//($1) is how you would pass a variable into the exec statement e.g. is would be
	//($1 $2) if you had two variables to pass in also helps protect from SQL injection
	statement := `INSERT INTO phone_numbers(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}



func must(err error) {
	if err != nil {
		panic(err)
	}
}


func normalize(phone string) string {
	//REGEX approach
	re := regexp.MustCompile("[^0-9]")

	return re.ReplaceAllString(phone, "")
}



//Buffer approach
// func normalize(phone string) string {
// 	// We want these - 0123456789

// 	var buf bytes.Buffer

// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buf.WriteRune(ch)
// 		}
// 	}

// 	return buf.String()
// }
