package main

import (
	"gaokao/db"
	"gaokao/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func ExtendSchoolLowestScore() {
	for _, year := range []int{2021, 2020, 2019, 2018, 2017, 2016} {
		var schools []model.School
		awesome_error.CheckFatal(db.DB.Model(&model.Subject{}).Select("subjects.school_id AS ID, max(subjects.lowest_rank) AS lowest_rank").Joins(
			"JOIN schools ON subjects.school_id=schools.id AND schools.year=subjects.year",
		).Where(
			"subjects.year=?", year,
		).Group("subjects.school_id").Find(&schools).Error)
		for _, school := range schools {
			awesome_error.CheckFatal(db.DB.Model(model.School{}).Where(school.ID).Update("lowest_rank", school.LowestRank).Error)
		}
	}
}
