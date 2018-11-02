package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

var (
	Infile  = "information.xlsx"
	OutFile = "out_information.xlsx"
)

type Info struct {
	Name   string
	Age    int
	Phone  string
	Gender string
	Mail   string
}

func Export() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("student_list")

	if err != nil {
		fmt.Printf(err.Error())
	}

	stus := getStudents()

	// add data
	for _, stu := range stus {

		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = stu.Name

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(stu.Age)

		phoneCell := row.AddCell()
		phoneCell.Value = stu.Phone

		genderCell := row.AddCell()
		genderCell.Value = stu.Gender

		mailCell := row.AddCell()
		mailCell.Value = stu.Mail

		err = file.Save(OutFile)
		if err != nil {
			fmt.Printf(err.Error())
		}

		fmt.Printf("Export Success!")
	}
}

func getStudents() []Info {
	students := make([]Info, 0)
	for i := 0; i < 10; i++ {
		stu := Info{}
		stu.Name = "name" + strconv.Itoa(i+1)
		stu.Mail = stu.Name + "@gmail.com"
		stu.Phone = "14505" + strconv.Itoa(i)
		stu.Age = 21
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}
