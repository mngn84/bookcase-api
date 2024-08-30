package bookmodel

type MoveBookRequest struct {
	ID       int      `json:"id"`
	Location Location `json:"location"`
}

type FullBookRequest struct {
	Title    string   `json:"title"`
	Author   string   `json:"author"`
	Genre    string   `json:"genre"`
	IsRead   bool     `json:"is_read"`
	Location Location `json:"location"`
}
