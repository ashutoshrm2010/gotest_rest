package services

import (
    "golang.org/x/crypto/bcrypt"
    "time"
    "encoding/json"
    "github.com/dsmart/gotest_rest/model"
    "database/sql"
    "github.com/dsmart/gotest_rest/utils"
    "errors"
    "fmt"
)

func SignUp(userDetails model.User) ([]byte, error) {
    db, _ := utils.InitMysql()
    name := userDetails.Name
    if name == "" {
        return nil, errors.New("name is missing")
    }
    userName := userDetails.UserName
    if userName == "" {
        return nil, errors.New("username is missing")
    }

    email := userDetails.Email
    if email == "" {
        return nil, errors.New("emailId is missing")
    }
    password := userDetails.Password
    if password == "" {
        return nil, errors.New("password is missing")
    }
    var user string
    if name != ""&&password != "" {
        err := db.QueryRow("SELECT UserName FROM" + " " + "user" + " " + "WHERE UserName=?", userName).Scan(&user)
        if err == nil {
            fmt.Println("err ",err)
            return nil, errors.New("user already exist")
        } else if err == sql.ErrNoRows{
            stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS" + " " + "user" + "(id int NOT NULL AUTO_INCREMENT, Name varchar(50) NOT NULL,UserName varchar(50) NOT NULL,Email varchar(50) NOT NULL," +
                "Password varchar(120) NOT NULL,CreatedOn varchar(100) NOT NULL, PRIMARY KEY (id));")
            if err != nil {
                fmt.Println(err.Error())
            }
            _, err = stmt.Exec()
            if err != nil {
                fmt.Println(err.Error())
            }

            hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
            if err != nil {
                fmt.Println(err)
                return nil, err
            }
            createdOn := time.Now().UTC()
            _, err = db.Exec("INSERT INTO" + " " + "user" + "(Name,UserName,Email,Password,CreatedOn) VALUES( ?, ?, ?, ?,?)", name, userName, email,hashedPassword,createdOn )
            if err != nil {
                fmt.Println(err)
                return nil, err
            }
            response := make(map[string]interface{})
            response["message"] = "User Created"
            finalResponse, _ := json.Marshal(response)
            return finalResponse, nil
        }


    }
    return nil, nil

}
