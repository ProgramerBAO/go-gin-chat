// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/index": {
            "get": {
                "description": "这是一个首页",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "首页"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "登陆校验",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "get": {
                "description": "用户模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "新建用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "paw",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "核对密码",
                        "name": "rePaw",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "PhoneNum",
                        "name": "phoneNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "description": "用户模块",
                "tags": [
                    "Users"
                ],
                "summary": "删除用户(逻辑删除)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUsers": {
            "get": {
                "description": "用户模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "获取用户列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "tags": [
                    "Users"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "PhoneNum",
                        "name": "phoneNum",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
