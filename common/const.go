package common

type TokenPayload struct {
	UId   int    `json:"user_id"`
	URole string `json:"role"`
}

func (p *TokenPayload) GetUserId() int {
	return p.UId
}

func (p *TokenPayload) GetRole() string {
	return p.URole
}
