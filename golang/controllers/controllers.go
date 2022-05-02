package controllers

import (
	"fmt"
	"net/http"
	"shorten/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

var SavePath string

func Ecoding_url(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

func UploadImage(c *gin.Context) {

	file, _ := c.FormFile("file")
	//name := c.PostForm("user_id")

	//filename := file.Filename
	//header := file.Header

	SavePath = "./image/" + time.Now().Format("20060102150405") + file.Filename
	setPath(&SavePath)
	_ = c.SaveUploadedFile(file, SavePath)

	//fmt.Println(SavePath)

	//	c.SaveUploadedFile(file, "tmp/"+filename) // save file to tmp folder in current directory

	//filename := name + ".png"

	//fmt.Println(header)

	Pass_data(c, SavePath)
	//c.Writer.WriteString(string(data))
}

func Pass_data(c *gin.Context, SavePath string) {
	fmt.Println(SavePath)
	f := middlewares.Dump_to_python(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", f)
	fmt.Println("finish pass_data")
	SavePath = ""
}

func Pass_data_test(c *gin.Context) {
	SavePath = "./image/20220415204023WIN_20220216_08_47_44_Pro.jpg"
	f := middlewares.Dump_to_python(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", f)
	fmt.Println("finish pass_data")

}
func setPath(s *string) string {
	Path := *s
	return Path
}

func Test(c *gin.Context) {
	a := middlewares.Calc()
	c.String(http.StatusOK, "您輸入的文字為: \n%s", a)

}
