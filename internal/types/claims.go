package types

type Claims struct {
	Id   int    `json:"user_id"`
	Role string `json:"user_role"`
}

func NewClaims(user_id int, role string) *Claims {
	return &Claims{user_id, role}
}
