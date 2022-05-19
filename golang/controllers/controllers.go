package controllers

import (
	"net/http"
	"shorten/middlewares"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var SavePath string

type Post_controller struct{}

func Ecoding_url(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
func Post_image(c *gin.Context) {
	a := GetInstance()
	a.UploadImage(c)
}

var once sync.Once
var Instance *Post_controller

func GetInstance() *Post_controller {
	once.Do(func() {
		Instance = &Post_controller{}
	})
	return Instance
}

func (Post_controller) UploadImage(c *gin.Context) {

	file, _ := c.FormFile("file")
	//name := c.PostForm("user_id")

	//filename := file.Filename
	//header := file.Header

	SavePath = "../Golang_AI/image/" + time.Now().Format("20060102150405") + file.Filename
	setPath(&SavePath)
	_ = c.SaveUploadedFile(file, SavePath)

	//fmt.Println(SavePath)

	//	c.SaveUploadedFile(file, "tmp/"+filename) // save file to tmp folder in current directory

	//filename := name + ".png"

	//fmt.Println(header)

	//Pass_data(c, SavePath)
	Test(c, SavePath)
	//c.Writer.WriteString(string(data))
}

/*
func Pass_data(c *gin.Context, SavePath string) {
	fmt.Println(SavePath)
	f := middlewares.Dump_to_python(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", f)
	fmt.Println("finish pass_data")
	SavePath = ""
}
*/
/*
func Pass_data_test(c *gin.Context) {
	SavePath = "./image/20220415204023WIN_20220216_08_47_44_Pro.jpg"
	f := middlewares.Dump_to_python(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", f)
	fmt.Println("finish pass_data")

}*/
func setPath(s *string) string {
	Path := *s
	return Path
}

func Test(c *gin.Context, SavePath string) {
	a := middlewares.Calc(SavePath)
	c.String(http.StatusOK, "您輸入的文字為: \n%s", a)

}

func Test_image(c *gin.Context) {
	SavePath = "./image/Cat_Sample_01.jpg"
	Test(c, SavePath)
}
