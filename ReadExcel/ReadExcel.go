package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var InFile = "information.xlsx"

func ImportFile() {
	xlFile, err := xlsx.OpenFile(InFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, sheet := range xlFile.Sheets {
		fmt.Println("Sheet Name: ", sheet.Name)

		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%20s", text)
			}
			fmt.Print("\n")
		}
	}

	fmt.Println("import success")
}
