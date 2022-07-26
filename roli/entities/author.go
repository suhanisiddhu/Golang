package entities

type Author struct {
	AuthorID  int    `json:"authorID,omitempty"`
	FirstName string `json:"firstName,omiempty"`
	LastName  string `json:"lastName"`
	DOB       string `json:"DOB"`
	PenName   string `json:"penName"`
}
