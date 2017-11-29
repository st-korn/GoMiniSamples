package main

import "github.com/tealeg/xlsx"

func main() {
    // Создаём новую таблицу Excel
    file := xlsx.NewFile()

    // Добавляем лист
    sheet, err := file.AddSheet("Лист1")
    if err != nil { panic(err) }

    // Помещаем в первую ячейку листа строку
    cell := sheet.Cell(0,0)
    cell.SetString("Hello, world!")

    // Сохраняем таблицу в файл
    err = file.Save("sample.xlsx")
    if err != nil { panic(err) }
}