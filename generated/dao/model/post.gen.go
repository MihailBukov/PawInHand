// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePost = "post"

// Post mapped from table <post>
type Post struct {
	ID          string      `gorm:"column:id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string      `gorm:"column:title;not null" json:"title"`
	Content     string      `gorm:"column:content;not null" json:"content"`
	AuthorID    string      `gorm:"column:author_id;not null" json:"author_id"`
	DateCreated time.Time   `gorm:"column:date_created;default:CURRENT_DATE" json:"date_created"`
	Images      interface{} `gorm:"column:images" json:"images"`
	CreatedAt   time.Time   `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Post's table name
func (*Post) TableName() string {
	return TableNamePost
}
