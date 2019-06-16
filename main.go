package main

import (
    "github.com/dsmart/gotest_rest/utils"
    "github.com/zenazn/goji"
    "github.com/dsmart/gotest_rest/route"
    "flag"
)

func main()  {
    var application = &utils.Application{}
    goji.Use(application.ApplyAuth)
    route.PrepareRoutes(application)
    flag.Set("bind","localhost:8080")
    goji.Serve()

}