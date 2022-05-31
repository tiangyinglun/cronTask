package until

import (
	"Gin/models"
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"strconv"
	"strings"
	"time"
)

func Md5(str string) string {

	c := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", c)
}

/**
生成验证码
*/
func CreateCaptchaImage(len, width, height int) map[string]string {
	var image bytes.Buffer
	captchaId := captcha.NewLen(len)

	err := captcha.WriteImage(&image, captchaId, width, height)
	if err != nil {
		panic(err)
	}
	imageData := base64.StdEncoding.EncodeToString([]byte(image.String()))

	cap := make(map[string]string)
	cap["captchaId"] = captchaId
	cap["captchaPic"] = imageData
	return cap
}

//验证验证码
func VerifyCaptcha(captchaId string, digits string) bool {
	res := captcha.VerifyString(captchaId, digits)
	if res == false {
		return false
	}
	return true
}

/**
获取时间
*/
func GetNowTime() string {
	t := time.Now()
	nowTime := t.Format("2006-01-02 15:04:05")
	return nowTime
}

func TimestampString(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	nowTime := tm.Format("2006-01-02 15:04:05")
	return nowTime
}

/**
格式化
*/
func TimeLocateToFormat(str string) string {
	Tim, err := time.Parse("2006-01-02T15:04:05+08:00", str)
	if err != nil {
		fmt.Println(err)
	}
	return Tim.Format("2006-01-02 15:04:05")
}

/**
获取数据
*/
func GetStrInterfaceToString(c *gin.Context) (Adm models.AdminUsers, b bool) {
	var user models.AdminUsers
	Userinfo, bool := c.Get("userinfo")
	if !bool {
		return user, bool
	}

	var key string
	switch Userinfo.(type) {
	case string:
		key = Userinfo.(string)
		strByte := []byte(key)
		err := json.Unmarshal(strByte, &user)
		if err != nil {
			fmt.Println(err)
		}
	}
	return user, bool
}

func Ip2Long(ip string) (ips int64) {
	var ip_pieces = strings.Split(ip, ".")
	ip_1, _ := strconv.ParseInt(ip_pieces[0], 10, 32)
	ip_2, _ := strconv.ParseInt(ip_pieces[1], 10, 32)
	ip_3, _ := strconv.ParseInt(ip_pieces[2], 10, 32)
	ip_4, _ := strconv.ParseInt(ip_pieces[3], 10, 32)
	var ip_bin string = fmt.Sprintf("%08b%08b%08b%08b", ip_1, ip_2, ip_3, ip_4)
	ip_int, _ := strconv.ParseInt(ip_bin, 2, 64)
	return ip_int
}

func GetCmdIpLong() int {
	var ips []string
	netinter, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := recover(); err != nil {
			LogObj().WithFields(logrus.Fields{
				"func": "Run",
			}).Error("获取参数：", err)
			fmt.Println("获取host 失败")
			fmt.Println(err)
			return
		}

	}()
	for _, address := range netinter {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	ip := ""
	if len(ips) < 1 {
		ip = viper.GetString("app.host")
	} else {
		ip = ips[0]
	}
	ipInt := int(Ip2Long(ip))
	return ipInt
}
