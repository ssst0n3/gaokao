package db

import (
	"fmt"
	"gaokao/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

type Result struct {
	SubjectName string `json:"subject_name"`
	SchoolName  string `json:"school_name"`
	LowestScore int
}

func QueryWaveSchool() {

}

func querySchoolByYear(year int) (schools []model.School, err error) {
	err = DB.Model(&schools).Find(&schools, "year=?", year).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func querySubjectByYear(year int) (subjects []model.Subject, err error) {
	err = DB.Model(&subjects).Find(&subjects, "year=?", year).Error
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func Query() (err error) {
	var allSubjects [][]Result
	years := []int{
		2021, 2020, 2019, 2018, 2017, 2016,
	}
	for _, year := range years {
		var results []Result
		err = DB.Model(&model.Subject{}).Select("subjects.name AS subject_name, schools.name AS school_name").Joins("JOIN schools ON subjects.school_id=schools.id AND schools.year=subjects.year").Find(&results, "lowest_rank > 60000 and lowest_rank < 90000 and subjects.year = ?", year).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		allSubjects = append(allSubjects, results)
	}
	counts := map[string]int{}
	for _, subjects := range allSubjects {
		for _, subject := range subjects {
			uuid := fmt.Sprintf("%s-%s", subject.SchoolName, subject.SubjectName)
			if _, ok := counts[uuid]; !ok {
				counts[uuid] = 1
			} else {
				counts[uuid] += 1
			}
		}
	}
	for id, count := range counts {
		if count < 6 {
			fmt.Println(id)
		}
	}
	return
}
