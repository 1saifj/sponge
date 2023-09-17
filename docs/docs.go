// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/userExample": {
            "post": {
                "description": "submit information to create userExample",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "create userExample",
                "parameters": [
                    {
                        "description": "userExample information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserExampleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CreateUserExampleRespond"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/condition": {
            "post": {
                "description": "get userExample by condition",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "get userExample by condition",
                "parameters": [
                    {
                        "description": "query condition",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Conditions"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUserExampleByConditionRespond"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/delete/ids": {
            "post": {
                "description": "delete userExamples by batch id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "delete userExamples",
                "parameters": [
                    {
                        "description": "id array",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUserExamplesByIDsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUserExamplesByIDsRespond"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/list": {
            "post": {
                "description": "list of userExamples by paging and conditions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "list of userExamples by query parameters",
                "parameters": [
                    {
                        "description": "query parameters",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_zhufuyi_sponge_internal_types.Params"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ListUserExamplesRespond"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/list/ids": {
            "post": {
                "description": "list of userExamples by batch id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "list of userExamples by batch id",
                "parameters": [
                    {
                        "description": "id array",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ListUserExamplesByIDsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ListUserExamplesByIDsRespond"
                        }
                    }
                }
            }
        },
        "/api/v1/userExample/{id}": {
            "get": {
                "description": "get userExample detail by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "get userExample detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.GetUserExampleByIDRespond"
                        }
                    }
                }
            },
            "put": {
                "description": "update userExample information by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "update userExample",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "userExample information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.UpdateUserExampleByIDRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.UpdateUserExampleByIDRespond"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete userExample by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userExample"
                ],
                "summary": "delete userExample",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.DeleteUserExampleByIDRespond"
                        }
                    }
                }
            }
        },
        "/codes": {
            "get": {
                "description": "list error codes info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "list error codes info",
                "responses": {}
            }
        },
        "/config": {
            "get": {
                "description": "show config info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "show config info",
                "responses": {}
            }
        },
        "/health": {
            "get": {
                "description": "check health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "check health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlerfunc.checkHealthResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "ping",
                "responses": {}
            }
        }
    },
    "definitions": {
        "github_com_zhufuyi_sponge_internal_types.Column": {
            "type": "object",
            "properties": {
                "exp": {
                    "description": "expressions, which default to = when the value is null, have =, !=, \u003e, \u003e=, \u003c, \u003c=, like",
                    "type": "string"
                },
                "logic": {
                    "description": "logical type, defaults to and when value is null, only \u0026(and), ||(or)",
                    "type": "string"
                },
                "name": {
                    "description": "column name",
                    "type": "string"
                },
                "value": {
                    "description": "column value"
                }
            }
        },
        "github_com_zhufuyi_sponge_internal_types.Params": {
            "type": "object",
            "properties": {
                "columns": {
                    "description": "query conditions",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_zhufuyi_sponge_internal_types.Column"
                    }
                },
                "page": {
                    "description": "page number, starting from page 0",
                    "type": "integer"
                },
                "size": {
                    "description": "lines per page",
                    "type": "integer"
                },
                "sort": {
                    "description": "sorted fields, multi-column sorting separated by commas",
                    "type": "string"
                }
            }
        },
        "handlerfunc.checkHealthResponse": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "types.Conditions": {
            "type": "object",
            "properties": {
                "columns": {
                    "description": "columns info",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_zhufuyi_sponge_internal_types.Column"
                    }
                }
            }
        },
        "types.CreateUserExampleRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "age",
                    "type": "integer"
                },
                "avatar": {
                    "description": "avatar",
                    "type": "string",
                    "minLength": 5
                },
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "gender": {
                    "description": "gender, 1:Male, 2:Female, other values:unknown",
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                },
                "name": {
                    "description": "username",
                    "type": "string",
                    "minLength": 2
                },
                "password": {
                    "description": "password",
                    "type": "string"
                },
                "phone": {
                    "description": "phone number, e164 rules, e.g. +8612345678901",
                    "type": "string"
                }
            }
        },
        "types.CreateUserExampleRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "id": {
                            "description": "id",
                            "type": "integer"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.DeleteUserExampleByIDRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.DeleteUserExamplesByIDsRequest": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "id list",
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "types.DeleteUserExamplesByIDsRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.GetUserExampleByConditionRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userExample": {
                            "$ref": "#/definitions/types.UserExampleObjDetail"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.GetUserExampleByIDRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userExample": {
                            "$ref": "#/definitions/types.UserExampleObjDetail"
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.ListUserExamplesByIDsRequest": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "id list",
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "types.ListUserExamplesByIDsRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userExamples": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UserExampleObjDetail"
                            }
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.ListUserExamplesRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data",
                    "type": "object",
                    "properties": {
                        "userExamples": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.UserExampleObjDetail"
                            }
                        }
                    }
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.UpdateUserExampleByIDRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "age",
                    "type": "integer"
                },
                "avatar": {
                    "description": "avatar",
                    "type": "string"
                },
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "gender": {
                    "description": "gender, 1:Male, 2:Female, other values:unknown",
                    "type": "integer"
                },
                "id": {
                    "description": "id",
                    "type": "integer"
                },
                "name": {
                    "description": "username",
                    "type": "string"
                },
                "password": {
                    "description": "password",
                    "type": "string"
                },
                "phone": {
                    "description": "phone number",
                    "type": "string"
                }
            }
        },
        "types.UpdateUserExampleByIDRespond": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "return code",
                    "type": "integer"
                },
                "data": {
                    "description": "return data"
                },
                "msg": {
                    "description": "return information description",
                    "type": "string"
                }
            }
        },
        "types.UserExampleObjDetail": {
            "type": "object",
            "properties": {
                "age": {
                    "description": "age",
                    "type": "integer"
                },
                "avatar": {
                    "description": "avatar",
                    "type": "string"
                },
                "createdAt": {
                    "description": "create time",
                    "type": "string"
                },
                "email": {
                    "description": "email",
                    "type": "string"
                },
                "gender": {
                    "description": "gender, 1:Male, 2:Female, other values:unknown",
                    "type": "integer"
                },
                "id": {
                    "description": "id",
                    "type": "string"
                },
                "loginAt": {
                    "description": "login timestamp",
                    "type": "integer"
                },
                "name": {
                    "description": "username",
                    "type": "string"
                },
                "phone": {
                    "description": "phone number",
                    "type": "string"
                },
                "status": {
                    "description": "account status, 1:inactive, 2:activated, 3:blocked",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "update time",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer your-jwt-token\" to Value",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{"http", "https"},
	Title:            "serverNameExample api docs",
	Description:      "http server api docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
