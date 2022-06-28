package db

import (
	"gaokao/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func CreateSchoolIfNotExists(school *model.School) (err error) {
	var count int64
	err = DB.Model(&model.School{}).Where(school).Count(&count).Error
	if count > 0 {
		return
	}
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = DB.Create(school).Model(school).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
