package repositories

import (
	"blog/db"
	"blog/models"
	"database/sql"
	"log"
)

// Main struct

type PostPostgresRepository struct {
	db *sql.DB
}

// Constructor

func NewPostPostgresRepository() *PostPostgresRepository {
	return &PostPostgresRepository{
		db: db.CreateDb(),
	}
}

// Methods

func (r *PostPostgresRepository) GetById(id int) models.Post {
	rows, err := r.db.Query("SELECT id, title, body, created_at FROM posts WHERE id = $1", id)
	if err != nil {
		log.Print("Failed to get post by id: ", err)
	}
	posts := bindToPost(rows)
	if len(posts) == 0 {
		return models.Post{}
	}
	return posts[0]
}

func (r *PostPostgresRepository) Get() []models.Post {
	rows, err := r.db.Query("SELECT id, title, body, created_at FROM posts")
	if err != nil {
		log.Print("Failed to get all posts: ", err)
	}
	return bindToPost(rows)
}

func (r *PostPostgresRepository) Create(input models.CreatePostInput) models.Post {
	row := r.db.QueryRow("INSERT INTO posts (title, body) VALUES ($1, $2) RETURNING id, title, body, created_at", input.Title, input.Body)
	var p = models.Post{}
	bindErr := row.Scan(&p.ID, &p.Title, &p.Body, &p.CreatedAt)
	if bindErr != nil {
		log.Print("Failed to bind row to post: ", bindErr)
	}
	return p
}

func (r *PostPostgresRepository) Delete(postID int) {
	r.db.Exec("DELETE FROM posts WHERE id = $1", postID)
}

// Binds all rows to a post models and returns a Posts array
func bindToPost(rows *sql.Rows) []models.Post {
	var posts []models.Post = make([]models.Post, 0)
	for {
		if !rows.Next() {
			break
		}
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt)
		if err != nil {
			log.Print("Failed to bind rows to post", rows, err)
			break
		}
		posts = append(posts, post)
	}
	return posts
}
