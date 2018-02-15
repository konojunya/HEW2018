package model

type Product struct {
	ID        int    `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Votes     int    `json:"votes"`
}
