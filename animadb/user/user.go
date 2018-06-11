package user

import (
	"projectanima/AnimaDB/domain"
)

// User -
type User struct {
	domain.AuditData

	Person domain.Person `json:"person,omitempty"`

	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Passkey  string `json:"passkey,omitempty"`
	Salt     string `json:"salt,omitempty"`
	Role     Role   `json:"role,omitempty"`
}

// Authenticate - Verify If The User Has Correctly Authenticated Themselves
func (userObject User) Authenticate() (bool, error) {
	return true, nil
}

// CheckAPIAccess - Check If The User Has Access To Get Information From A Certain API/Screen
// func (userObject User) CheckAPIAccess(context *persist.IDataContext, api string) (bool, error) {
// 	return true, nil
// }

// GetStamp - Get The User Stamp That Was Saved In The Database
func (userObject User) GetStamp() string {
	return userObject.Username
}
