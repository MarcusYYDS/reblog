// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "GPL-V3"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/login": {
            "post": {
                "description": "管理员使用用户名和密码进行登录，若登录成功，返回token",
                "tags": [
                    "站点管理"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名或邮箱",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "登录成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.AdminLoginResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "缺少必要参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "401": {
                        "description": "用户名或密码错误",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/admin/site": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "更新站点的名称、URL、描述和图标",
                "tags": [
                    "站点管理"
                ],
                "summary": "更新站点信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "站点名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点URL",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点图标(base64格式)",
                        "name": "icon",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功, 部分主题可能需重新部署生效",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "缺少参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/admin/tokenState": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取当前token的状态",
                "tags": [
                    "站点管理"
                ],
                "summary": "获取token状态",
                "responses": {
                    "200": {
                        "description": "若值为true则token有效",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "boolean"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/admin/userInfo": {
            "get": {
                "description": "获取管理员信息",
                "tags": [
                    "站点管理"
                ],
                "summary": "获取管理员信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.RespUserInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/article/list": {
            "get": {
                "description": "分页获取文章列表",
                "tags": [
                    "文章"
                ],
                "summary": "分页获取文章列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码, 默认为1",
                        "name": "pageIndex",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数量, 默认为10",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功返回文章列表",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.ArticlesListResp"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "参数不合法",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "500": {
                        "description": "服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/article/{slug}": {
            "get": {
                "description": "根据slug获取文章详情",
                "tags": [
                    "文章"
                ],
                "summary": "获取文章详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "缺少必要参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "404": {
                        "description": "未知的slug",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
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
                "description": "根据slug更新文章的标题、描述和内容",
                "tags": [
                    "文章"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章的slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章的标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章的描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章的内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新成功",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "缺失参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "404": {
                        "description": "未知的文章",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
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
                "description": "添加一篇新的文章",
                "tags": [
                    "文章"
                ],
                "summary": "添加文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章标题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文章内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "操作成功",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "缺少必要参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "409": {
                        "description": "slug已被其他文章使用",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
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
                "description": "根据slug删除文章",
                "tags": [
                    "文章"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "string",
                        "description": "文章的slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除成功",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "缺少必要参数",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "404": {
                        "description": "未知的文章",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/init": {
            "post": {
                "description": "使用给定的参数初始化站点",
                "tags": [
                    "站点管理"
                ],
                "summary": "初始化站点",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点名称",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点URL",
                        "name": "url",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "站点描述",
                        "name": "desc",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "站点图标(base64格式)",
                        "name": "icon",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "初始化成功",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "无效的邮箱或URL",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "403": {
                        "description": "此站点已初始化",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/rss": {
            "get": {
                "description": "获取包含所有文章的RSS",
                "produces": [
                    "text/xml"
                ],
                "tags": [
                    "Rss"
                ],
                "summary": "获取Rss",
                "responses": {
                    "200": {
                        "description": "RSS Feed"
                    },
                    "500": {
                        "description": "服务器错误",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/site": {
            "get": {
                "description": "获取站点信息",
                "tags": [
                    "站点"
                ],
                "summary": "获取站点信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.Site"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/{username}": {
            "put": {
                "description": "管理员更新用户信息",
                "tags": [
                    "站点管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "昵称",
                        "name": "nickname",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.Resp"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "获取reblog版本信息",
                "tags": [
                    "版本"
                ],
                "summary": "获取reblog版本信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Resp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.RespVersion"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Resp": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "handler.AdminLoginResp": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "handler.ArticlesListResp": {
            "type": "object",
            "properties": {
                "articles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "count": {
                    "type": "integer"
                }
            }
        },
        "handler.RespUserInfo": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handler.RespVersion": {
            "type": "object",
            "properties": {
                "app_name": {
                    "type": "string"
                },
                "commit": {
                    "type": "string"
                },
                "runtime": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Site": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "desc": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "reblog api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
