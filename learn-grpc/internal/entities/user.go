package entities

import "time"

type BaseUser struct {
	ID   string `json:"user_id"`
	Name string `json:"username"`
}

///////////////////////////////////////////////////////////
// User Model
// INFO: this model will be used for response API / GQL
///////////////////////////////////////////////////////////

type User struct {
	BaseUser
	RegisteredDate   time.Time `json:"registered_date"`
	LastActivityDate time.Time `json:"last_activity_date"`
	Avatar           Asset     `json:"avatar"`
}

///////////////////////////////////////////////////////////
// User Request Model
// INFO: this model will be used for parameters API / GQL
///////////////////////////////////////////////////////////

type UserRequest struct {
	BaseUser
	Avatar string `json:"avatar"`
}
