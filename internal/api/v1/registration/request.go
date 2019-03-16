package registration

type request struct {
	UserID int64 `json:"userId"`
	Password string `json:"password"`
	Credentials []requestCredential `json:"credentials"`
}

type requestCredential struct {
	ID int64 `json:"id"`
	Type string `json:"type"`
}
