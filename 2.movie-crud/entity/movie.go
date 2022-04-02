package entity

// struct movie (id, title, actor)
type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Actor *Actor `json:"actor"`
}