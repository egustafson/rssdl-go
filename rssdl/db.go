package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Feed struct {
	gorm.Model
	Title    string
	Href     string
	Subdir   string
	Active   bool
	Articles []Article
}

type Article struct {
	gorm.Model
	Title       string
	Link        string
	Guid        string
	PubAt       time.Time
	Feed        Feed
	FeedID      int
	Attachments []Attachment
}

type Attachment struct {
	gorm.Model
	MediaType string
	Link      string
	Path      string
	Expired   bool
	Article   Article
	ArticleID int
}

func openDB(dbPath string) (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", dbPath)
	return
}
