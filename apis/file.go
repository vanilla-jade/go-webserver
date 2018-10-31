package apis

import (
	"os"
	"io"
	"io/ioutil"
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	db "webServer/database"
	. "webServer/types"
)

func UploadFileApi(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://10.100.5.131:8080")
	c.Header("Access-Control-Allow-Methods", "POST")
	file, header, err := c.Request.FormFile("upload")
        filename := header.Filename
        out, err := os.Create("./uploadfile/" + filename)
        if err != nil {
                log.Fatal(err)
        }
        defer out.Close()
        _, err = io.Copy(out, file)
        if err != nil {
                log.Fatal(err)
        }
}

func WriteFileApi(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://10.100.5.131:8080")
	c.Header("Access-Control-Allow-Methods", "POST")
	info := c.Request.FormValue("info")
	filename := "./dailyfile/" + time.Now().Local().Format("2017_01_01") + ".txt"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	n, err := io.WriteString(f, info)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
			"msg": n,
		},
	)
	db.SqlDB.AutoMigrate(&Record{})
	db.SqlDB.Create(&Record{Msg: info, Time: time.Now()})
}

func ReadFileApi(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://10.100.5.131:8080")
	c.Header("Access-Control-Allow-Methods", "GET")
	filename := "./dailyfile/" + time.Now().Local().Format("2017_01_01") + ".txt"
	dat, err := ioutil.ReadFile(filename)
	var record Record
	db.SqlDB.First(&record, 1)
	if err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"success": false,
				"msg": "",
			},
		)
	}else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"success": true,
				"msg": string(dat) + record.Msg,
			},
		)
	}
}
