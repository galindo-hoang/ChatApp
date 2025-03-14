package logic

type CreateAccountParams struct {
	Email       string `json:"email"`
	AccountName string `json:"account_name"`
	Password    string `json:"password"`
}

type CreateAccountResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type CreateSessionParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountResponse struct {
	ID          uint64 `json:"id"`
	Email       string `json:"email"`
	AccountName string `json:"account_name"`
}

type CreateSessionResponse struct {
	Account      AccountResponse `json:"account"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
}
