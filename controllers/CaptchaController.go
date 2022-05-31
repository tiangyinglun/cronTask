package controllers

import (
	"Gin/until"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
获取验证码
*/
func Captchimage(c *gin.Context) {
	Captcha := until.CreateCaptchaImage(4, 108, 36)
	data := until.Return(until.Success, Captcha, "")
	c.JSON(http.StatusOK, data)
}
