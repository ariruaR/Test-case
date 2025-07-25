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
        "/subs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Получить все подписки",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Создать подписку",
                "parameters": [
                    {
                        "description": "Данные подписки",
                        "name": "sub",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Subcription"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Subcription"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/subs/price/{userID}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Получить суммарную стоимость подписок",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Фильтры: название сервиса, даты",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/subs/{userID}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Получить подписки пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Subcription"
                            }
                        }
                    },
                    "502": {
                        "description": "Bad Gateway",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Изменить подписку пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Изменяемое поле",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ChangeData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "subscriptions"
                ],
                "summary": "Удалить подписку пользователя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ChangeData": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "main.Data": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "models.Subcription": {
            "type": "object"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Subscription API",
	Description:      "API для управления подписками и подсчёта их стоимости",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
