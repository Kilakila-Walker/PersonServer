package request

import "perServer/model"

type SysDictionarySearch struct {
	model.Sys_Dic
	PageInfo
}
