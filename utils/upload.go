package utils

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

//上传文件 单文件
func Upload(uid string, req *http.Request) (err error, filePath string, re int) {
	re = 0
	if req.Method != "POST" {
		re = -1
		return err, filePath, re
	} else {
		req.ParseMultipartForm(32 << 20)
		uploadFile, handle, err := req.FormFile("postFile")
		if err != nil {
			re = -2
			return err, filePath, re
		}
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" && ext != ".gif" {
			re = -3
			return err, filePath, re
		}
		if !Exists("./upload/") {
			os.Mkdir("./upload/", 0777)
		}
		filePath = "./upload/" + uid + ext
		saveFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			re = -4
			return err, filePath, re
		}
		io.Copy(saveFile, uploadFile)
		defer uploadFile.Close()
		defer saveFile.Close()
		return err, filePath, re
	}
}

//上传多个文件
func UploadFiles(uploadFile multipart.File, fileName string) (err error, filePath string, re int) {
	re = 0
	if !Exists("./upload/") {
		os.Mkdir("./upload/", 0777)
	}
	filePath = "./upload/" + fileName
	saveFile, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		re = -4
		return err, filePath, re
	}
	io.Copy(saveFile, uploadFile)
	defer saveFile.Close()
	return err, filePath, re
}

//删除文件
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
