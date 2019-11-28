package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    //基本認證
    authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
        "Eric": "123456",
        "Andy": "556888",
    }))
    authorized.GET("/hello/:name/*action", func(c *gin.Context) {
        //一般取值
        name := c.Param("name")
        action := c.Param("action")
        //設定預設值
        firstname := c.DefaultQuery("firstname", "None")
        lastname := c.Query("lastname")
        /**
        c.JSON(200, gin.H{
            "message": "hello " + name,
        })
        **/
        c.JSON(http.StatusOK, gin.H{
            "name": name,
            "action": action,
            "firstname": firstname,
            "lastname": lastname,
        })
    })
    r.Run(":3000")
}
