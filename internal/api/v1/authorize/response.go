package authorize

type response struct {
	Result resultSuccess `json:"result"`
}

type resultSuccess struct {
	Success bool `json:"success"`
}
