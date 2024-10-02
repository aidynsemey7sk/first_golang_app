package book

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Book struct {
	ID           int
	Title, Autor string
	Year         string
	Pages        int
}

func AddBook(books []Book, title, autor, year string, pages int) ([]Book, error) {
	if title == "" || autor == "" || year == "" || pages <= 0 {
		return books, errors.New("некорректные данные для книги")
	}
	newBook := Book{Title: title, Autor: autor, Year: year, Pages: pages, ID: len(books) + 1}

	log.Printf("Книга добавлена: %+v\n", newBook)

	return append(books, newBook), nil
}

// String возвращает строковое представление книги
func (b Book) String() string {
	return fmt.Sprintf("ID: %d, Название: %s, Автор: %s, Год: %s, Страницы: %d", b.ID, b.Title, b.Autor, b.Year, b.Pages)
}

// Поиск книги по ID
func FindBookByID(books []Book, id int) (Book, bool) {
	for _, b := range books {
		if b.ID == id {
			return b, true
		}
	}
	return Book{}, false
}

// RemoveBookByID удаляет книгу из списка по ID
func RemoveBookByID(books []Book, idToRemove int) []Book {
	var updatedBooks []Book
	for _, book := range books {
		if book.ID != idToRemove {
			updatedBooks = append(updatedBooks, book)
		}
	}
	return updatedBooks
}

func FindBookByTitle(books []Book, title string) (Book, bool) {
	title = strings.TrimSpace(strings.ToLower(title))
	for _, b := range books {
		if strings.EqualFold(strings.TrimSpace(b.Title), title) {
			return b, true
		}
	}
	return Book{}, false
}

// ListBooks выводит все книги
func ListBooks(books []Book) {
	if len(books) == 0 {
		fmt.Println("Нет ни одной книги")
		return
	}

	var bookWord string
	count := len(books)

	if count == 1 {
		bookWord = "книга"
	} else if count > 1 && count < 5 {
		bookWord = "книги"
	} else {
		bookWord = "книг"
	}

	fmt.Printf("Найдено %d %s\n", count, bookWord)

	for _, book := range books {
		fmt.Println(book.String())
	}
}
