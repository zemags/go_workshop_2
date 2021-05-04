package store

// User contain user-related info
type User struct {
	ID int `json:"-" db:"id"`
	// binding:"required" - validate field data in request body
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
