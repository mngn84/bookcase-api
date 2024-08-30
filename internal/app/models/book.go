package bookmodel

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Genre      string `json:"genre"`
	IsRead     bool   `json:"isRead"`
	LocationId *int    `json:"locationId,omitempty"`
}

type Location struct {
	ID         int    `json:"id"`
	CaseId    int    `json:"caseId"`
	ShelfName string `json:"shelfName"`
}
