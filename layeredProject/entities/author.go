package entities

type Author struct {
	AuthorID  int    `json:"authorID,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	DOB       string `json:"dOB,omitempty"`
	PenName   string `json:"penName,omitempty"`
}
