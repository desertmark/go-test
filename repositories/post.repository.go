package repositories

import "time"

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

var posts map[int]Post = make(map[int]Post)
var nextID int = 0

func GetById(id int) Post {
	return posts[id]
}

func Get() []Post {
	var postList []Post = make([]Post, 0)
	for _, p := range posts {
		postList = append(postList, p)
	}
	return postList
}

func Create(input CreatePostInput) Post {
	post := Post{
		ID:        nextID,
		Title:     input.Title,
		Body:      input.Body,
		CreatedAt: time.Now(),
	}
	posts[post.ID] = post
	nextID++
	return post
}

func Delete(postID int) {
	delete(posts, postID)
}
