// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        string    `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	FirstName string    `gorm:"column:first_name;not null" json:"first_name"`
	LastName  string    `gorm:"column:last_name;not null" json:"last_name"`
	Username  string    `gorm:"column:username;not null" json:"username"`
	Email     string    `gorm:"column:email;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
