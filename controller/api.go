package controller

import (
    "encoding/json"
    "github.com/dsmart/gotest_rest/utils"
    "github.com/zenazn/goji/web"
    "net/http"
    "github.com/dsmart/gotest_rest/model"
    "fmt"
    "github.com/dsmart/gotest_rest/services"
    "strconv"
)


type Controller struct {
    utils.Controller
}

//SignUp
func (controller *Controller) SignUp(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data model.User
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println("error ",err)
        return nil, err
    }
    response, err := services.SignUp(data)

    if err != nil {
        return nil, err
    }
    return response, nil
}

//Login
func (controller *Controller) Login(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data map[string]interface{}
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    response, err := services.UserLogin(data)

    if err != nil {
        return nil, err
    }
    return response, nil
}

//CreateContact
func (controller *Controller) CreateContact(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data model.Contacts
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    response, err := services.CreateContact(data,controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}

//ReadContact
func (controller *Controller) ReadContact(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    contactId,err:= strconv.Atoi(c.URLParams["id"])
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    response, err := services.ReadContact(contactId,controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}

//UpdateContact
func (controller *Controller) UpdateContact(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data model.Contacts
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    response, err := services.UpdateContact(data,controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}

//DeleteContact
func (controller *Controller) DeleteContact(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    contactId,err:= strconv.Atoi(c.URLParams["id"])
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    response, err := services.DeleteContact(contactId,controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}


//ListContact
func (controller *Controller) ListContact(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    response, err := services.ListContacts(controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}

//Logout
func (controller *Controller) Logout(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data map[string]string
    err := decoder.Decode(&data)
    if err != nil {
        return nil, err
    }
    response, err := services.Logout(controller.GetUserContext(c))

    if err != nil {
        return nil, err
    }
    return response, nil
}