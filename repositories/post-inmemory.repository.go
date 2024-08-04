package repositories

import (
	"blog/models"
	"time"
)

// Main struct

type PostInMemoryRepository struct {
	posts  map[int]models.Post
	nextID int
}

// Constructor

func NewPostInMemoryRepository() *PostInMemoryRepository {
	return &PostInMemoryRepository{
		posts:  make(map[int]models.Post),
		nextID: 0,
	}
}

// Methods

func (r *PostInMemoryRepository) GetById(id int) models.Post {
	return r.posts[id]
}

func (r *PostInMemoryRepository) Get() []models.Post {
	var postList []models.Post = make([]models.Post, 0)
	for _, p := range r.posts {
		postList = append(postList, p)
	}
	return postList
}

func (r *PostInMemoryRepository) Create(input models.CreatePostInput) models.Post {
	post := models.Post{
		ID:        r.nextID,
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: time.Now(),
	}
	r.posts[post.ID] = post
	r.nextID++
	return post
}

func (r *PostInMemoryRepository) Delete(postID int) {
	delete(r.posts, postID)
}
