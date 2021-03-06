package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	xlsx "github.com/tealeg/xlsx"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: git-xlsx-textconv file.xslx")
	}
	excelFileName := os.Args[1]

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			cels := make([]string, len(row.Cells))
			for _, cell := range row.Cells {
				s := cell.String()
				s = strings.Replace(s, "\\", "\\\\", -1)
				s = strings.Replace(s, "\n", "\\n", -1)
				s = strings.Replace(s, "\r", "\\r", -1)
				s = strings.Replace(s, "\t", "\\t", -1)
				cels = append(cels, s)
			}
			fmt.Printf("[%s] %s\n", sheet.Name, strings.Join(cels, "\t"))
		}
	}
}
