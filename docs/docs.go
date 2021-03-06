// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/genres": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"получение списка всех жанров\"",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка получения"
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"получение фильмов\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category",
                        "name": "category",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "genre",
                        "name": "category",
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
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка получения фильмов"
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"получение информации о фильме по Id\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "404": {
                        "description": "Фильм не найден"
                    }
                }
            }
        },
        "/movies/{id}/similar": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"получение списка похожих фильмов\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
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
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    },
                    "500": {
                        "description": "Ошибка получения"
                    }
                }
            }
        },
        "/movies/{id}/watch": {
            "post": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Установка у юзера статус просмотренно для фильма\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка простановки статуса"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Удаление у юзера статус просмотренно для фильма\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка удаления статуса"
                    }
                }
            }
        },
        "/subscriptions/:username": {
            "post": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Подписка на другого пользователя\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username пользователя",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка подписки"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Отписка от пользователя\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username пользователя",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка отписки"
                    }
                }
            }
        },
        "/user/:username": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Получение пользователя\"",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username пользователя",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка получения"
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Получение текущего пользователя\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user_id пользователя",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка разлогина"
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Обновление инфо о пользователе\"",
                "parameters": [
                    {
                        "description": "Инфа пользователя",
                        "name": "info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка обновления"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Создание пользователя\"",
                "parameters": [
                    {
                        "description": "Данные пользователя",
                        "name": "signapData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.signupData"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Все создалось"
                    },
                    "400": {
                        "description": "Ошибка введеных данных"
                    },
                    "500": {
                        "description": "Ошибка создания"
                    }
                }
            }
        },
        "/users/admin/:username": {
            "post": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Логин пользователя\"",
                "parameters": [
                    {
                        "description": "Данные пользователя для логина",
                        "name": "loginData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.loginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка логина"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "\"Разлогин пользователя\"",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "session_id пользователя",
                        "name": "session_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Все ок"
                    },
                    "500": {
                        "description": "Ошибка разлогина"
                    }
                }
            }
        }
    },
    "definitions": {
        "http.loginData": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "http.signupData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Actor": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "biography": {
                    "type": "string"
                },
                "birthdate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_liked": {
                    "type": "boolean"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MovieReference"
                    }
                },
                "movies_count": {
                    "type": "integer"
                },
                "movies_rating": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "origin": {
                    "type": "string"
                },
                "profession": {
                    "type": "string"
                }
            }
        },
        "models.ActorData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ActorData"
                    }
                },
                "artist": {
                    "type": "string"
                },
                "banner": {
                    "type": "string"
                },
                "budget": {
                    "type": "string"
                },
                "composer": {
                    "type": "string"
                },
                "country": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "director": {
                    "type": "string"
                },
                "duration": {
                    "type": "string"
                },
                "genre": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "id": {
                    "type": "string"
                },
                "is_watched": {
                    "type": "boolean"
                },
                "montage": {
                    "type": "string"
                },
                "operator": {
                    "type": "string"
                },
                "poster": {
                    "type": "string"
                },
                "producer": {
                    "type": "string"
                },
                "production_year": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "rating_count": {
                    "type": "integer"
                },
                "scriptwriter": {
                    "type": "string"
                },
                "slogan": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "trailer_preview": {
                    "type": "string"
                }
            }
        },
        "models.MovieReference": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "favorite_actors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Actor"
                    }
                },
                "movies_watched": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "reviews_number": {
                    "type": "integer"
                },
                "subscribers": {
                    "type": "integer"
                },
                "subscriptions": {
                    "type": "integer"
                },
                "username": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}
