///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Secret Store",
    "title": "Secret Store",
    "version": "0.0.1"
  },
  "basePath": "/v1/secret",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "secret"
        ],
        "operationId": "getSecrets",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter based on tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An array of registered secrets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/secret"
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "addSecret",
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The created secret.",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{secretName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "getSecret",
        "responses": {
          "200": {
            "description": "The secret identified by the secretName",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "updateSecret",
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "The updated secret",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "secret"
        ],
        "operationId": "deleteSecret",
        "parameters": [
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deletion"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "name of the secret to operate on",
          "name": "secretName",
          "in": "path",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "description": "Error error",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "code",
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "function-error": {
          "description": "function error",
          "type": "object",
          "x-go-name": "FunctionError"
        },
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        },
        "user-error": {
          "description": "user error",
          "type": "object",
          "x-go-name": "UserError"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "secret": {
      "description": "Secret secret",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "description": "id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "ID",
          "readOnly": true
        },
        "kind": {
          "description": "kind",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Kind",
          "readOnly": true
        },
        "name": {
          "description": "name",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Name"
        },
        "secrets": {
          "description": "SecretValue secret value",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        },
        "tags": {
          "description": "tags",
          "type": "array",
          "items": {
            "$ref": "#/definitions/secretTagsItems"
          },
          "x-go-name": "Tags"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "secretTagsItems": {
      "description": "Tag tag",
      "type": "object",
      "properties": {
        "key": {
          "description": "key",
          "type": "string",
          "x-go-name": "Key"
        },
        "value": {
          "description": "value",
          "type": "string",
          "x-go-name": "Value"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Operations on secrets",
      "name": "secret"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "VMware Dispatch Secret Store",
    "title": "Secret Store",
    "version": "0.0.1"
  },
  "basePath": "/v1/secret",
  "paths": {
    "/": {
      "get": {
        "tags": [
          "secret"
        ],
        "operationId": "getSecrets",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter based on tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "An array of registered secrets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/secret"
              }
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "addSecret",
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "The created secret.",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "409": {
            "description": "Already Exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{secretName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "getSecret",
        "responses": {
          "200": {
            "description": "The secret identified by the secretName",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Standard error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "secret"
        ],
        "operationId": "updateSecret",
        "parameters": [
          {
            "name": "secret",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "201": {
            "description": "The updated secret",
            "schema": {
              "$ref": "#/definitions/secret"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "secret"
        ],
        "operationId": "deleteSecret",
        "parameters": [
          {
            "pattern": "^[\\w\\d\\-]+$",
            "type": "string",
            "name": "secretName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Successful deletion"
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "Resource Not Found if no secret exists with the given name",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "name of the secret to operate on",
          "name": "secretName",
          "in": "path",
          "required": true
        },
        {
          "type": "array",
          "items": {
            "type": "string"
          },
          "collectionFormat": "multi",
          "description": "Filter based on tags",
          "name": "tags",
          "in": "query"
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "description": "Error error",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "code",
          "type": "integer",
          "format": "int64",
          "x-go-name": "Code"
        },
        "function-error": {
          "description": "function error",
          "type": "object",
          "x-go-name": "FunctionError"
        },
        "message": {
          "description": "message",
          "type": "string",
          "x-go-name": "Message"
        },
        "user-error": {
          "description": "user error",
          "type": "object",
          "x-go-name": "UserError"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "secret": {
      "description": "Secret secret",
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "description": "id",
          "type": "string",
          "format": "uuid",
          "x-go-name": "ID",
          "readOnly": true
        },
        "kind": {
          "description": "kind",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Kind",
          "readOnly": true
        },
        "name": {
          "description": "name",
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$",
          "x-go-name": "Name"
        },
        "secrets": {
          "description": "SecretValue secret value",
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
        },
        "tags": {
          "description": "tags",
          "type": "array",
          "items": {
            "$ref": "#/definitions/secretTagsItems"
          },
          "x-go-name": "Tags"
        }
      },
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    },
    "secretTagsItems": {
      "description": "Tag tag",
      "type": "object",
      "properties": {
        "key": {
          "description": "key",
          "type": "string",
          "x-go-name": "Key"
        },
        "value": {
          "description": "value",
          "type": "string",
          "x-go-name": "Value"
        }
      },
      "x-go-gen-location": "models",
      "x-go-package": "github.com/vmware/dispatch/pkg/api/v1"
    }
  },
  "securityDefinitions": {
    "bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    },
    "cookie": {
      "description": "use cookies for authentication, when the user already logged in",
      "type": "apiKey",
      "name": "Cookie",
      "in": "header"
    }
  },
  "security": [
    {
      "cookie": []
    },
    {
      "bearer": []
    }
  ],
  "tags": [
    {
      "description": "Operations on secrets",
      "name": "secret"
    }
  ]
}`))
}
