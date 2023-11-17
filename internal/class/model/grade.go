package model

import "time"

type Grade struct {
	Id           string
	ClassId      string
	StudentId    string
	Grade        string
	LastModified time.Time
}
