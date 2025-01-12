package services

import (
	api "PawInHand/generated/api/PawInHand/rest"
	"PawInHand/generated/dao/model"
	"PawInHand/repositories"
	"PawInHand/resources"
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type PostAPIService struct {
	DB      repositories.PostRepo
	Convert resources.ConverterI
}

func NewPostAPIService(db repositories.PostRepo, convert resources.ConverterI) api.PostsAPIServicer {
	return &PostAPIService{
		DB:      db,
		Convert: convert,
	}
}

// GetAllPosts -
func (s *PostAPIService) GetAllPosts(ctx context.Context) (api.ImplResponse, error) {
	posts, err := s.DB.GetAll()
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	var postApis []*api.Post
	for _, post := range posts {
		postApi := s.Convert.ConvertModelPostToApiPost(post)
		postApis = append(postApis, postApi)
	}

	return api.Response(http.StatusOK, postApis), nil
}

// GetPostById -
func (s *PostAPIService) GetPostById(ctx context.Context, postId string) (api.ImplResponse, error) {
	if postId == "" {
		return api.Response(http.StatusBadRequest, "postId is required"), nil
	}

	post, err := s.DB.FindByID(postId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Post not found"), err
	}

	postApi := s.Convert.ConvertModelPostToApiPost(post)
	return api.Response(http.StatusOK, postApi), nil
}

// CreatePost -
func (s *PostAPIService) CreatePost(ctx context.Context, post api.Post) (api.ImplResponse, error) {
	postModel := s.Convert.ConvertApiPostToModelPost(&post)
	newPost, err := s.DB.Create(postModel)
	if err != nil {
		return api.Response(http.StatusInternalServerError, nil), err
	}

	postApi := s.Convert.ConvertModelPostToApiPost(newPost)
	return api.Response(http.StatusCreated, postApi), nil
}

// DeletePost -
func (s *PostAPIService) DeletePostById(ctx context.Context, postId string) (api.ImplResponse, error) {
	if postId == "" {
		return api.Response(http.StatusBadRequest, "postId is required"), nil
	}

	_, err := s.DB.DeleteByID(postId)
	if err != nil {
		return api.Response(http.StatusNotFound, "Post not found"), err
	}

	return api.Response(http.StatusNoContent, nil), nil
}
func (s *PostAPIService) UpdatePostById(ctx context.Context, postID string, post api.Post) (api.ImplResponse, error) {
	// Validate input: Ensure required fields are present
	if post.Title == "" || post.Content == "" {
		return api.Response(http.StatusBadRequest, "Title and content are required"), nil
	}

	// Parse and validate optional date fields if applicable
	var postDate time.Time
	var err error
	if post.DateCreated != "" {
		postDate, err = time.Parse("2006-01-02", post.DateCreated)
		if err != nil {
			return api.Response(http.StatusBadRequest, "Invalid date format"), nil
		}
	}

	// Map API model to DAO/database model
	updatedPost := &model.Post{
		ID:          postID,
		Title:       post.Title,
		Content:     post.Content,
		AuthorID:    post.AuthorId,
		DateCreated: postDate,
		Image:       post.Image,
	}

	// Update the post in the database/repository
	_, err = s.DB.UpdateByID(postID, updatedPost)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return api.Response(http.StatusNotFound, "Post not found"), nil
		}
		return api.Response(http.StatusInternalServerError, "Failed to update post"), err
	}

	// Return success response
	return api.Response(http.StatusOK, "Post updated successfully"), nil
}
