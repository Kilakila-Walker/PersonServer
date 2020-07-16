// swagger(api测试用)

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/createApi": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "创建基础api",
                "parameters": [
                    {
                        "description": "创建api",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/deleteApi": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "删除指定api",
                "parameters": [
                    {
                        "description": "删除api",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getAllApis": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "获取所有的Api 不分页",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getApiById": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "根据id获取api",
                "parameters": [
                    {
                        "description": "根据id获取api",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.GetById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/getApiList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "分页获取API列表",
                "parameters": [
                    {
                        "description": "分页获取API列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.SearchApiParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/updateApi": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "创建基础api",
                "parameters": [
                    {
                        "description": "创建api",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysApi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/copyAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "拷贝角色",
                "parameters": [
                    {
                        "description": "拷贝角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/response.SysAuthorityCopyResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"拷贝成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/createAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "创建角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysAuthority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/deleteAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "description": "删除角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysAuthority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/getAuthorityList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "分页获取角色列表",
                "parameters": [
                    {
                        "description": "分页获取用户列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/setDataAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "设置角色资源权限",
                "parameters": [
                    {
                        "description": "设置角色资源权限",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysAuthority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"设置成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/updateAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authority"
                ],
                "summary": "设置角色资源权限",
                "parameters": [
                    {
                        "description": "设置角色资源权限",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysAuthority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"设置成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/autoCode/createTemp": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "自动代码模板",
                "parameters": [
                    {
                        "description": "创建自动代码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.AutoCodeStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/captcha": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "生成验证码",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/captcha/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "base"
                ],
                "summary": "生成验证码图片路径",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录接口",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.RegisterAndLoginStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"登陆成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户注册账号",
                "parameters": [
                    {
                        "description": "用户注册接口",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/casbin/CasbinTest": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "casbin"
                ],
                "summary": "casb RBAC RESTFUL测试路由",
                "parameters": [
                    {
                        "description": "获取权限列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.CasbinInReceive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/casbin/UpdateCasbin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "casbin"
                ],
                "summary": "更改角色api权限",
                "parameters": [
                    {
                        "description": "更改角色api权限",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.CasbinInReceive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/casbin/getPolicyPathByAuthorityId": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "casbin"
                ],
                "summary": "获取权限列表",
                "parameters": [
                    {
                        "description": "获取权限列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.CasbinInReceive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/customer": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "获取单一客户信息",
                "parameters": [
                    {
                        "description": "获取单一客户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "更新客户信息",
                "parameters": [
                    {
                        "description": "创建客户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "创建客户",
                "parameters": [
                    {
                        "description": "创建客户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "删除客户",
                "parameters": [
                    {
                        "description": "删除客户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/customer/customerList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysApi"
                ],
                "summary": "获取权限客户列表",
                "parameters": [
                    {
                        "description": "获取权限客户列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/breakpointContinue": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "断点续传到服务器",
                "parameters": [
                    {
                        "type": "file",
                        "description": "an example for breakpoint resume, 断点续传示例",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/deleteFile": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "删除文件",
                "parameters": [
                    {
                        "description": "传入文件里面id即可",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.ExaFileUploadAndDownload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/findFile": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "查找文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传文件完成",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"file uploaded, 文件创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/getFileList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "分页文件列表",
                "parameters": [
                    {
                        "description": "分页获取文件户列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/removeChunk": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "删除切片",
                "parameters": [
                    {
                        "type": "file",
                        "description": "删除缓存切片",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"查找成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/fileUploadAndDownload/upload": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaFileUploadAndDownload"
                ],
                "summary": "上传文件示例",
                "parameters": [
                    {
                        "type": "file",
                        "description": "上传文件示例",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/jwt/jsonInBlacklist": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jwt"
                ],
                "summary": "jwt加入黑名单",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"拉黑成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/GetMenuAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorityAndMenu"
                ],
                "summary": "获取指定角色menu",
                "parameters": [
                    {
                        "description": "增加menu和角色关联关系",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.AuthorityIdInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/addBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "新增菜单",
                "parameters": [
                    {
                        "description": "新增菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysBaseMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/addMenuAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorityAndMenu"
                ],
                "summary": "增加menu和角色关联关系",
                "parameters": [
                    {
                        "description": "增加menu和角色关联关系",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.AddMenuAuthorityInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/deleteBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "删除菜单",
                "parameters": [
                    {
                        "description": "删除菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.GetById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/getBaseMenuById": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "根据id获取菜单",
                "parameters": [
                    {
                        "description": "根据id获取菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.GetById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/getBaseMenuTree": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorityAndMenu"
                ],
                "summary": "获取用户动态路由",
                "parameters": [
                    {
                        "description": "可以什么都不填",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.RegisterAndLoginStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/getMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authorityAndMenu"
                ],
                "summary": "获取用户动态路由",
                "parameters": [
                    {
                        "description": "可以什么都不填",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.RegisterAndLoginStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/getMenuList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "分页获取基础menu列表",
                "parameters": [
                    {
                        "description": "分页获取基础menu列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/menu/updateBaseMenu": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "更新菜单",
                "parameters": [
                    {
                        "description": "更新菜单",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysBaseMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/system/ReloadSystem": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "设置配置文件内容",
                "parameters": [
                    {
                        "description": "设置配置文件内容",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.System"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/system/getSystemConfig": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "获取配置文件内容",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/system/setSystemConfig": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "设置配置文件内容",
                "parameters": [
                    {
                        "description": "设置配置文件内容",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.System"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"返回成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/changePassword": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "description": "用户修改密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.ChangePasswordStruct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "删除用户",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.SetUserAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "分页获取用户列表",
                "parameters": [
                    {
                        "description": "分页获取用户列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/setUserAuthority": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "设置用户权限",
                "parameters": [
                    {
                        "description": "设置用户权限",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/request.SetUserAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"修改成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/uploadHeaderImg": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "用户上传头像",
                "parameters": [
                    {
                        "type": "file",
                        "description": "用户上传头像",
                        "name": "headerImg",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户上传头像",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"上传成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/workflow/createWorkFlow": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workflow"
                ],
                "summary": "注册工作流",
                "parameters": [
                    {
                        "description": "注册工作流接口",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.SysWorkflow"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Captcha": {
            "type": "object",
            "properties": {
                "imgHeight": {
                    "type": "integer"
                },
                "imgWidth": {
                    "type": "integer"
                },
                "keyLong": {
                    "type": "integer"
                }
            }
        },
        "config.Casbin": {
            "type": "object",
            "properties": {
                "modelPath": {
                    "type": "string"
                }
            }
        },
        "config.JWT": {
            "type": "object",
            "properties": {
                "signingKey": {
                    "type": "string"
                }
            }
        },
        "config.Log": {
            "type": "object",
            "properties": {
                "file": {
                    "type": "string"
                },
                "logFile": {
                    "type": "boolean"
                },
                "prefix": {
                    "type": "string"
                },
                "stdout": {
                    "type": "string"
                }
            }
        },
        "config.Mysql": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "string"
                },
                "dbname": {
                    "type": "string"
                },
                "logMode": {
                    "type": "boolean"
                },
                "maxIdleConns": {
                    "type": "integer"
                },
                "maxOpenConns": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "config.Qiniu": {
            "type": "object",
            "properties": {
                "accessKey": {
                    "type": "string"
                },
                "secretKey": {
                    "type": "string"
                }
            }
        },
        "config.Redis": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "string"
                },
                "db": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "config.Server": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "object",
                    "$ref": "#/definitions/config.Captcha"
                },
                "casbin": {
                    "type": "object",
                    "$ref": "#/definitions/config.Casbin"
                },
                "jwt": {
                    "type": "object",
                    "$ref": "#/definitions/config.JWT"
                },
                "log": {
                    "type": "object",
                    "$ref": "#/definitions/config.Log"
                },
                "mysql": {
                    "type": "object",
                    "$ref": "#/definitions/config.Mysql"
                },
                "qiniu": {
                    "type": "object",
                    "$ref": "#/definitions/config.Qiniu"
                },
                "redis": {
                    "type": "object",
                    "$ref": "#/definitions/config.Redis"
                },
                "sqlite": {
                    "type": "object",
                    "$ref": "#/definitions/config.Sqlite"
                },
                "system": {
                    "type": "object",
                    "$ref": "#/definitions/config.System"
                }
            }
        },
        "config.Sqlite": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "string"
                },
                "logMode": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "config.System": {
            "type": "object",
            "properties": {
                "addr": {
                    "type": "integer"
                },
                "dbType": {
                    "type": "string"
                },
                "env": {
                    "type": "string"
                },
                "useMultipoint": {
                    "type": "boolean"
                }
            }
        },
        "model.AutoCodeStruct": {
            "type": "object",
            "properties": {
                "abbreviation": {
                    "type": "string"
                },
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Field"
                    }
                },
                "packageName": {
                    "type": "string"
                },
                "structName": {
                    "type": "string"
                }
            }
        },
        "model.ExaCustomer": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string"
                },
                "customerPhoneData": {
                    "type": "string"
                },
                "sysUser": {
                    "type": "object",
                    "$ref": "#/definitions/model.SysUser"
                },
                "sysUserAuthorityID": {
                    "type": "string"
                },
                "sysUserId": {
                    "type": "integer"
                }
            }
        },
        "model.ExaFileUploadAndDownload": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.Field": {
            "type": "object",
            "properties": {
                "columnName": {
                    "type": "string"
                },
                "fieldJson": {
                    "type": "string"
                },
                "fieldName": {
                    "type": "string"
                },
                "fieldType": {
                    "type": "string"
                }
            }
        },
        "model.SysApi": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "model.SysAuthority": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "authorityName": {
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "dataAuthorityId": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "deletedAt": {
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenu"
                    }
                },
                "parentId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.SysBaseMenu": {
            "type": "object",
            "properties": {
                "authoritys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenu"
                    }
                },
                "component": {
                    "type": "string"
                },
                "defaultMenu": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "icon": {
                    "type": "string"
                },
                "keepAlive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.SysUser": {
            "type": "object",
            "properties": {
                "authority": {
                    "type": "object",
                    "$ref": "#/definitions/model.SysAuthority"
                },
                "authorityId": {
                    "type": "string"
                },
                "headerImg": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "model.SysWorkflow": {
            "type": "object",
            "properties": {
                "workflowDescription": {
                    "description": "工作流描述",
                    "type": "string"
                },
                "workflowName": {
                    "description": "工作流英文id",
                    "type": "string"
                },
                "workflowNickName": {
                    "description": "工作流名称",
                    "type": "string"
                },
                "workflowStep": {
                    "description": "工作流步骤",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysWorkflowStepInfo"
                    }
                }
            }
        },
        "model.SysWorkflowStepInfo": {
            "type": "object",
            "properties": {
                "isEnd": {
                    "description": "是否是完结流节点",
                    "type": "boolean"
                },
                "isStrat": {
                    "description": "是否是开始流节点",
                    "type": "boolean"
                },
                "stepAuthorityID": {
                    "description": "操作者级别id",
                    "type": "string"
                },
                "stepName": {
                    "description": "工作流名称",
                    "type": "string"
                },
                "stepNo": {
                    "description": "步骤id （第几步）",
                    "type": "number"
                },
                "workflowID": {
                    "description": "所属工作流ID",
                    "type": "integer"
                }
            }
        },
        "model.System": {
            "type": "object",
            "properties": {
                "config": {
                    "type": "object",
                    "$ref": "#/definitions/config.Server"
                }
            }
        },
        "request.AddMenuAuthorityInfo": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenu"
                    }
                }
            }
        },
        "request.AuthorityIdInfo": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                }
            }
        },
        "request.CasbinInReceive": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "casbinInfos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.CasbinInfo"
                    }
                }
            }
        },
        "request.CasbinInfo": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "request.ChangePasswordStruct": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.GetById": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "number"
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        },
        "request.RegisterAndLoginStruct": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "captchaId": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.SearchApiParams": {
            "type": "object",
            "properties": {
                "apiGroup": {
                    "type": "string"
                },
                "desc": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "method": {
                    "type": "string"
                },
                "orderKey": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "request.SetUserAuth": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "response.SysAuthorityCopyResponse": {
            "type": "object",
            "properties": {
                "authority": {
                    "type": "object",
                    "$ref": "#/definitions/model.SysAuthority"
                },
                "oldAuthorityId": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample Server pets",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
