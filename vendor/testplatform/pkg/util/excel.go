package util

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	
	"testplatform/models"
)

func WriteToExcel(data []models.Testcase, path string) (err error) {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "编号")
	f.SetCellValue("Sheet1", "B1", "模块")
	f.SetCellValue("Sheet1", "C1", "重要性")
	f.SetCellValue("Sheet1", "D1", "描述")
	f.SetCellValue("Sheet1", "E1", "步骤")
	f.SetCellValue("Sheet1", "F1", "结果")

	if (len(data) > 0){
		for index, value := range data {
			f.SetCellValue("Sheet1", "A" + strconv.Itoa(index + 2), value.Id)
			f.SetCellValue("Sheet1", "B" + strconv.Itoa(index + 2), value.Module)
			f.SetCellValue("Sheet1", "C" + strconv.Itoa(index + 2), value.Importance)
			f.SetCellValue("Sheet1", "D" + strconv.Itoa(index + 2), value.Describe)
			f.SetCellValue("Sheet1", "E" + strconv.Itoa(index + 2), value.Step)
			f.SetCellValue("Sheet1", "F" + strconv.Itoa(index + 2), value.Result)
		}
	}
	
	err = f.SaveAs(path)

	return err
}