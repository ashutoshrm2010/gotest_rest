package utils

import (
      "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

type ConnectionContext struct {
    MySqlConnection *sql.DB
}

var ApplicationContext ConnectionContext


func InitMysql() (db *sql.DB, err error) {

    database := DATABASE_DSMART
    username := "root"
    password := "password"

    db, err = sql.Open("mysql", username + ":" + password + "@/" + database)

    if err != nil {
        fmt.Println(err)
        panic(err.Error())
    }

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    return db, err

}

func PrepareConnection() {
    var err error
    ApplicationContext.MySqlConnection,err=InitMysql()

    if err!=nil{
        fmt.Println(err)
     log.Println("Error in DB connection")
    }

    log.Println("DB connected successfully")

}