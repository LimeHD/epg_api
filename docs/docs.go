// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-04-03 21:24:46.8952935 +0300 MSK m=+0.080784401

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/channels": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show list of all channels",
                "operationId": "get-string-by-int",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entries.Channel"
                            }
                        }
                    }
                }
            }
        },
        "/channels/{id}/programm": {
            "get": {
                "description": "get string by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Show TV programm list",
                "operationId": "get-string-by-int",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Channel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entries.ProgrammResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entries.Channel": {
            "type": "object",
            "properties": {
                "aspectRatio": {
                    "type": "string"
                },
                "cdnvideo": {
                    "type": "integer"
                },
                "dayArchive": {
                    "type": "integer"
                },
                "descriptionEn": {
                    "type": "string"
                },
                "descriptionRu": {
                    "type": "string"
                },
                "foreignEpgId": {
                    "type": "integer"
                },
                "foreignUrl": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isForeign": {
                    "type": "integer"
                },
                "nameEn": {
                    "type": "string"
                },
                "nameRu": {
                    "type": "string"
                },
                "ourId": {
                    "type": "integer"
                },
                "playlistUrl": {
                    "type": "object",
                    "$ref": "#/definitions/entries.ChannelUrl"
                },
                "public": {
                    "type": "integer"
                },
                "tvprogram": {
                    "type": "integer"
                },
                "withArchive": {
                    "type": "integer"
                }
            }
        },
        "entries.ChannelUrl": {
            "type": "object",
            "properties": {
                "epgId": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "playlistOurId": {
                    "type": "integer"
                },
                "tz": {
                    "type": "integer"
                },
                "urlArchive": {
                    "type": "string"
                },
                "urlProtocol": {
                    "type": "string"
                },
                "urlSound": {
                    "type": "string"
                },
                "urlStuff": {
                    "type": "string"
                }
            }
        },
        "entries.Programm": {
            "type": "object",
            "properties": {
                "aspect_ratio": {
                    "type": "string"
                },
                "begin": {
                    "type": "integer"
                },
                "current": {
                    "type": "boolean"
                },
                "desc": {
                    "type": "string"
                },
                "end": {
                    "type": "integer"
                },
                "rating": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "entries.ProgrammResponse": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entries.Programm"
                    }
                },
                "name": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
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
	Host:        "epg.iptv2021.com",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a sample server celler server.",
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
