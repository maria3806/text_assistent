package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text_assistent/processor"
)

func main() {
	folder := "./texts"
	files := processor.ListFiles(folder)
	if len(files) == 0 {
		fmt.Println("У папці немає файлів")
		return
	}

	fmt.Println("Файли у папці:")
	for i, f := range files {
		fmt.Printf("[%d] %s\n", i+1, f)
	}

	fmt.Print("Виберіть файл за номером: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	idx, _ := strconv.Atoi(input)
	if idx < 1 || idx > len(files) {
		fmt.Println("Невірний номер файлу")
		return
	}
	filePath := filepath.Join(folder, files[idx-1])
	data, _ := os.ReadFile(filePath)
	text := string(data)

	fmt.Println("Виберіть завдання:")
	fmt.Println("[1] Показати розмір і переіменувати")
	fmt.Println("[2] Головна мета тексту")
	fmt.Println("[3] Короткий виклад тексту")
	fmt.Println("[4] Переписати текст у стилі")
	fmt.Println("[5] Показати ключові слова")
	fmt.Println("[6] Конвертувати у формат MD/HTML")
	fmt.Println("[7] Стиснути текст до певної кількості слів")
	fmt.Println("[8] Згенерувати запитання до тексту")
	fmt.Print("Номер завдання: ")
	taskInput, _ := reader.ReadString('\n')
	taskInput = strings.TrimSpace(taskInput)

	switch taskInput {
	case "1":
		fmt.Println("Розмір файлу:", processor.FileSize(filePath), "байт")
		fmt.Print("Введіть нове ім'я файлу: ")
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)
		err := processor.RenameFile(filePath, newName)
		if err != nil {
			fmt.Println("Помилка при переіменуванні")
		} else {
			fmt.Println("Файл переіменовано:", newName)
		}
	case "2":
		fmt.Println("Головна мета тексту:", processor.GetMainGoal(text))
	case "3":
		fmt.Println("Короткий виклад тексту:", processor.ShortSummary(text))
	case "4":
		fmt.Print("Введіть стиль (казка, пост для соцмереж, жартівливий): ")
		style, _ := reader.ReadString('\n')
		style = strings.TrimSpace(style)
		fmt.Println("Переписаний текст:", processor.RewriteInStyle(text, style))
	case "5":
		fmt.Println("Ключові слова:", processor.GetKeywords(text, 10))
	case "6":
		fmt.Println("Введіть формат (md/html): ")
		format, _ := reader.ReadString('\n')
		format = strings.TrimSpace(format)
		if format == "md" {
			processor.SaveAsMD(text, filepath.Join(folder, files[idx-1]))
		} else {
			processor.SaveAsHTML(text, filepath.Join(folder, files[idx-1]))
		}
	case "7":
		fmt.Print("Введіть кількість слів: ")
		numStr, _ := reader.ReadString('\n')
		numStr = strings.TrimSpace(numStr)
		num, _ := strconv.Atoi(numStr)
		fmt.Println("Стиснутий текст:", processor.CompressText(text, num))
	case "8":
		fmt.Print("Скільки питань згенерувати: ")
		countStr, _ := reader.ReadString('\n')
		countStr = strings.TrimSpace(countStr)
		count, _ := strconv.Atoi(countStr)
		fmt.Println("Запитання:\n", processor.GenerateQuestions(text, count))

	default:
		fmt.Println("Невірний номер завдання")
	}
}
