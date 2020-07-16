package response

import "perServer/model/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
