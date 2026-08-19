package model

// User represents an application user.
type User struct {
	ID    int64
	Name  string
	Email string
}

// UpdateName is supposed to change the user's name.
//
// TODO(candidate): as written, callers never see the change take effect.
// Fix the receiver type so this method mutates the caller's User.
func (u User) UpdateName(name string) {
	u.Name = name
}
