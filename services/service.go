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
    "github.com/pborman/uuid"
    "log"
)

//SignUp--
func SignUp(userDetails model.User) ([]byte, error) {

    sessionDB:=utils.ApplicationContext.MySqlConnection
    err:=sessionDB.Ping()
    if err!=nil{
        log.Println("err in DB")
        return nil,err
    }
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
        err := sessionDB.QueryRow("SELECT UserName FROM" + " " + utils.TABLE_USER + " " + "WHERE UserName=?", userName).Scan(&user)
        if err == nil {
            return nil, errors.New("user already exist")
        } else if err == sql.ErrNoRows{
            stmt, err := sessionDB.Prepare("CREATE TABLE IF NOT EXISTS" + " " + utils.TABLE_USER + "(id int NOT NULL AUTO_INCREMENT, Name varchar(50) NOT NULL,UserName varchar(50) NOT NULL,Email varchar(50) NOT NULL," +
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
            _, err = sessionDB.Exec("INSERT INTO" + " " + utils.TABLE_USER + "(Name,UserName,Email,Password,CreatedOn) VALUES( ?, ?, ?, ?,?)", name, userName, email,hashedPassword,createdOn )
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

//UserLogin--
func UserLogin(userDetails map[string]interface{}) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()
    username := userDetails["username"].(string)
    password := userDetails["password"].(string)

    var databaseUsername string
    var databasePassword string

    err := db.QueryRow("SELECT UserName, Password FROM" + " " + utils.TABLE_USER + " " + "WHERE UserName=?", username).Scan(&databaseUsername, &databasePassword)
    if err != nil {
        return nil, errors.New("Invalid username or password")
    }

    err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))

    if err != nil {
        return nil, errors.New("Invalid username or password")
    } else {
        fmt.Println("login success")
        stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS"+" "+utils.TABLE_TOKEN+"(AccessToken varchar(100) NOT NULL, UserName varchar(50) NOT NULL,PRIMARY KEY (AccessToken));")
        if err != nil {
            fmt.Println(err.Error())
        }
        _, err = stmt.Exec()
        if err != nil {
            fmt.Println(err.Error())
        }
        var accesstoken string
        err = db.QueryRow("select AccessToken from" + " " + utils.TABLE_TOKEN + " " + "where UserName= ?", username).Scan(&accesstoken)
        if err == sql.ErrNoRows {
            accessToken := uuid.New()
            _, err = db.Exec("INSERT INTO" + " " + utils.TABLE_TOKEN+ "(accessToken,userName) VALUES(?, ?)", accessToken, username)
            if err != nil {
                return nil, err
            }
            response := make(map[string]interface{})
            response["accessToken"] = accessToken
            response["userName"] = username
            finalResponse, _ := json.Marshal(response)
            return finalResponse, nil

        } else {
            response := make(map[string]interface{})
            response["accessToken"] = accesstoken
            response["userName"] = username
            finalResponse, _ := json.Marshal(response)
            return finalResponse, nil
        }

    }

    return nil, nil
}

//CreateContact
func CreateContact(contact model.Contacts,userContext *utils.UserContext,) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()
    firstName := contact.FirstName
    lastName := contact.LastName
    if firstName == "" || lastName=="" {
        return nil, errors.New("name is missing")
    }
    email := contact.Email
    if email == "" {
        return nil, errors.New("email is missing")
    }

    mobile := contact.PhoneNumber
    if mobile == "" {
        return nil, errors.New("Phone is missing")
    }


    stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS" + " " + utils.TABLE_CONTACT + "(id int NOT NULL AUTO_INCREMENT, FirstName varchar(50) NOT NULL, LastName varchar(50) NOT NULL,Organization varchar(1000) NOT NULL,Phone varchar(50) NOT NULL," +
        "Email varchar(120) NOT NULL,Website varchar(120) NOT NULL,CreatedBy varchar(100) NOT NULL,CreatedOn varchar(100) NOT NULL, PRIMARY KEY (id));")
    if err != nil {
        fmt.Println(err.Error())
    }
    _, err = stmt.Exec()
    if err != nil {
        fmt.Println(err.Error())
    }
    _, err = db.Exec("INSERT INTO" + " " + utils.TABLE_CONTACT + "(FirstName,LastName,Organization,Phone,Email,Website,CreatedBy,CreatedOn) VALUES(?, ?,?,?,?,?,?,?)", firstName, lastName,contact.Organization,mobile,email,contact.Website,userContext.UserId,time.Now())
    if err != nil {
        return nil, err
    }

    response := make(map[string]interface{})
    response["message"] = "Contact successfully created"
    finalResponse, _ := json.Marshal(response)
    return finalResponse, nil
}

