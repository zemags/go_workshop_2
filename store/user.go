package store

// User contain user-related info
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
