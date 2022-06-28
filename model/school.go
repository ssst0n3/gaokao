package model

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type School struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}

var SchemaSchool schema.Schema

func init() {
	awesome_error.CheckFatal(InitSchema(&SchemaSchool, &School{}))
}
