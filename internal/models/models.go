package models

// swagger:model User
// User represents a user in the system.
type User struct {
	// ユーザーID
	// required: true
	ID int `json:"id"`
	// ユーザー名
	// required: true
	// example: John Doe
	Name string `json:"name"`
	// メールアドレス
	// required: true
	// example: john@example.com
	Email string `json:"email"`
}

// swagger:model Post
// Post represents a blog post in the system.
type Post struct {
	// 記事ID
	// required: true
	ID int `json:"id"`
	// 記事タイトル
	// required: true
	// example: My First Post
	Title string `json:"title"`
	// 記事本文
	// required: true
	Content string `json:"content"`
	// 投稿者情報
	// required: true
	Author User `json:"author"`
}

// Comment represents a comment on a blog post.
type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	User    User   `json:"user"`
	Content string `json:"content"`
}
