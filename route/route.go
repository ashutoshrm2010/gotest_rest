package route

import (
    "github.com/dsmart/gotest_rest/utils"
    "github.com/dsmart/gotest_rest/controller"
    "github.com/zenazn/goji"
)


//PrepareRoutes
func PrepareRoutes(application *utils.Application) {

    //Signup
    goji.Post("/services/user/signup", application.Route(&controller.Controller{}, true,"SignUp"))

    //Login
    goji.Post("/services/user/login", application.Route(&controller.Controller{}, true,"Login"))

    //Create Contact
    goji.Post("/services/user/contact", application.Route(&controller.Controller{}, false,"CreateContact"))

    //Read Contact
    goji.Get("/services/user/contact/:id", application.Route(&controller.Controller{}, false,"ReadContact"))

    //Update Contact
    goji.Put("/services/user/contact", application.Route(&controller.Controller{}, false,"UpdateContact"))

    //Delete Contact
    goji.Delete("/services/user/contact/:id", application.Route(&controller.Controller{}, false,"DeleteContact"))

    //List All contacts
    goji.Get("/services/user/contacts/list", application.Route(&controller.Controller{}, false,"ListContact"))

    //Logout
    goji.Get("/services/user/logout", application.Route(&controller.Controller{}, false,"Logout"))


}

