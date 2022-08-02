package comment

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) PostComment(comment Comment) (Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Service) DeleteComment(ID uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *Service) GetAllComments() ([]Comment, error) {
	panic("not implemented") // TODO: Implement
}
