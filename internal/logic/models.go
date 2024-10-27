package logic

type CreateAccountParams struct {
	Email       string
	AccountName string
	Password    string
}

type CreateAccountResponse struct {
	ID    uint64
	Email string
}

type CreateSessionParams struct {
	AccountName string
	Password    string
}

type AccountResponse struct {
	ID          uint64 `json:"id"`
	Email       string `json:"email"`
	AccountName string `json:"account_name"`
}

type CreateSessionResponse struct {
	account      AccountResponse
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
