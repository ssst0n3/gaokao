package db

import (
	"fmt"
	"gaokao/model"
)

func QueryNewSubject() (err error) {
	schools2017, err := querySchoolByYear(2017)
	if err != nil {
		return
	}
	for _, school := range schools2017 {
		var names []string
		var newSubjects []model.Subject
		DB.Model(&model.Subject{}).Select("name").Where("year=? AND school_id=?", 2016, school.ID).Find(&names)
		DB.Model(&model.Subject{}).Select("name").Where(
			"year=? AND school_id=?", 2017, school.ID,
		).Not(map[string]interface{}{"name": names}).Find(&newSubjects)
		fmt.Println(newSubjects)
	}
	return
}
