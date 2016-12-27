package gojraphql

const SCHEMA = `{
  "@schema": {
    "enumStatus": {
      "$type": "enum",
      "$values": [
        {
          "name": "active",
          "value": "A"
        },
        {
          "name": "suspended",
          "value": "I"
        },
        {
          "mame": "deleted",
          "value": "D"
        }
      ]
    },
    "userInfo": {
      "id": {
        "$type": "int!",
        "readonly": true
      },
      "firstName": "str!",
      "lastName": "str",
      "address": "$addressInfo",
      "status": "$enumStatus"
    },
    "addressInfo": {
      "id": "int!"
    }
  },
  "@query": {
    "listUsers": {
        "@return": {
        "arrayOf": "$userInfo"
        },
        "@args": {
        "limit": "int",
        "count": "int"
        }
    },
    "searchUsers": {
        "@return": {
        "arrayOf": "$userInfo"
        },
        "@args": {
        "name": "str!"
        }
    }
  },
  "@mutation": {
      "saveUser": {
        "@return": "$userInfo",
        "@args": "$userInfo"
      }
  }
}`
