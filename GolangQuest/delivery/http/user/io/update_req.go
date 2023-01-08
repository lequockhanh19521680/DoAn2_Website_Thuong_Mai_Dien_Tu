package io

import "github.com/eNViDAT0001/Backend/internal/user/entities"

type UpdateUserReq struct {
	Username string  `json:"username"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   *bool   `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Type     entities.UserType
}
type AdminUpdateUserReq struct {
	Username string  `json:"username"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Gender   *bool   `json:"gender"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Type     entities.UserType
}
