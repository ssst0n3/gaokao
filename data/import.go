package main

import (
	"fmt"
	"gaokao/db"
	"gaokao/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func Read(filename string) (err error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err = f.Close(); err != nil {
			awesome_error.CheckErr(err)
		}
	}()
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	var lastSchool model.School
	for _, row := range rows[1:] {
		//fmt.Println(row)
		var year, lowestScore, lowestRank, plan int
		yearRaw := row[0]
		year, err = strconv.Atoi(yearRaw)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		schoolCode := row[1]
		schoolName := row[2]
		subjectOrderRaw := row[3]
		var subjectOrder int
		switch subjectOrderRaw {
		case "本一":
			subjectOrder = 1
		case "本二":
			subjectOrder = 2
		}
		subjectTypeRaw := row[4]
		var subjectType int
		switch subjectTypeRaw {
		case "理科":
			subjectType = 1
		case "文科":
			subjectType = 2
		}
		subjectCode := row[5]
		subjectName := row[6]
		note := row[7]
		lowestScoreRaw := row[9]
		lowestScore, err = strconv.Atoi(lowestScoreRaw)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		lowestRankRaw := row[10]
		lowestRank, err = strconv.Atoi(lowestRankRaw)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		planRaw := row[14]
		plan, err = strconv.Atoi(planRaw)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		school := model.School{
			Name: schoolName,
			Code: schoolCode,
		}
		if school.Name != lastSchool.Name {
			err = db.CreateSchoolIfNotExists(&school)
			if err != nil {
				awesome_error.CheckErr(err)
				return
			}
			lastSchool = school
		}
		subject := model.Subject{
			Name:        subjectName,
			SchoolId:    school.ID,
			Code:        subjectCode,
			Order:       subjectOrder,
			Type:        subjectType,
			LowestScore: lowestScore,
			LowestRank:  lowestRank,
			Plan:        plan,
			Note:        note,
			Year:        year,
		}
		err = db.DB.Model(&subject).Create(&subject).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}

	}
	return
}
