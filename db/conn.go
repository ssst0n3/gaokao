package db

import (
	"github.com/cloudquery/sqlite"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"gorm.io/gorm"
	"os"
	"sync"
)

var DB *gorm.DB
var lock = &sync.Mutex{}

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	awesome_error.CheckFatal(err)
	awesome_error.CheckFatal(Migrate())
}
