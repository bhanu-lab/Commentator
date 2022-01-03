package comment

import (
	"time"

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
	Created time.Time
}

type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetAllComments() ([]Comment, error)
	PostComment(comment Comment) (bool, error)
	UpdateComment(ID uint, comment Comment) (Comment, error)
	DeleteComment(ID uint)
}

//GetComment - retrieves comment from the ID given
func (s *Service) GetComment(ID int) (Comment, error) {
	var comment Comment
	if res := s.DB.First(&comment, ID); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

//GetAllComments - retrieves all comments present
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if res := s.DB.Find(&comments); res.Error != nil {
		return comments, res.Error
	}
	return comments, nil
}

//PostComment - adds new comment
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if res := s.DB.Save(&comment); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

//UpdateComment - updates comment for the ID given
func (s *Service) UpdateComment(ID int, newcomment Comment) (Comment, error) {
	var comment Comment
	if res := s.DB.Find(&comment, ID); res.Error != nil {
		return Comment{}, res.Error
	}

	if res := s.DB.Model(comment).Update(newcomment); res.Error != nil {
		return Comment{}, res.Error
	}
	return comment, nil
}

//DeleteComment - deletes comment based on ID given
func (s *Service) DeleteComment(ID int) error {
	if res := s.DB.Delete(&Comment{}, ID); res.Error != nil {
		return res.Error
	}
	return nil
}
