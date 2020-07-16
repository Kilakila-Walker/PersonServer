package api

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

// @Tags ExaFileUploadAndDownload
// @Summary 断点续传到服务器
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "an example for breakpoint resume, 断点续传示例"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /fileUploadAndDownload/breakpointContinue [post]
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

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "Find the file, 查找文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/findFile [post]
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

// @Tags ExaFileUploadAndDownload
// @Summary 查找文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件完成"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"file uploaded, 文件创建成功"}"
// @Router /fileUploadAndDownload/findFile [post]
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

// @Tags ExaFileUploadAndDownload
// @Summary 删除切片
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "删除缓存切片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查找成功"}"
// @Router /fileUploadAndDownload/removeChunk [post]
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
