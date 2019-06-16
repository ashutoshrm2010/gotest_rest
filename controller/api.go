package controller

import (
    "encoding/json"
    "github.com/dsmart/gotest_rest/utils"
    "github.com/zenazn/goji/web"
    "net/http"
    "github.com/dsmart/gotest_rest/model"
    "fmt"
    "github.com/dsmart/gotest_rest/services"
)


type Controller struct {
    utils.Controller
}

func (controller *Controller) SignUp(c web.C, w http.ResponseWriter, r *http.Request) ([]byte, error) {
    decoder := json.NewDecoder(r.Body)
    var data model.User
    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println("error ",err)
        return nil, err
    }
    fmt.Println("data ",data)
    response, err := services.SignUp(data)

    if err != nil {
        return nil, err
    }
    return response, nil
}
