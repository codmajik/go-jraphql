package gojraphql

import (
	"bytes"
	"testing"
)

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
      "id": { "$type": "int!" },
      "fullName": "str#",
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
    "allFriends": {
      "@return": {
        "friendCount": "int",
        "friends":["$userInfo"],
        "lastUpdated": "str"
      }
    },
    "me": {
      "@return": "$userInfo"
    },
    "allUsers": {
        "@return": ["$userInfo", "int"],
        "@args": {
        "limit": "int",
        "count": "int"
        }
    },
    "searchUsers": {
        "@return": ["$userInfo"],
        "@args": {
        "name": "str!"
        }
    }
  },
  "@mutation": {
      "saveUser": {
        "@return": "$userInfo",
        "@args": "$userInfo"
      },
      "updateUser": {
        "@return": "$userInfo",
        "@args": {
          "id": "str!",
          "values":"$userInfo"
        }
      }
  }
}`

func TestSchema(t *testing.T) {
	s, err := NewSchema(bytes.NewBufferString(SCHEMA))
	println(s, err)
}
