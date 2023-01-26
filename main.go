package main

import (
	"fmt"
	"net/http"

	"github.com/MouriMai/demo-youbride/batch"
	"github.com/MouriMai/demo-youbride/models"

	"github.com/MouriMai/demo-youbride/config/database"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Go project started!")
	// batch.AutoLogin()
	// models.SubProcess()
	db := database.New()
	fmt.Println(db)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/members", func(c *gin.Context) {
		var members []models.Member
		members = models.GetMembers()
		c.JSON(http.StatusOK, gin.H{
			"members": members,
		})
	})
	r.GET("/membersLastInfo", func(c *gin.Context) {
		membersLastInfo := models.GetMembersLastInfo()
		c.JSON(http.StatusOK, gin.H{
			"membersLastInfo": membersLastInfo,
		})
	})
	r.GET("/autologin", batch.AutoLogin)
	r.GET("/footprint", batch.Footprint)
	r.GET("/talk", batch.Talk)
	r.GET("/appapisecret", batch.Talk)
	r.GET("/registeremailtoken", batch.RegisterEmailToken)
	r.GET("/cachewarmer", batch.CacheWarmer)
	// {
	// 	results := models.AutoLogin()
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"results": results,
	// 	})
	// }

	r.Run(":3001")

}
