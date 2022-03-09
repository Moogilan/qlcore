package main

import (
	"database/sql"

	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host        = "localhost"
	port        = 5432
	user        = "postgres"
	password    = "user"
	dbname      = "test"
	DB_USER     = "postgres"
	DB_PASSWORD = "user"
	DB_NAME     = "test"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("QL server!")
	})

	app.Get("/retrieve", retrieveRecord)

	// http.HandleFunc("/retrieve", retrieveRecord) // (1)
	// http.ListenAndServe(":3000", nil)

	app.Use("/cardaccessraw", func(c *fiber.Ctx) error {

		fmt.Println(string(c.Body()))

		csv := string(c.Body())

		csvarray := strings.Split(csv, ",")

		d, err := time.Parse("02/01/2006", csvarray[2])
		if err != nil {
			log.Fatal(err)
		}

		t, err := time.Parse("15:04:05", csvarray[3])
		if err != nil {
			log.Fatal(err)
		}

		f := d.Add(time.Hour*time.Duration(t.Hour()) + time.Minute*time.Duration(t.Minute()) + time.Second*time.Duration(t.Second()) + time.Nanosecond*time.Duration(t.Nanosecond()))

		savedatabase(csvarray[0], csvarray[1], f.Format("2006-01-02 15:04:05"))
		return c.SendString("")

	})

	// http.HandleFunc("/retrieve", retrieveRecord)
	// http.ListenAndServe(":3000",nil)
	// (2)

	app.Listen(":3000")

	// http.HandleFunc("/retrieve", retrieveRecord) // (1)
	// http.ListenAndServe(":8000", nil)            // (2)

}

// type Employee struct {

// id int
// 	Employee_id string
// 	Employeename        string
// 	Employeetimestamp   string
// }

// type JsonResponse struct {
// 	Type    string     `json:"type"`
// 	Data    []Employee `json:"data"`
// 	Message string     `json:"message"`
// }

// func GetData(w http.ResponseWriter, r *http.Request) ([]byte, error) {

// 	db := saveg()

// 	printMessage("Getting data...")

// 	// Get all movies from movies table that don't have movieID = "1"
// 	rows, err := db.Query("SELECT * FROM public.card_access_log")

// 	// check errors
// 	checkErr(err)

// 	// var response []JsonResponse
// 	var data []Employee

// 	// Foreach movie
// 	for rows.Next() {
// 		var id int
// 		var employeeID string
// 		var employeeName string

// 		err = rows.Scan(&id, &employeeID, &employeeName)

// 		// check errors
// 		checkErr(err)

// 		data = append(data, Employee{Employeeemployee_id: employeeID, Employeename: employeeName})
// 	}

// 	var response = JsonResponse{Type: "success", Data: data}

// 	json.NewEncoder(w).Encode(response)
// }

// func printMessage(s string) {
// 	panic("unimplemented")
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

var db *sql.DB

func init() {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

type sandbox struct {
	Id         string `json:"id"`
	Employeeid string `json:"employee_id"`
	Name       string `json:"name"`
	Timestamp  string `json:"timestamp"`
}

func savedatabase(employee_id string, name string, timestamp string) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	  INSERT INTO card_access_log (employee_id, name, timestamp)
	  VALUES ($1, $2, $3)
	  RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, employee_id, name, timestamp).Scan(&id)
	if err != nil {
		panic(err)
	}

}

// var db *sql.DB
// func setupDB() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 	  "password=%s dbname=%s sslmode=disable",
// 	  host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 	  panic(err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 	  panic(err)
// 	}

// 	fmt.Println("Successfully connected!")
//   }

func retrieveRecord(c *fiber.Ctx) error {

	// // checks if the request is a "GET" request
	// // if r.Method != "GET" {
	// 	// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	// 	return
	// }
	// We assign the result to 'rows'
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rowsRs, _ := db.Query("SELECT * FROM public.card_access_log")

	// if err != nil {
	// 	// http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	// 	// return
	// }
	defer rowsRs.Close()

	// creates placeholder of the sandbox
	// snbs := make([]sandbox, 0)

	// we loop through the values of rows
	// for rowsRs.Next() {
	// 	snb := sandbox{}
	// 	err := rowsRs.Scan(&snb.id, &snb.employeeid, &snb.name, &snb.timestamp)
	// 	if err != nil {
	// 		log.Println(err)
	// 		// http.Error(w, http.StatusText(500), 500)
	// 		// return
	// 	}
	// 	snbs = append(snbs, snb)
	// }

	// if err = rowsRs.Err(); err != nil {
	// 	// http.Error(w, http.StatusText(500), 500)
	// 	// return
	// }

	// a := ""
	// // loop and display the result in the browser
	// for _, snb := range snbs {
	// 	a = a + fmt.Sprintf("%s %s %s %s\n", snb.id, snb.employeeid, snb.name, snb.timestamp)
	// }

	var sand sandbox

	var sandboxArray []sandbox

	for rowsRs.Next() {
		rowsRs.Scan(
			&sand.Id,
			&sand.Employeeid,
			&sand.Name,
			&sand.Timestamp,
		)
		sandboxArray = append(sandboxArray, sand)
	}
	defer rowsRs.Close()
	fmt.Println(sandboxArray)
	c.JSON(sandboxArray)
	return c.SendStatus(200)
}
