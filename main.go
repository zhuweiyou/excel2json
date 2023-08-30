package main

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

func main() {
	defer func() {
		os.Stdin.Read(make([]byte, 1))
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: https://github.com/zhuweiyou/excel2json")
		return
	}

	from := os.Args[1]
	fmt.Printf("From: %s\n", from)

	file, err := excelize.OpenFile(from)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", from, err)
		return
	}
	defer file.Close()

	sheetList := file.GetSheetList()
	fmt.Printf("Sheet list: %v\n", sheetList)

	for index, sheet := range file.GetSheetList() {
		fmt.Printf("Sheet (%d/%d): %s\n", index+1, len(sheetList), sheet)
		rows, err := file.GetRows(sheet)
		if err != nil {
			fmt.Printf("Error getting rows: %v\n", err)
			return
		}
		fmt.Printf("Rows: %d\n", len(rows))

		marshal, err := json.Marshal(rows)
		if err != nil {
			fmt.Printf("Error marshaling rows: %v\n", err)
			return
		}

		dir, filename := filepath.Split(from)
		to := filepath.Join(dir, filename[:len(filename)-len(filepath.Ext(from))]+"_"+sheet+".json")
		err = os.WriteFile(to, marshal, 0666)
		if err != nil {
			fmt.Printf("Error writing to %v: %v\n", to, err)
			return
		}

		fmt.Printf("Successfully converted to %v\n", to)
	}
}
