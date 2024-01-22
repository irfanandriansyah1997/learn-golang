package enytities

type User struct {
	ID               uint   `json:"user_id"`
	Name             string `json:"username"`
	RegisteredDate   Date   `json:"registered_dated"`
	LastActivityDate Date   `json:"last_activity_dated"`
	Avatar           Asset  `json:"avatar"`
}
