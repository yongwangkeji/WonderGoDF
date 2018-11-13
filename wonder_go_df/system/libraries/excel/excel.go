package excel

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

/**
 * fileName 生成的文件名 MyExcel
 * exportData 导出的数据对象
 * path 服务器的存储临时路径 /tmp
 */
func ExportExcel(fileName string, exportData [][]string, path string) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")

	for i := 0; i < len(exportData); i++ {
		row = sheet.AddRow() // 创建行
		for j := 0; j < len(exportData[i]); j++ {
			cell = row.AddCell()
			cell.Value = exportData[i][j]
		}
	}

	err = file.Save(path + fileName + ".xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
