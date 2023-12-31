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
        "/all": {
            "get": {
                "tags": [
                    "File"
                ],
                "summary": "api for retrieve all files",
                "operationId": "GetAllFiles",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page Number",
                        "name": "page_number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Count of Files",
                        "name": "count",
                        "in": "query"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully retrieved all files",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.SingleFile"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Failed to retrieve all files",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/uploads/one": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "File"
                ],
                "summary": "api for upload single file",
                "operationId": "UploadSingleFile",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Upload one file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File Name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "File Description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully file uploaded",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid inputs",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to upload file",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.SingleFile": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "uploaded_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
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
