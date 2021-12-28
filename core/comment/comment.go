package comment

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	DB *gorm.DB
}

/*
NewService - new service creates Service struct
and returns pointer to it
*/
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.time
}

type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetAllComments() ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, comment Comment) (Comment, error)
	DeleteComment(ID uint)
}
