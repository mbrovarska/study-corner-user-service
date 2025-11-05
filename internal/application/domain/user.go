package domain

import "time"

type User struct{
	ID				int		   `json:"id"`
	Username 		string	   `json:"username"`
	Email			string 	   `json:"email"`
	PasswordHash 	string	   `json:"password"`
	Avatar			string     `json:"avatar"`
	IsBanned		bool	   `json:"isBanned"`
	DeactivatedAt 	time.Time  `json:"deactivatedAt"`
	CreatedAt 		time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`

}
