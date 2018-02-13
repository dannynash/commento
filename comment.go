package main

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Comment   string    `json:"comment"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Parent    int       `json:"parent"`
}

type CommentService interface {
	CreateComment(comment *Comment) error
	GetComments(url string) ([]Comment, error)
	GetCommentsWithParent(url string, parentID string) ([]Comment, error)
	UpdateComment(comment *Comment) (*Comment, error)
}
