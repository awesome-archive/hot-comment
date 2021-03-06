// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-10 10:46:56.664787 +0800 CST m=+0.123184401

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/comments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据歌曲ID获取评论列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "歌曲ID",
                        "name": "song_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/hot_comments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "热门评论，按点赞数排行",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/hot_songs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "获取热门歌曲，按评论数排行",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/search/artists": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "搜索歌手",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/search/comments": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "按关键词搜索评论",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/search/songs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "按关键词搜索歌曲",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/songs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "根据歌手ID获取歌曲列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page_num",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "歌手ID",
                        "name": "artist_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "47.99.131.182",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "云音乐歌曲、评论搜索API",
	Description: "使用Go、Gin、Elasticsearch开发的一个网易云音乐歌曲、评论搜索API，可以在web上点击发请求哦，能查到数据，不过当然不是全部的数据啦",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
