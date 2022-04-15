package main

import (

    "database/sql"
    "fmt"
    "net/http"

    _ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

var err error

func home(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Hello Home Page!")

}



type users struct {

    Name     string

    address    string

    marks   int   

}


func connect() {

    db, err = sql.Open("mysql", "root:Anup1234@tcp(127.0.0.1:3306)/nov_data")

    if err != nil {

        fmt.Println(err.Error())

    } else {

        fmt.Println("successfully connected to database.")

    }

}

func insert(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "isnert endpoint got hit!")

  
    name := "anup rahangdale"

    address:="pune"
     
      marks:=87

   

    db.Exec("insert into student123 values(?,?,?)", name, address,marks)
    fmt.Println("data inserted.")

}

func main() {

    defer db.Close()

    connect()

    http.HandleFunc("/", home)

    http.HandleFunc("/insert", insert)


    fmt.Println("listening on port 8000")

    http.ListenAndServe(":8000", nil)

}
