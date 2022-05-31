package controllers

import (
	"Gin/until"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	Captchimage := until.CreateCaptchaImage(4, 80, 36)
	//fmt.Println(Captchimage)
	data := gin.H{
		"title":   "name",
		"sb":      "good",
		"list":    "fgdg",
		"listsg":  "dsds",
		"capid":   Captchimage["id"],
		"capbase": Captchimage["base"],
	}
	fmt.Println(data)

	c.HTML(http.StatusOK, "index.html", data)
}

func GetName(c *gin.Context) {
	id := c.DefaultQuery("id", "枯藤")
	fmt.Println(id)
	fc := until.Md5(id)
	fmt.Println(fc)
	c.String(http.StatusOK, fc)
}

type Jobs struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Work string `json:"work"`
	Addr string `json:"addr"`
}

func Ujson(c *gin.Context) {

	var j Jobs
	j.Age = 12
	j.Name = "田英伦"
	j.Work = "sb"
	j.Addr = "家里蹲"

	//fc := models.GetOne()
	//fmt.Println(fc)

	c.JSON(http.StatusOK, j)
}

func Sjson(c *gin.Context) {
	id := c.DefaultQuery("id", "枯藤")
	fmt.Println(id)
	data := map[string]interface{}{
		"name": "GO name",
		"tag":  "good",
	}

	fmt.Println(c.Get("user_id"))

	c.JSON(http.StatusOK, data)
}
