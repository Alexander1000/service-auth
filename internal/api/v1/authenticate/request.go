package authenticate

import "github.com/Alexander1000/service-auth/internal/model"

type request struct {
	Credential model.Credential
	Password string
}
