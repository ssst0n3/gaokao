package main

import (
	"fmt"
	"gaokao/db"
	"gaokao/model"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/xuri/excelize/v2"
	"strconv"
)

type Index struct {
	Year        int
	SchoolName  int
	SchoolCode  int
	SubjectType int
	SubjectName int
	SubjectCode int
	Order       int
	LowestScore int
	LowestRank  int
	Plan        int
}

type Parser interface {
	filename() string
	sheet() string
	columnIndex() (index Index)
	parseOrder(raw string) (order int)
}

func Parse() (err error) {
	for _, parser := range []Parser{
		Parser2021{}, Parser2020{}, Parser2019{}, Parser2018{}, Parser2017{}, Parser2016{},
	} {
		err = parseYear(parser)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

func parseYear(parser Parser) (err error) {
	f, err := excelize.OpenFile(parser.filename())
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

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows(parser.sheet())
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	index := parser.columnIndex()
	var lastSchool model.School
	for _, row := range rows[1:] {
		var year, lowestScore, lowestRank, plan int
		yearRaw := row[index.Year]
		year, err = strconv.Atoi(yearRaw)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
		schoolCode := row[index.SchoolCode]
		schoolName := row[index.SchoolName]
		orderRaw := row[index.Order]
		subjectOrder := parser.parseOrder(orderRaw)
		subjectTypeRaw := row[index.SubjectType]
		var subjectType int
		switch subjectTypeRaw {
		case "理科":
			subjectType = 1
		case "文科":
			subjectType = 2
		}
		subjectCode := row[index.SubjectCode]
		subjectName := row[index.SubjectName]
		lowestScoreRaw := row[index.LowestScore]
		lowestScore, err = strconv.Atoi(lowestScoreRaw)
		if err != nil {
			if lowestScoreRaw == "-" {
				lowestScore = -1
			} else {
				f, err2 := strconv.ParseFloat(lowestScoreRaw, 32)
				if err2 != nil {
					fmt.Println(row)
					err = err2
					awesome_error.CheckErr(err)
					return
				}
				lowestScore = int(f)
			}
		}
		lowestRankRaw := row[index.LowestRank]
		lowestRank, err = strconv.Atoi(lowestRankRaw)
		if err != nil {
			if lowestRankRaw == "-" {
				lowestRank = -1
			} else {
				fmt.Println(row)
				awesome_error.CheckErr(err)
				return
			}
		}
		planRaw := row[index.Plan]
		plan, err = strconv.Atoi(planRaw)
		if err != nil {
			if planRaw == "-" {
				plan = -1
			} else {
				awesome_error.CheckErr(err)
				return
			}
		}
		school := model.School{
			Year: year,
			Name: schoolName,
			Code: schoolCode,
		}
		subject := model.Subject{
			Name:        subjectName,
			Code:        subjectCode,
			Order:       subjectOrder,
			Type:        subjectType,
			LowestScore: lowestScore,
			LowestRank:  lowestRank,
			Plan:        plan,
			Year:        year,
		}
		if school.Code != lastSchool.Code {
			err = db.CreateSchoolIfNotExists(&school)
			if err != nil {
				awesome_error.CheckErr(err)
				return
			}
			lastSchool = school
		}
		if subjectOrder != 1 {
			continue
		}
		if subjectType != 1 {
			continue
		}
		subject.SchoolId = lastSchool.ID
		err = db.DB.Model(&subject).Create(&subject).Error
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

type Parser2016 struct{}

func (p Parser2016) filename() string { return "A安徽2016.xlsx" }
func (p Parser2016) sheet() string    { return "专业分数" }
func (p Parser2016) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolName:  1,
		SchoolCode:  2,
		SubjectType: 5,
		SubjectName: 6,
		SubjectCode: 7,
		Order:       8,
		LowestScore: 12,
		LowestRank:  14,
		Plan:        18,
	}
}
func (p Parser2016) parseOrder(raw string) (order int) {
	switch raw {
	case "本科提前批":
		order = 0
	case "本科第一批":
		order = 1
	case "本科第二批":
		order = 2
	case "高职专科批":
		order = 3
	case "国家专项计划本科批":
		order = 4
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}

type Parser2017 struct{}

func (p Parser2017) filename() string { return "A安徽2017.xlsx" }
func (p Parser2017) sheet() string    { return "专业分数" }
func (p Parser2017) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolName:  1,
		SchoolCode:  2,
		SubjectType: 5,
		SubjectName: 6,
		SubjectCode: 7,
		Order:       8,
		LowestScore: 12,
		LowestRank:  14,
		Plan:        18,
	}
}
func (p Parser2017) parseOrder(raw string) (order int) {
	switch raw {
	case "本科提前批":
		order = 0
	case "本科第一批":
		order = 1
	case "本科第二批":
		order = 2
	case "高职专科批":
		order = 3
	case "国家专项计划本科批":
		order = 4
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}

type Parser2018 struct{}

func (p Parser2018) filename() string { return "A安徽2018.xlsx" }
func (p Parser2018) sheet() string    { return "专业分数" }
func (p Parser2018) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolName:  1,
		SchoolCode:  2,
		SubjectType: 5,
		SubjectName: 6,
		SubjectCode: 7,
		Order:       8,
		LowestScore: 12,
		LowestRank:  14,
		Plan:        18,
	}
}
func (p Parser2018) parseOrder(raw string) (order int) {
	switch raw {
	case "本科提前批":
		order = 0
	case "本科第一批":
		order = 1
	case "本科第二批":
		order = 2
	case "高职专科批":
		order = 3
	case "国家专项计划本科批":
		order = 4
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}

type Parser2019 struct{}

func (p Parser2019) filename() string { return "A安徽2019.xlsx" }
func (p Parser2019) sheet() string    { return "专业分数" }
func (p Parser2019) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolName:  1,
		SchoolCode:  2,
		SubjectType: 5,
		SubjectName: 6,
		SubjectCode: 7,
		Order:       8,
		LowestScore: 12,
		LowestRank:  14,
		Plan:        18,
	}
}
func (p Parser2019) parseOrder(raw string) (order int) {
	switch raw {
	case "本科提前批":
		order = 0
	case "本科第一批":
		order = 1
	case "本科第二批":
		order = 2
	case "高职专科批":
		order = 3
	case "国家专项计划本科批":
		order = 4
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}

type Parser2020 struct{}

func (p Parser2020) filename() string { return "A安徽-2020.xlsx" }
func (p Parser2020) sheet() string    { return "专业分数" }
func (p Parser2020) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolName:  1,
		SchoolCode:  2,
		SubjectType: 5,
		SubjectName: 6,
		SubjectCode: 7,
		Order:       8,
		LowestScore: 12,
		LowestRank:  14,
		Plan:        18,
	}
}
func (p Parser2020) parseOrder(raw string) (order int) {
	switch raw {
	case "本科提前批":
		order = 0
	case "本科第一批":
		order = 1
	case "本科第二批":
		order = 2
	case "高职专科批":
		order = 3
	case "国家专项计划本科批":
		order = 4
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}

type Parser2021 struct{}

func (p Parser2021) filename() string { return "A安徽2021.xlsx" }
func (p Parser2021) sheet() string    { return "Sheet1" }
func (p Parser2021) columnIndex() Index {
	return Index{
		Year:        0,
		SchoolCode:  1,
		SchoolName:  2,
		Order:       3,
		SubjectType: 4,
		SubjectCode: 5,
		SubjectName: 6,
		LowestScore: 9,
		LowestRank:  10,
		Plan:        14,
	}
}
func (p Parser2021) parseOrder(raw string) (order int) {
	switch raw {
	case "本一":
		order = 1
	case "本二":
		order = 2
	case "专科":
		order = 3
	default:
		panic(fmt.Sprintf("not valid: %s", raw))
	}
	return
}
