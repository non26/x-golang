package main

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	excelFile := excelize.NewFile()
	//// the sheet needs to be exising sheet in order to fill the cell in that sheet
	sheet1 := "Sheet1"
	excelFile.SetCellValue(sheet1, "A1", "column-1")

	//// save the file at the same level of main.go
	excelFile.SaveAs("./xExcel.xlsx")

	//// need to save on existing folder
	// excelFile.SaveAs("./excelFolder/xExcel.xls")
}
