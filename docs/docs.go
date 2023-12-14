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
        "/story": {
            "put": {
                "description": "更新历史记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "更新历史记录",
                "parameters": [
                    {
                        "description": "更新历史记录请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "删除历史记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "删除历史记录",
                "parameters": [
                    {
                        "description": "删除历史记录请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.DeleteStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/story/generate": {
            "post": {
                "description": "生成故事",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "生成故事",
                "parameters": [
                    {
                        "description": "生成故事请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.GenerateStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/story/list": {
            "post": {
                "description": "得到用户的故事列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "得到用户的故事列表",
                "parameters": [
                    {
                        "description": "故事列表请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ListStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/story/save": {
            "post": {
                "description": "创建历史记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "创建历史记录",
                "parameters": [
                    {
                        "description": "创建历史记录请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/story/select": {
            "post": {
                "description": "根据mood分类历史记录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "历史记录操作"
                ],
                "summary": "根据mood分类历史记录",
                "parameters": [
                    {
                        "description": "分类历史记录请求体",
                        "name": "story",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SelectStoryReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/information": {
            "get": {
                "description": "得到用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录状态下用户操作"
                ],
                "summary": "得到用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "用户修改信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录状态下用户操作"
                ],
                "summary": "用户修改信息",
                "parameters": [
                    {
                        "description": "用户修改信息请求体",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserUpdateInfoReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "description": "用户进行登录操作",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserServiceReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/password": {
            "put": {
                "description": "用户修改密码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录状态下用户操作"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "description": "用户修改密码请求体",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserUpdatePwdReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/register": {
            "post": {
                "description": "注册一个新用户",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户操作"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户注册请求体",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UserServiceReq"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/userMenu/create": {
            "post": {
                "description": "添加彩蛋成就",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "彩蛋操作"
                ],
                "summary": "添加彩蛋成就",
                "parameters": [
                    {
                        "description": "添加彩蛋成就请求体",
                        "name": "userMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserMenuReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/userMenu/isMenu": {
            "post": {
                "description": "判断是否触发彩蛋成就",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "彩蛋操作"
                ],
                "summary": "判断是否触发彩蛋成就",
                "parameters": [
                    {
                        "description": "判断彩蛋成就请求体",
                        "name": "userMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.SelectMenuReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/userMenu/list": {
            "post": {
                "description": "得到彩蛋成就列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "彩蛋操作"
                ],
                "summary": "得到彩蛋成就列表",
                "parameters": [
                    {
                        "description": "彩蛋成就列表请求体",
                        "name": "userMenu",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ListUserMenuReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "身份验证令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "types.CreateStoryReq": {
            "type": "object",
            "required": [
                "content",
                "keywords",
                "mood",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "example": "content1"
                },
                "keywords": {
                    "type": "string",
                    "example": "室友+电脑"
                },
                "mood": {
                    "type": "string",
                    "example": "开心"
                },
                "title": {
                    "type": "string",
                    "example": "story1"
                }
            }
        },
        "types.CreateUserMenuReq": {
            "type": "object",
            "required": [
                "content",
                "keywords"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "example": "彩蛋:焦虑的夜晚"
                },
                "keywords": {
                    "type": "string",
                    "example": "大学生+未完成的作业"
                }
            }
        },
        "types.DeleteStoryReq": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "example": "story1"
                }
            }
        },
        "types.GenerateStoryReq": {
            "type": "object",
            "required": [
                "mood"
            ],
            "properties": {
                "keywords": {
                    "type": "string",
                    "example": "室友+电脑"
                },
                "mood": {
                    "type": "string",
                    "example": "开心"
                }
            }
        },
        "types.ListStoryReq": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "types.ListUserMenuReq": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "types.SelectMenuReq": {
            "type": "object",
            "required": [
                "keywords"
            ],
            "properties": {
                "keywords": {
                    "type": "string",
                    "example": "大学生+未完成的作业"
                }
            }
        },
        "types.SelectStoryReq": {
            "type": "object",
            "required": [
                "mood"
            ],
            "properties": {
                "mood": {
                    "type": "string",
                    "example": "开心"
                }
            }
        },
        "types.UpdateStoryReq": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "title": {
                    "type": "string",
                    "example": "story1"
                },
                "update_content": {
                    "type": "string",
                    "example": "content2"
                },
                "update_title": {
                    "type": "string",
                    "example": "story2"
                }
            }
        },
        "types.UserServiceReq": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "12345678"
                },
                "user_name": {
                    "type": "string",
                    "example": "john"
                }
            }
        },
        "types.UserUpdateInfoReq": {
            "type": "object",
            "properties": {
                "kitchen": {
                    "type": "string",
                    "example": "John Doe的厨房"
                },
                "update_name": {
                    "type": "string",
                    "example": "John Doe"
                }
            }
        },
        "types.UserUpdatePwdReq": {
            "type": "object",
            "required": [
                "originPwd",
                "updatePwd"
            ],
            "properties": {
                "originPwd": {
                    "type": "string",
                    "example": "12345678"
                },
                "updatePwd": {
                    "type": "string",
                    "example": "123456789"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Story-Cook",
	Description:      "Story-Cook API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
