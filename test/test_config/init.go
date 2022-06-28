package test_config

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"os"
)

var AlreadyInit bool

func Init() {
	if AlreadyInit {
		log.Logger.Debug("already init")
	} else {
		path := "../data/gaokao.sqlite"
		//err := os.Remove(path)
		//if err != nil {
		//	if !os.IsNotExist(err) {
		//		awesome_error.CheckFatal(err)
		//	}
		//}
		awesome_error.CheckFatal(os.Setenv("DB_DSN", path))
		AlreadyInit = true
	}
}
