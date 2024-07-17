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
        "/kittens": {
            "get": {
                "description": "Return list of kittens.",
                "summary": "Get All kittens.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "obejct"
                        }
                    }
                }
            },
            "post": {
                "description": "Save kittens data in Db.",
                "produces": [
                    "application/json"
                ],
                "summary": "Create kittens",
                "parameters": [
                    {
                        "description": "Create kittens",
                        "name": "kittens",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateKittensRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/kittens/{kittenID}": {
            "delete": {
                "description": "Remove kittens data by id.",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete kittens",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update kittens data.",
                "produces": [
                    "application/json"
                ],
                "summary": "Update kittens",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update kittens by id",
                        "name": "kittenId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update kittens",
                        "name": "kittens",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateKittensRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/kittens/{kittenId}": {
            "get": {
                "description": "Return the tahs whoes kittenId valu mathes id.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get Single kittens by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "update kittens by id",
                        "name": "kittenId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateKittensRequest": {
            "type": "object",
            "required": [
                "age",
                "breed",
                "color",
                "name",
                "price"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 200,
                    "minimum": 1
                },
                "breed": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                },
                "color": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                },
                "count": {
                    "type": "integer"
                },
                "in_stock": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "maxLength": 200,
                    "minLength": 1
                },
                "price": {
                    "type": "number",
                    "maximum": 200,
                    "minimum": 1
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Kitten Service API",
	Description:      "A Kitten service API in Go using Gin framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
