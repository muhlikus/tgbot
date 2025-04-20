package models

type Message struct {
	MessageId int
	Date      int
	Text      string
	Chat      *Chat
	Document  *Document
}

type Chat struct {
	Id   int
	Type string
}

type Document struct {
	FileID       string
	FileUniqueID string
	FileSize     int
	FileName     string
	MimeType     string
}
