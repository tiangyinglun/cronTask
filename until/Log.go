package until

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

var glog = logrus.New()

func init() {

	//defer file.Close()
}

func LogObj() *logrus.Logger {

	glog.Out = os.Stdout

	t := time.Now().String()
	collection_str := t[:10]
	collect := "Log_" + strings.Replace(collection_str, "-", "_", -1) + ".log"

	file, err := os.OpenFile("./logs/"+collect, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		glog.Out = file
	} else {
		glog.Info("Failed to log to file, using default stderr")
	}
	//defer file.Close()
	return glog
}
