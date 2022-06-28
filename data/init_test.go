package main

import (
	"gaokao/db"
	"gaokao/test/test_config"
)

func init() {
	test_config.Init()
	db.Init()
}
