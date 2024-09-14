package main

import (
	"fmt"
	"gin-example/config"
	"gin-example/model"
	"gin-example/router"

	"github.com/gin-gonic/gin"
)

func main() {

    //Config instance
    cfg := config.ReadConfig("./config.toml")


    err := model.InitDB(&cfg)

    if err != nil {
        panic(err)
    }

    fmt.Println("Server started on port 8080")

    //Gin instance
    r := gin.Default()
    router.SetRouter(r)
    r.Run() // listen and serve on
}
