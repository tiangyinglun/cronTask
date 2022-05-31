package controllers

import (
	"Gin/models"
	"Gin/until"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Login(c *gin.Context) {

	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		pwd := c.PostForm("password")
		Captcha := c.PostForm("captcha")
		CaptchaId := c.PostForm("captchaId")
		status := until.VerifyCaptcha(CaptchaId, Captcha)
		fmt.Println(status)
		if !status {
			data := until.Return(until.CaptchaError, "", "")
			c.JSON(http.StatusOK, data)
			return
		}
		pwd_md5 := until.Md5(pwd)
		user, err := models.GetUserInfo(username, pwd_md5)
		fmt.Println(user)
		if err != nil {
			fmt.Println(err)
		}
		if username == user.Username && pwd_md5 == user.Password {
			Domain := ""
			host := c.Request.Host
			if strings.Contains(host, ":") {
				hosts := strings.Split(host, ":")
				Domain = hosts[0]
			} else {
				Domain = host
			}
			idstr := strconv.Itoa(user.Id)
			cookie_val := user.Name + "|" + idstr

			c.SetCookie("auth", cookie_val, 3600*12, "/", Domain, false, true)
			//c.Set("userinfo", user)
			data := until.Return(until.Success, "", "")
			c.JSON(http.StatusOK, data)
			return
		}
		data := until.Return(until.LoginError, "", "")
		c.JSON(http.StatusOK, data)
	}

	Hdata := gin.H{
		"time": time.Now().Unix(),
	}
	c.HTML(http.StatusOK, "login.html", Hdata)
}

func LoginOut(c *gin.Context) {
	Domain := ""
	host := c.Request.Host
	if strings.Contains(host, ":") {
		hosts := strings.Split(host, ":")
		Domain = hosts[0]
	} else {
		Domain = host
	}
	cookie_val := ""
	c.SetCookie("auth", cookie_val, -1, "/", Domain, false, true)
	c.Redirect(302, "login")
}

func UserLogin(c *gin.Context) {

	username := c.PostForm("username")
	pwd := c.PostForm("password")
	Captcha := c.PostForm("captcha")
	CaptchaId := c.PostForm("captchaId")

	status := until.VerifyCaptcha(CaptchaId, Captcha)
	fmt.Println(status)
	if !status {
		data := until.Return(until.CaptchaError, "", "")
		c.JSON(http.StatusOK, data)
		return
	}
	pwd_md5 := until.Md5(pwd)
	user, err := models.GetUserInfo(username, pwd_md5)
	fmt.Println(user)
	if err != nil {
		fmt.Println(err)
	}

	if username == user.Username && pwd_md5 == user.Password {
		Domain := ""
		host := c.Request.Host
		if strings.Contains(host, ":") {
			hosts := strings.Split(host, ":")
			Domain = hosts[0]
		} else {
			Domain = host
		}
		idstr := strconv.Itoa(user.Id)
		cookie_val := user.Name + "|" + idstr

		c.SetCookie("auth", cookie_val, 3600*8, "/", Domain, false, true)
		//c.Set("userinfo", user)
		data := until.Return(until.Success, "", "")
		c.JSON(http.StatusOK, data)
		return
	}
	data := until.Return(until.LoginError, "", "")
	c.JSON(http.StatusOK, data)
}
