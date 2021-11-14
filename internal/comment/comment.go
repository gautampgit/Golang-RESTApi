package comment

import (
	"github.com/jinzhu/gorm"
)

//Service- the struct for the comment service
type Service struct {
	DB *gorm.DB
}

//Comment - defines the structure of the comment
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

//CommentService - the interface for the comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

//NewService - the function returns a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

//GetComment - returns a comment by their ID in database
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.DB.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//GetCommentsBySlug - retuens all comments by slug(path -.va/comments)

func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

//PostComment - Posts a new comment in the database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//UpdateComment - Updates an existing comment and returns the new Comment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(&newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

//DeleteComment - Deletes a comment from the database
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

//GetAllComments = Retrieves all comments from the Database
func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}
