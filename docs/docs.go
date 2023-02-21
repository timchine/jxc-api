// Code generated by swaggo/swag. DO NOT EDIT
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
        "/cargo_kind": {
            "put": {
                "description": "修改货品种类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改货品种类",
                "parameters": [
                    {
                        "description": "货品种类和属性",
                        "name": "货品种类",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReqAddCargoKind"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status 200 表示成功 否则提示msg内容",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "新增货品种类，新增种类时同时新增种类相关规格属性",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新增货品种类",
                "parameters": [
                    {
                        "description": "货品种类和属性",
                        "name": "货品种类",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReqAddCargoKind"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status 200 表示成功 否则提示msg内容",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/cargo_kind/_search": {
            "get": {
                "description": "通过货品code 或者 货品名称搜索 货品种类",
                "summary": "搜索货品种类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "货品code 或者 货品名称",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status 200 表示成功 否则提示msg内容",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ReqAddCargoKind"
                            }
                        }
                    }
                }
            }
        },
        "/cargo_kind/{ck_id}": {
            "get": {
                "description": "获取货品详情，获取货品详情和相关属性",
                "summary": "获取货品详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "货品种类ID",
                        "name": "ck_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status 200 表示成功 否则提示msg内容",
                        "schema": {
                            "$ref": "#/definitions/dto.ReqAddCargoKind"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除货品",
                "summary": "删除货品",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "货品种类ID",
                        "name": "ck_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "status 200 表示成功 否则提示msg内容",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "dto.ReqAddCargoKind": {
            "type": "object",
            "properties": {
                "cargo_attrs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CargoAttr"
                    }
                },
                "cargo_kind": {
                    "$ref": "#/definitions/model.CargoKind"
                }
            }
        },
        "model.CargoAttr": {
            "type": "object",
            "properties": {
                "attr_name": {
                    "description": "属性名称",
                    "type": "string"
                },
                "attr_value": {
                    "description": "属性值 ｜ 符号分割",
                    "type": "string"
                },
                "ca_id": {
                    "description": "属性ID",
                    "type": "integer"
                },
                "ck_id": {
                    "description": "关联货品种类",
                    "type": "integer"
                },
                "created_at": {
                    "type": "integer"
                },
                "status": {
                    "description": "状态 1 正常 8 删除",
                    "type": "integer"
                },
                "type": {
                    "description": "1 选择 2 文本",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        },
        "model.CargoKind": {
            "type": "object",
            "properties": {
                "cargo_attrs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CargoAttr"
                    }
                },
                "ck_code": {
                    "description": "货品编码",
                    "type": "string"
                },
                "ck_id": {
                    "type": "integer"
                },
                "ck_name": {
                    "description": "货品名称",
                    "type": "string"
                },
                "created_at": {
                    "type": "integer"
                },
                "intro": {
                    "description": "货品简介",
                    "type": "string"
                },
                "status": {
                    "description": "状态 1 正常 8 删除",
                    "type": "integer"
                },
                "type": {
                    "description": "1:原材料 2:半成品 3:成品",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/jxc",
	Schemes:          []string{},
	Title:            "进销存系统",
	Description:      "以实现无纸化办公为目标",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
