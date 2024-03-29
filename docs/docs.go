// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "soberkader@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/license/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/comments": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of all comment",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post details of comment corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Post details for a given Id",
                "parameters": [
                    {
                        "description": "create comment",
                        "name": "models.InputComment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputComment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of comment corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get details for a given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update the comment corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update comment identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be updated",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update comment",
                        "name": "models.InputComment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputComment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete the book corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete comment identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the comment to be deleted",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok"
                    }
                }
            }
        },
        "/photos": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of all photo",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post details of photo corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Post details for a given Id",
                "parameters": [
                    {
                        "description": "create photo",
                        "name": "models.InputPhoto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputPhoto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            }
        },
        "/photos/{photoId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of photo corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Get details for a given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update the photo corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Update photo identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be updated",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update photo",
                        "name": "models.InputPhoto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputPhoto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Photo"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete the book corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "photos"
                ],
                "summary": "Delete photo identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the photo to be deleted",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok"
                    }
                }
            }
        },
        "/socialmedias": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of all socialmedia",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "socialmedias"
                ],
                "summary": "Get details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Post details of socialmedia corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedias"
                ],
                "summary": "Post details for a given Id",
                "parameters": [
                    {
                        "description": "create socialmedia",
                        "name": "models.InputSocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            }
        },
        "/socialmedias/{socialmediaId}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get details of socialmedia corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedias"
                ],
                "summary": "Get details for a given id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialmedia",
                        "name": "socialmediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update the socialmedia corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedias"
                ],
                "summary": "Update socialmedia identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialmedia to be updated",
                        "name": "socialmediaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update socialmedia",
                        "name": "models.InputSocialMedia",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputSocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SocialMedia"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete the book corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "socialmedias"
                ],
                "summary": "Delete socialmedia identified by the given Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the socialmedia to be deleted",
                        "name": "socialmediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok"
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Post details of user corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Post details for a given Id",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "models.LoginUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Post details of user corresponding to the input Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Post details for a given Id",
                "parameters": [
                    {
                        "description": "create user",
                        "name": "models.InputUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.InputUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Comment": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "photo": {
                    "$ref": "#/definitions/models.Photo"
                },
                "photo_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.InputComment": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "photo_id": {
                    "type": "integer"
                }
            }
        },
        "models.InputPhoto": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.InputSocialMedia": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                }
            }
        },
        "models.InputUser": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.LoginUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Photo": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo_url": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.SocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "social_media_url": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Comment"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "photo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Photo"
                    }
                },
                "social_media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.SocialMedia"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "mygram-production-ab05.up.railway.app",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MyGram API",
	Description:      "This is a sample service for managing books",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
