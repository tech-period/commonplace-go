package models

// User represents a user in the system.
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Post represents a blog post in the system.
type Post struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Author  User   `json:"author"`
}

// Comment represents a comment on a blog post.
type Comment struct {
    ID     int    `json:"id"`
    PostID int    `json:"post_id"`
    User   User   `json:"user"`
    Content string `json:"content"`
}