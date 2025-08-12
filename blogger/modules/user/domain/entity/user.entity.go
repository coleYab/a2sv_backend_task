// Package entity: user entity
package entity

import "time"

const (
	UserStatusVerified = "VERIFIED"
	UserStatusUnVerified = "UNVERIFIED"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role 		string 		`json:"role"`
	Status 		string `json:"status"`
	Password  string    `json:"-"`
	Profile  Profile   `json:"profile"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Profile struct {
	Bio            string            `json:"bio,omitempty"`             // optional
	ProfilePicture string            `json:"profilePicture,omitempty"`  // URL to profile picture
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}


func NewUser(id, userName, email, role, password string, createdAt, updatedAT time.Time) User {
	return User{
		ID: id,
		Username: userName,
		Email: email,
		Role: role,
		Status: UserStatusUnVerified,
		Profile: Profile{
			Bio: "",
			ProfilePicture: "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Password: password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAT,
	}
}

func (u *User) UpdateProfile(bio, profilePicture string) {
	u.Profile.Bio = bio
	u.Profile.ProfilePicture = profilePicture
	u.Profile.UpdatedAt = time.Now()
}

func (u *User) Update(userName, password string) {
	u.Username = userName
	u.Password = password
	u.UpdatedAt = time.Now()
}

func (u *User) Promote(role string) {
	u.Role = role
}

func (u *User) Verfiy() {
	u.Status = UserStatusVerified
}