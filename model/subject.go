package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Subject struct {
	gorm.Model
	Name        string `json:"name"`
	SchoolId    uint   `json:"school_id"`
	Code        string `json:"code"`
	Order       int    `json:"order"`
	Type        int    `json:"type"`
	LowestScore int    `json:"lowest_score"`
	LowestRank  int    `json:"lowest_rank"`
	Plan        int    `json:"plan"`
	Note        string `json:"note"`
	Year        int    `json:"year"`
}

var SchemaSubject schema.Schema

func init() {
	awesome_error.CheckFatal(InitSchema(&SchemaSubject, &Subject{}))
}
