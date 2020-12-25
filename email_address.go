package amply

import (
	"encoding/json"
	"errors"
	"strings"
)

type EmailAddress struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func NewEmailAddress(data interface{}) (EmailAddress, error) {
	var email EmailAddress

	switch data.(type) {
	case string:
		email = fromString(data.(string))
	case EmailAddress:
		email = data.(EmailAddress)
	case map[string]interface{}, map[string]string:
		toJson, _ := json.Marshal(data)
		dataMap := make(map[string]string, 0)
		json.Unmarshal(toJson, &dataMap)

		email = EmailAddress{dataMap["name"], dataMap["email"]}
	default:
		return EmailAddress{}, errors.New("Expect `string`, `amply.EmailAddress`, or `map[string]string` for email address")
	}

	if email.Email == "" {
		return EmailAddress{}, errors.New("Must provide `email`")
	}

	return email, nil
}

func fromString(data string) EmailAddress {
	index := indexOf('<', []rune(data))

	if index == -1 {
		return EmailAddress{Email: data}
	}

	name := data[0:index]
	email := strings.Trim(data[index:len(data)], "<> ")

	return EmailAddress{name, email}
}

func indexOf(element rune, data []rune) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}

	return -1
}
