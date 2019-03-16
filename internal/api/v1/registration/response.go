package registration

type response struct {
	Result responseSuccess `json:"result"`
}

type responseSuccess struct {
	Success bool `json:"success"`
}
