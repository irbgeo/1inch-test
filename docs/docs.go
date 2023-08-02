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
        "/get-amount-out": {
            "get": {
                "description": "Return outputAmount that corresponding uniswap_v2 pool will return if you try to swap inputAmount of fromToken in poolID",
                "tags": [
                    "requests"
                ],
                "summary": "get amount out",
                "parameters": [
                    {
                        "type": "string",
                        "default": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
                        "description": "from token address",
                        "name": "fromToken",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1e18",
                        "description": "amount for swapping",
                        "name": "inputAmount",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0xdac17f958d2ee523a2206206994597c13d831ec7",
                        "description": "to token address",
                        "name": "toToken",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852",
                        "description": "pool address",
                        "name": "poolID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "amountOut",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error description",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error description",
                        "schema": {
                            "type": "string"
                        }
                    }
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
