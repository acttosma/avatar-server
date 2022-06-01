// Package docs GENERATED BY SWAG; DO NOT EDIT
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
        "/account/changePwd": {
            "post": {
                "description": "Change the login password, remember user must be online when do this action, otherwise please see '/account/resetPassword'",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audience - Account Module"
                ],
                "summary": "User change the password while logged in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The JWT (called 'authorization' in the return value) after user logged in",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ChangePassword",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Base"
                        }
                    }
                }
            }
        },
        "/account/changeTradePwd": {
            "post": {
                "description": "Change the trade password, remember user must be online when do this action",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audience - Account Module"
                ],
                "summary": "User change the trade password while logged in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The JWT (called 'authorization' in the return value) after user logged in",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "ChangePassword",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Base"
                        }
                    }
                }
            }
        },
        "/account/login": {
            "post": {
                "description": "User login action  with resp.ActLogin returned",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audience - Account Module"
                ],
                "summary": "User login with mail and password",
                "parameters": [
                    {
                        "description": "ActLogin",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ActLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ActLogin"
                        }
                    }
                }
            }
        },
        "/account/register": {
            "post": {
                "description": "Provide mail,password and invite code to register as a new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audience - Account Module"
                ],
                "summary": "User register as an account through this api",
                "parameters": [
                    {
                        "description": "ActRegister",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.ActRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.ActRegister"
                        }
                    }
                }
            }
        },
        "/account/setTradePwd": {
            "post": {
                "description": "Set the trade password, remember user must be online when do this action",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Audience - Account Module"
                ],
                "summary": "User set the trade password while logged in",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The JWT (called 'authorization' in the return value) after user logged in",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "SetPassword",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.SetPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.Base"
                        }
                    }
                }
            }
        },
        "/captcha/check": {
            "get": {
                "description": "This interface checks the captcha code (called 'cc')",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Utilities-Captcha Module"
                ],
                "summary": "Check the captchaCode",
                "parameters": [
                    {
                        "type": "string",
                        "description": "the id of the catpcha",
                        "name": "ci",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the letters on the captcha image, case insensitive",
                        "name": "cc",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.CheckCaptcha"
                        }
                    }
                }
            }
        },
        "/captcha/get": {
            "get": {
                "description": "This interface returns an image data with the given parameters, the size of the request should be considered, can NOT be too high or too low.",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Utilities-Captcha Module"
                ],
                "summary": "Get a captcha image",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "the width of the image, default value is 100, max value should no more than 500. unit:px",
                        "name": "w",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "the heigh of the image, default value is 30, max value should no more than 100. unit:px",
                        "name": "h",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "the length of the letters shown on the image, default value is 4, it should be in the zone of [4,8]",
                        "name": "l",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.CaptchaGet"
                        }
                    }
                }
            }
        },
        "/qrcode/gen": {
            "get": {
                "description": "Generates a QRCode image with the given text",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Utilities-QRCode Module"
                ],
                "summary": "QRCode generator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The JWT (called 'authorization' in the return value) after user logged in",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the length of the square holding the QRCode image",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "the text of the QRCode image. if it is a http-url, it should be encoded with 'url-encode'",
                        "name": "text",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/resp.QRCodeImg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "req.ActLogin": {
            "type": "object",
            "required": [
                "mail",
                "password"
            ],
            "properties": {
                "mail": {
                    "description": "E-Mail address [required]",
                    "type": "string"
                },
                "password": {
                    "description": "Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]",
                    "type": "string"
                }
            }
        },
        "req.ActRegister": {
            "type": "object",
            "required": [
                "inviteCode",
                "mail",
                "password"
            ],
            "properties": {
                "inviteCode": {
                    "description": "invite code [required]",
                    "type": "string"
                },
                "mail": {
                    "description": "E-Mail address [required]",
                    "type": "string"
                },
                "password": {
                    "description": "Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]",
                    "type": "string"
                }
            }
        },
        "req.ChangePassword": {
            "type": "object",
            "required": [
                "password",
                "prePassword"
            ],
            "properties": {
                "password": {
                    "description": "Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]",
                    "type": "string"
                },
                "prePassword": {
                    "description": "PreviousPassword, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]",
                    "type": "string"
                }
            }
        },
        "req.SetPassword": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "password": {
                    "description": "Password, for safety reason,the request should mix this parameter with encryption,such as MD5 and SHA256. [required]",
                    "type": "string"
                }
            }
        },
        "resp.ActLogin": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "The ID of the account",
                    "type": "integer"
                },
                "authorization": {
                    "description": "The authorization token for the account",
                    "type": "string"
                }
            }
        },
        "resp.ActRegister": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "The ID of the account",
                    "type": "integer"
                },
                "authorization": {
                    "description": "The authorization token for the account",
                    "type": "string"
                }
            }
        },
        "resp.Base": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "the return code, 200 means OK, other codes mean failed",
                    "type": "integer"
                },
                "msg": {
                    "description": "The simple description of the code, the request should NOT use this value directly, it must be translated to another suitable message",
                    "type": "string"
                }
            }
        },
        "resp.CaptchaGet": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "The captcha image of request, in BASE64 format",
                    "type": "string"
                },
                "nonce": {
                    "description": "The nonce key of the captcha",
                    "type": "string"
                }
            }
        },
        "resp.CheckCaptcha": {
            "type": "object",
            "properties": {
                "captchaNonce": {
                    "description": "The nonce value of the captcha check response",
                    "type": "string"
                },
                "captchaNonceKey": {
                    "description": "The nonce key of the captcha check response",
                    "type": "string"
                }
            }
        },
        "resp.QRCodeImg": {
            "type": "object",
            "properties": {
                "image": {
                    "description": "The QRCode data, in BASE64 format",
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
