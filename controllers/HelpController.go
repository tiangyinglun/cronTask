package controllers

import (
	"Gin/models"
	"Gin/until"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Help(c *gin.Context) {
	data := gin.H{}
	c.HTML(http.StatusOK, "help.html", data)
}

func RetsetIpLong(c *gin.Context) {
	ip := until.GetCmdIpLong()

	b, err := models.TaskIpLongUpdate(ip)
	if err != nil {
		fmt.Println(err)
	}

	data := until.Return(until.Success, b, "")
	c.JSON(http.StatusOK, data)
}
