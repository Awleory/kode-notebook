package entity

type NoteCreating struct {
	OwnerId int    `json:"ownerId"`
	Title   string `json:"title"`
	Text    string `json:"text"`
}

type Note struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
