package domain

// User is domain
type User struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
