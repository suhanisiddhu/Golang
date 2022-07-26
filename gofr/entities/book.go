package entities

type Book struct {
	BookID        int    `json:"bookID,omitempty"`
	AuthorID      int    `json:"authorID,omitempty"`
	Title         string `json:"title,omitempty"`
	Publication   string `json:"publication,omitempty"`
	PublishedDate string `json:"publishedDate,omitempty"`
	Author        Author `json:"author,omitempty"`
}
