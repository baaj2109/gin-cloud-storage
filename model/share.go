package model

type Share struct {
	Id       int
	Code     string
	FileId   int
	Username string
	Hash     string
}

func (Share) TableName() string {
	return "share"
}
