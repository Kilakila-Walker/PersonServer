package api

//断点续传API （暂时未搞懂是干嘛的，或许之后做传输的时候回用到)
import (
	"fmt"
	"io/ioutil"
	"perServer/global/response"
	resp "perServer/model/response"
	"perServer/service"
	"perServer/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 断点续传到服务器
func BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, err.Error(), c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, err.Error(), c)
		return
	}
	defer f.Close()
	cen, _ := ioutil.ReadAll(f)
	if flag := utils.CheckMd5(cen, chunkMd5); !flag {
		return
	}
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, err.Error(), c)
		return
	}
	err, pathc := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, err.Error(), c)
		return
	}

	if err = service.CreateFileChunk(file.ID, pathc, chunkNumber); err != nil {
		response.ToJson(response.ERROR, gin.H{}, err.Error(), c)
		return
	}
	response.ToJson(response.SUCCESS, gin.H{}, "切片创建成功", c)
}

// 查找文件
func FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	err, file := service.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		response.ToJson(response.ERROR, gin.H{}, "查找失败", c)
	} else {
		response.ToJson(response.SUCCESS, resp.FileResponse{File: file}, "成功", c)
	}
}

// 查找文件
func BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	err, filePath := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		response.ToJson(response.ERROR, resp.FilePathResponse{FilePath: filePath}, fmt.Sprintf("文件创建失败：%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, resp.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// 删除切片
func RemoveChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath := c.Query("filePath")
	err := utils.RemoveChunk(fileMd5)
	err = service.DeleteFileChunk(fileMd5, fileName, filePath)
	if err != nil {
		response.ToJson(response.ERROR, resp.FilePathResponse{FilePath: filePath}, fmt.Sprintf("缓存切片删除失败：%v", err), c)
	} else {
		response.ToJson(response.SUCCESS, resp.FilePathResponse{FilePath: filePath}, "缓存切片删除成功", c)
	}
}
