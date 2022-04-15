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

    ID       int

    Name     string

    Email    string

    Password string

    AccType  string

}


func connect() {

    db, err = sql.Open("mysql", "root:1qaz2wsx@tcp(127.0.0.1:3306)/emp_data")

    if err != nil {

        fmt.Println(err.Error())

    } else {

        fmt.Println("successfully connected to database.")

    }

}





func insert(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "isnert endpoint got hit!")

  

    name := "Harshal Rahangde"

    id:=123
     
     salary:=456787

    

   

    db.Exec("insert into emp values(?,?,?)", id, name,salary)
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

