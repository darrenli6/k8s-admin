package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	SR_File_Max_Bytes = 1024 * 1024 * 2
)

func main() {
	router := gin.Default()
	//本地上传到服务器 csv格式，其他类似 读取内容
	router.GET("/upload", uploadFile)

	//下载文件 读取内容
	router.GET("/download/read", downloadReadFile)

	//下载文件 写内容
	router.GET("/download/write", downloadWriteFile)

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run()
}

func uploadFile(c *gin.Context) {
	rFile, err := c.FormFile("file")
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}

	if rFile.Size > SR_File_Max_Bytes {
		c.String(400, "文件大小超过2M")
		return
	}

	file, err := rFile.Open()
	if err != nil {
		c.String(400, "文件格式错误")
		return
	}
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()
		if err != nil {
			c.String(400, err.Error())
			return
		}
		//line 就是每一行的内容
		fmt.Println(line)
		//line[0] 就是第几列
		fmt.Println(line[0])
	}

}

func downloadReadFile(c *gin.Context) {
	//http下载地址 csv
	csvFileUrl := c.PostForm("file_name")
	res, err := http.Get(csvFileUrl)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	defer res.Body.Close()
	//读取csv
	reader := csv.NewReader(bufio.NewReader(res.Body))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.String(400, err.Error())
			return
		}
		//line 就是每一行的内容
		fmt.Println(line)
		//line[0] 就是第几列
		fmt.Println(line[0])
	}
}

func downloadWriteFile(c *gin.Context) {
	//写文件
	var filename = "./output1.csv"
	if !checkFileIsExist(filename) {
		file, err := os.Create(filename) //创建文件
		if err != nil {
			c.String(400, err.Error())
			return
		}
		buf := bufio.NewWriter(file) //创建新的 Writer 对象
		buf.WriteString("test")
		buf.Flush()
		defer file.Close()
	}
	//返回文件流
	c.File(filename)
}

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
