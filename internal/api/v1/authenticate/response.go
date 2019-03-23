package authenticate

import "github.com/Alexander1000/service-auth/internal/model"

type response struct {
	Result model.Token `json:"result"`
}
