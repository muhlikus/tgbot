package models

const (
	TextMessage = iota
	FileMessage
)

type messageType int

type Message struct {
	Id     int
	Type   messageType
	Date   int
	Text   string
	ChatID int
	File   *File
}

type Chat struct {
	Id   int
	Type string
}

type File struct {
	ID       string
	UniqueID string
	Size     int
	Name     string
	MimeType string
}
