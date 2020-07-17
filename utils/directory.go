package utils

import (
	"fmt"
	"os"
)

// 文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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

//不存在则创建文件夹
func NoExistCreateDir(path string) (bool, error) {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true, err
		} else {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				fmt.Printf("create directory %s error: %s", path, err)
				return false, err
			}
			return true, err
		}
	}
	return false, err
}

//创建文件
func CreateFile(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		_, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("create file %s error: %s", filePath, err)
			return false, err
		}
		return true, err
	}
	return false, err
}

//删除文件或文件夹
func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 批量创建文件夹
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				fmt.Printf("create directory %s error: %s", v, err)
			}
		}
	}
	return err
}
