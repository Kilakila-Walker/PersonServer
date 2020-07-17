package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

//上传文件 单文件 uid为新建文件名
func Upload(uid string, req *http.Request) (err error, filePath string, re int) {
	re = 0
	if req.Method != "POST" {
		re = -1 //非post上传
		return err, filePath, re
	} else {
		err = req.ParseMultipartForm(32 << 20)
		if err != nil {
			re = -5 //文件过大
			return err, filePath, re
		}
		uploadFile, handle, err := req.FormFile("postFile")
		if err != nil {
			re = -2 //获取文件失败
			return err, filePath, re
		}
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			re = -3 //文件格式不对
			return err, filePath, re
		}
		exist, err := NoExistCreateDir("../upload")
		if exist {
			re = -4 //创建文件失败
			return err, filePath, re
		}
		filePath = "../upload/" + uid + ext
		saveFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			re = -4 //创建文件失败
			return err, filePath, re
		}
		io.Copy(saveFile, uploadFile)
		defer uploadFile.Close()
		defer saveFile.Close()
		return err, filePath, re
	}
}

//上传多文件示例
func upldFiles(c *gin.Context) (err error) {
	uid := "uuid" //由uuid生成一个唯一码
	err = c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		return err
	}
	// 获取表单
	form := c.Request.MultipartForm
	// 获取参数upload后面的多个文件名，存放到数组files里面，
	files := form.File["upload"]
	// 遍历数组，每取出一个file就拷贝一次
	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return err
		}
		ext := strings.ToLower(path.Ext(files[i].Filename))
		filePath := "../upload/" + uid + ext
		UploadFiles(file, filePath)
		//然后把UploadFiles传回的地址信息写入数据库
	}
	return err
}

//上传多个文件
func UploadFiles(uploadFile multipart.File, fileName string) (err error, filePath string, re int) {
	re = 0
	if !Exists("../upload") {
		os.Mkdir("../upload", os.ModePerm)
	}
	filePath = "../upload/" + fileName
	saveFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		re = -4
		return err, filePath, re
	}
	io.Copy(saveFile, uploadFile)
	defer saveFile.Close()
	return err, filePath, re
}
