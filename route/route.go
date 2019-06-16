package route

import (
    "github.com/dsmart/gotest_rest/utils"
    "github.com/dsmart/gotest_rest/controller"
    "github.com/zenazn/goji"
)

func PrepareRoutes(application *utils.Application) {
    goji.Post("/services/user/signup", application.Route(&controller.Controller{}, true,"SignUp"))
}

