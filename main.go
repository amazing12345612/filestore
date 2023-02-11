package main

import (
	cmn "filestore/common"
	"filestore/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/gettoken", func(c *gin.Context) {
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")
		fmt.Println(username, password)
		atoken, rtoken, err := util.GenToken2(username, password)
		fmt.Println(atoken, rtoken)
		if err != nil {
			c.JSON(int(cmn.StatusServerError), gin.H{
				"error": err,
			})
			return

		}
		c.JSON(200, gin.H{
			"atoken": atoken,
			"rtoken": rtoken,
		})
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "",
		})
	})
	r.Run(":8080")

}
