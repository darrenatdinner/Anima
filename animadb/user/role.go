package user

import "projectanima/AnimaDB/domain"

// Role - A Single Role Of A User
type Role struct {
	domain.AuditData

	Name  string `json:"name,omitempty"`
	Level int    `json:"level,omitempty"`
}
