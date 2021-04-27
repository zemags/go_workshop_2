package store

// TodoList contain main info for one todo
type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodUserListoList contain relation between user and list
type UserList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItem
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListsItem
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}
