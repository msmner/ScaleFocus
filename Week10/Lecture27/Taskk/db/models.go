// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"time"
)

type Story struct {
	StoryID    int32
	StoryTitle string
	StoryScore int32
	CreatedAt  time.Time
}