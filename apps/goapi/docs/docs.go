// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/connectfour": {
            "post": {
                "description": "Submit a connectfour board state and return the best move at the depth provided",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "solvers"
                ],
                "summary": "Connect Four solver",
                "parameters": [
                    {
                        "description": "Board State",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ConnectFourRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ConnectfourResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorReponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ConnectFourRequest": {
            "type": "object",
            "required": [
                "board",
                "depth",
                "empty_piece",
                "opp_piece",
                "computer_piece"
            ],
            "properties": {
                "board": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        ".....",
                        ".....",
                        ".....",
                        ".....",
                        "..X.."
                    ]
                },
                "depth": {
                    "type": "integer",
                    "example": 3
                },
                "empty_piece": {
                    "type": "string",
                    "example": "."
                },
                "opp_piece": {
                    "type": "string",
                    "example": "O"
                },
                "computer_piece": {
                    "type": "string",
                    "example": "X"
                }
            }
        },
        "main.ConnectfourResponse": {
            "type": "object",
            "properties": {
                "column": {
                    "type": "integer",
                    "example": 3
                },
                "value": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "main.ErrorReponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "NDSquared GOAPI",
	Description:      "Golang backend service for NDSquared",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