//ReadContact
func ReadContact(contactId int,userContext *utils.UserContext) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()
    var userData model.Contacts

    rows, err := db.Query("select * from" + " " + utils.TABLE_CONTACT + " " + " WHERE (id,CreatedBy)=(?,?)", contactId,userContext.UserId)
    if err != nil {
        fmt.Println("err ",err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Organization, &userData.PhoneNumber,&userData.Email,&userData.Email,&userData.CreatedBy, &userData.CreatedOn)
        if err != nil {
            fmt.Println("err ",err)
        }

    }
    response := make(map[string]interface{})
    response["contactDetail"] = userData
    finalResponse, _ := json.Marshal(response)
    return finalResponse, nil
}

//UpdateContact
func UpdateContact(contact model.Contacts,userContext *utils.UserContext,) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()

    stmt, err := db.Prepare("UPDATE" + " " + utils.TABLE_CONTACT + " " + "set FirstName=?,LastName=?,Organization=?,Phone=?,Email=?,Website=?,CreatedBy=?,CreatedOn=? WHERE (id,CreatedBy)=(?,?)")
    if err!=nil{
        return nil,err
    }

    _, err = stmt.Exec(contact.FirstName,contact.LastName,contact.Organization,contact.PhoneNumber,contact.Email,contact.Website,userContext.UserId,time.Now().UTC(), contact.ID,userContext.UserId)
    if err!=nil{
        return nil,err
    }
    response := make(map[string]interface{})
    response["message"] = "Contact successfully updated"
    finalResponse, _ := json.Marshal(response)
    return finalResponse, nil
}

//DeleteContact
func DeleteContact(contactId int,userContext *utils.UserContext) ([]byte, error) {
    db, _ := utils.InitMysql()
    stmt, err := db.Prepare("DELETE" + " from " + utils.TABLE_CONTACT + " " + "WHERE (id,CreatedBy)=(?,?)")
        _, err = stmt.Exec(contactId,userContext.UserId)
        if err != nil {
            fmt.Println(err)
        }
        response := make(map[string]interface{})

        response["message"] = "Contact Deleted successfully"
        finalResponse, _ := json.Marshal(response)
        return finalResponse, nil

}

//ListContacts
func ListContacts(userContext *utils.UserContext) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()
    var userData model.Contacts
    var userAllData []model.Contacts

    rows, err := db.Query("select * from" + " " + utils.TABLE_CONTACT + " " + "where createdBy = ?", userContext.UserId)
    if err != nil {
        fmt.Println("err ",err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&userData.ID, &userData.FirstName, &userData.LastName, &userData.Organization, &userData.PhoneNumber,&userData.Email,&userData.Email,&userData.CreatedBy, &userData.CreatedOn)
        if err != nil {
            fmt.Println("err ",err)
        }
        userAllData = append(userAllData, userData)

    }
    response := make(map[string]interface{})
    response["contactList"] = userAllData
    finalResponse, _ := json.Marshal(response)
    return finalResponse, nil
}

//Logout
func Logout(userContext *utils.UserContext) ([]byte, error) {
    db, _ := utils.InitMysql()
    defer db.Close()
    var databaseUsername string
    err := db.QueryRow("SELECT UserName FROM" + " " + utils.TABLE_TOKEN + " " + "WHERE UserName=?", userContext.UserName).Scan(&databaseUsername)
    if err != nil {
        return nil, errors.New("Invalid Token")
    }
    stmt, err := db.Prepare("DELETE FROM" + " " + utils.TABLE_TOKEN + " " + "WHERE UserName=?")
    if err != nil {
        fmt.Print(err.Error())
    }
    _, err = stmt.Exec(userContext.UserName)
    if err != nil {
        fmt.Println("accessToken Delete")
    }
    response := make(map[string]interface{})

    response["message"] = "Logout successful"
    finalResponse, _ := json.Marshal(response)
    return finalResponse, nil

}