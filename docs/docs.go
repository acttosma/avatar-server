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
        "/account/login": {
            "post": {
                "description": "此接口用于在用户完成微信OAuth授权后,注册或登录使用,新用户完成注册,老用户直接登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "普通用户-账号相关"
                ],
                "summary": "用户完成微信OAuth授权后注册或登录使用",
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
        "/captcha/check": {
            "get": {
                "description": "接收图片验证码参数对输入的captcha对儿进行校验",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "不限用户-图片验证码相关"
                ],
                "summary": "图片验证码校验接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "图片验证码的ID(captchaId)",
                        "name": "ci",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "图片验证码显示的字符(captchaCode),大小写不敏感",
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
                "description": "根据给定的参数生成图片验证码,为提高安全性,不建议图片宽高太大或字符个数太少",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "不限用户-图片验证码相关"
                ],
                "summary": "图片验证码生成接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "图片验证码的宽,默认值100,最大值500,单位:px",
                        "name": "w",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "图片验证码的高,默认值30,最大值200,单位:px",
                        "name": "h",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "图片验证码包含字符的个数,范围[4,8],请根据宽高显示的实际效果来决定此参数,默认值4,单位:个",
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
                "description": "根据给定的text生成二维码图片",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "不限用户-二维码图片相关"
                ],
                "summary": "二维码图片生成接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户登录后返回的token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "需要生成的二维码的边长",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "二维码所承载的文字信息,如果是http-url,需要进行url-encode编码",
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
                "mobile",
                "openId",
                "tableId",
                "verifyCode"
            ],
            "properties": {
                "mobile": {
                    "description": "手机号码 [必填]",
                    "type": "string"
                },
                "openId": {
                    "description": "开放平台ID,支持微信,支付宝,微博等平台用户 [必填]",
                    "type": "string"
                },
                "tableId": {
                    "description": "桌台Id [必填]",
                    "type": "integer"
                },
                "verifyCode": {
                    "description": "手机验证码 [必填]",
                    "type": "string"
                }
            }
        },
        "resp.ActLogin": {
            "type": "object",
            "properties": {
                "accountId": {
                    "description": "账户ID 老用户情况下返回此值,新用户尚需绑定手机号码",
                    "type": "integer"
                },
                "isPwdSet": {
                    "description": "是否已经设置密码, true:已设置,fale:未设置",
                    "type": "boolean"
                },
                "openId": {
                    "description": "第三方开放ID",
                    "type": "string"
                },
                "token": {
                    "description": "令牌 老用户情况下返回此值,后续响应可携带鉴权用",
                    "type": "string"
                }
            }
        },
        "resp.CaptchaGet": {
            "type": "object",
            "properties": {
                "captcha": {
                    "description": "验证码图片,含图片bytes的BASE64编码图片,格式png",
                    "type": "string"
                },
                "nonce": {
                    "description": "图片验证码对应的唯一随机串",
                    "type": "string"
                }
            }
        },
        "resp.CheckCaptcha": {
            "type": "object",
            "properties": {
                "captchaNonce": {
                    "description": "图片验证码校验码,验证成功情况下会返回此值,单次有效,用于在需要使用图片验证码防止机刷接口调用",
                    "type": "string"
                },
                "captchaNonceKey": {
                    "description": "图片验证码校验码对应的Key",
                    "type": "string"
                }
            }
        },
        "resp.QRCodeImg": {
            "type": "object",
            "properties": {
                "image": {
                    "description": "二维码图片,图片bytes的BASE64编码图片,格式png",
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