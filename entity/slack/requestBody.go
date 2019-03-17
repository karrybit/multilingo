package slack

// RequestBody -
type RequestBody struct {
	Token       string        `json:"token"`
	Channel     string        `json:"channel"`
	Attachments []*Attachment `json:"attachments"`
	UserName    string        `json:"username"`
}

// Attachment -
// https://api.slack.com/docs/message-attachments
type Attachment struct {
	Color     string `json:"color"` // good or warning or danger or colorcode
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Text      string `json:"text"`
}
