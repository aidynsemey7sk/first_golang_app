package main

import (
	"bufio"
	"first_project/pkg/book"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var action int
	var mainBooks []book.Book

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatalf("Ошибка создания директории logs: %v", err)
	}

	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Ошибка открытия лог файла: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Println("--- Меню ---")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Удалить книгу")
		fmt.Println("3. Найти книгу по названию")
		fmt.Println("4. Показать все книги")
		fmt.Println("5. Выйти")
		fmt.Printf("\n\nВведите номер действия: ")
		fmt.Scan(&action)

		switch action {
		case 1:
			var year string
			var pages int

			fmt.Print("Введите название книги: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Введите автора книги: ")
			scanner.Scan()
			autor := scanner.Text()

			for {
				fmt.Print("Введите год издания: ")
				scanner.Scan()
				year = scanner.Text()
				if _, err := strconv.Atoi(year); err == nil {
					break
				}
				fmt.Println("Неверный ввод. Пожалуйста, введите год в формате числа.")
			}

			for {
				fmt.Print("Введите количество страниц: ")
				scanner.Scan()
				pagesInput := scanner.Text()

				var err error
				pages, err = strconv.Atoi(pagesInput) // Удаление лишнего объявления переменной
				if err == nil && pages > 0 {
					break
				}
				fmt.Println("Неверный ввод. Пожалуйста, введите положительное число для количества страниц.")

			}

			// Вызов функции AddBook

			mainBooks, err = book.AddBook(mainBooks, title, autor, year, pages)
			if err != nil {
				log.Printf("Ошибка при добавлении книги: %v\n", err)
			}

		case 2:
			var id int
			fmt.Printf("Введите айди книги: ")
			fmt.Scan(&id)

			bookToDelete, found := book.FindBookByID(mainBooks, id)
			if !found {
				fmt.Println("Книга с таким ID не найдена.")
				break
			}
			mainBooks = book.RemoveBookByID(mainBooks, id)
			log.Printf("Книга удалена: ID: %d, Название: %s, Автор: %s, Год: %s, Страницы: %d\n", bookToDelete.ID, bookToDelete.Title, bookToDelete.Autor, bookToDelete.Year, bookToDelete.Pages)
			fmt.Printf("Книга удалена!\n")
		case 3:

			fmt.Println("Введите название книги:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			bookTitle := scanner.Text()
			fmt.Println("Вывод", bookTitle)

			res, found := book.FindBookByTitle(mainBooks, bookTitle)

			if found {
				fmt.Printf("ID: %d, Название: %s, Автор: %s, Год: %s, Страницы: %d\n\n", res.ID, res.Title, res.Autor, res.Year, res.Pages)

			} else {
				fmt.Println("Книга с таким названием не найдена.")
			}

		case 4:
			book.ListBooks(mainBooks)
		case 5:
			fmt.Println("Выход из программы.")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}

}
