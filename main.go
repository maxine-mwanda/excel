package main

import "github.com/360EntSecGroup-Skylar/excelize"

func main() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		println(err.Error())
		return
	}
	// Get value from cell by given worksheet name and axis.
	var cell = f.GetCellValue("Sheet1", "B2")
	if err != nil {
		println(err.Error())
		return
	}
	println(cell)

	var rows = f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			print(colCell, "\t")
		}
		println()
	}
}
