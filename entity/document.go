package entity

type Document struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	UserID int64  `json:"-"`
	User   *User  `json:"user,omitempty"`
}
