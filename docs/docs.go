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
        "/cinema/available-seats": {
            "get": {
                "description": "Get available seats",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seating"
                ],
                "summary": "Get available seats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Seat"
                        }
                    }
                }
            }
        },
        "/cinema/cancel": {
            "post": {
                "description": "Cancel seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seating"
                ],
                "summary": "Cancel seat",
                "parameters": [
                    {
                        "description": "SeatAction",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SeatAction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SeatAction"
                        }
                    }
                }
            }
        },
        "/cinema/configure": {
            "post": {
                "description": "Configure cinema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seating"
                ],
                "summary": "Configure cinema",
                "parameters": [
                    {
                        "description": "Cinema",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Cinema"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Cinema"
                        }
                    }
                }
            }
        },
        "/cinema/reserve": {
            "post": {
                "description": "Reserve seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seating"
                ],
                "summary": "Reserve seat",
                "parameters": [
                    {
                        "description": "SeatAction",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.SeatAction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SeatAction"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Cinema": {
            "type": "object",
            "required": [
                "columns",
                "rows"
            ],
            "properties": {
                "columns": {
                    "type": "integer"
                },
                "minDistance": {
                    "type": "integer",
                    "default": 0
                },
                "rows": {
                    "type": "integer"
                },
                "seats": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/entity.Seat"
                        }
                    }
                }
            }
        },
        "entity.Seat": {
            "type": "object",
            "properties": {
                "column": {
                    "type": "integer"
                },
                "row": {
                    "type": "integer"
                },
                "status": {
                    "description": "Status 0: Available, 1: Reserved",
                    "allOf": [
                        {
                            "$ref": "#/definitions/entity.SeatStatus"
                        }
                    ]
                }
            }
        },
        "entity.SeatAction": {
            "type": "object",
            "properties": {
                "col": {
                    "description": "Col is required but can be 0",
                    "type": "integer",
                    "default": 0
                },
                "row": {
                    "description": "Row is required but can be 0",
                    "type": "integer",
                    "default": 0
                }
            }
        },
        "entity.SeatStatus": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-comments": {
                "Available": "0",
                "Reserved": "1"
            },
            "x-enum-varnames": [
                "Available",
                "Reserved"
            ]
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