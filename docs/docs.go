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
        "/partners/:id": {
            "get": {
                "description": "This endpoint retrieves a partner using its UUID.",
                "tags": [
                    "Partners"
                ],
                "summary": "Get Partner",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Partner"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/partners/match": {
            "post": {
                "description": "This endpoint is used to show customers the available partners within their radius",
                "tags": [
                    "Partners"
                ],
                "summary": "Match Partners",
                "parameters": [
                    {
                        "description": "The match request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/partner.MatcherRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Partner"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Partner": {
            "type": "object",
            "properties": {
                "address_lat": {
                    "type": "number"
                },
                "address_long": {
                    "type": "number"
                },
                "flooring_materials": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "operating_radius": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "partner.MatcherRequest": {
            "type": "object",
            "properties": {
                "address_lat": {
                    "type": "number"
                },
                "address_long": {
                    "type": "number"
                },
                "floor_area": {
                    "type": "number"
                },
                "floor_material": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Partners API",
	Description:      "This service is responsible for managing partners.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
