package response

import "perServer/model"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File model.ExaFile `json:"file"`
}
