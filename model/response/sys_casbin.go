package response

import "perServer/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
