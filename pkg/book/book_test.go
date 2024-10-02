package book_test

import (
	"bytes"
	"first_project/pkg/book"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func TestRemoveBookByID(t *testing.T) {
	books := []book.Book{
		{ID: 1, Title: "Book 1", Autor: "Author 1", Year: "2000", Pages: 100},
		{ID: 2, Title: "Book 2", Autor: "Author 2", Year: "2001", Pages: 200},
	}

	// Удаляем книгу с ID 1
	updatedBooks := book.RemoveBookByID(books, 1)

	// Ожидаем, что осталась только одна книга с ID 2
	if len(updatedBooks) != 1 {
		t.Errorf("Ожидается 1 книга, получено %d", len(updatedBooks))
	}

	if updatedBooks[0].ID != 2 {
		t.Errorf("Ожидаемая книга с идентификатором 2, получен %d", updatedBooks[0].ID)
	}

	// Ожидаем что невозможно удалить несуществующюю книгу

	unchangedBooks := book.RemoveBookByID(updatedBooks, 10)
	// Проверяем, что длина списка осталась прежней
	if len(unchangedBooks) != len(updatedBooks) {
		t.Errorf("Длина списка книг изменилась, ожидалась %d, получено %d", len(updatedBooks), len(unchangedBooks))
	}

}

func TestFindBookByTitle(t *testing.T) {
	books := []book.Book{
		{ID: 1, Title: "Война и Мир", Autor: "Толстой", Year: "1869", Pages: 1225},
		{ID: 2, Title: "Преступление и наказание", Autor: "Достоевский", Year: "1866", Pages: 671},
	}

	// Ищем книгу по названию
	foundBook, found := book.FindBookByTitle(books, "Война и Мир")

	if !found {
		t.Errorf("Ожидал найти книгу, но она не была найдена")
	}

	if foundBook.Title != "Война и Мир" {
		t.Errorf("Ожидал 'Война и Мир', получил'%s'", foundBook.Title)
	}

	// Проверка на книгу, которой нет
	_, found = book.FindBookByTitle(books, "Unknown Title")
	if found {
		t.Errorf("Ожидал не найти книгу, но она нашлась")
	}

	_, found = book.FindBookByTitle(books, "Гордость и предубеждение")
	if found {
		t.Errorf("Ожидал не найти книгу 'Гордость и предубеждение', но она нашлась")
	}

}

func TestFindBookByID(t *testing.T) {
	books := []book.Book{
		{ID: 1, Title: "Book 1", Autor: "Author 1", Year: "2000", Pages: 100},
		{ID: 2, Title: "Book 2", Autor: "Author 2", Year: "2001", Pages: 200},
	}

	// Ищем книгу по ID
	foundBook, found := book.FindBookByID(books, 1)

	if !found {
		t.Errorf("Ожидали найти книгу с ID 1, но она не была найдена.")
	}

	if foundBook.ID != 1 {
		t.Errorf("Ожидался ID 1, получен %d", foundBook.ID)
	}

	// Проверяем на книгу с несуществующим ID
	_, found = book.FindBookByID(books, 3)
	if found {
		t.Errorf("Ожидалось, что книга с ID 3 не будет найдена, но она была найдена.")
	}
}

func TestAddBook(t *testing.T) {
	// Исходный список книг
	books := []book.Book{}

	// Сценарий 1: Успешное добавление книги
	newBooks, err := book.AddBook(books, "Война и Мир", "Толстой", "1869", 1225)
	if err != nil {
		t.Errorf("Не ожидалось ошибки, но получили: %v", err)
	}

	if len(newBooks) != 1 {
		t.Errorf("Ожидалось, что добавится 1 книга, но в списке %d", len(newBooks))
	}

	if newBooks[0].Title != "Война и Мир" {
		t.Errorf("Ожидалось название 'Война и Мир', получено: '%s'", newBooks[0].Title)
	}

	// Сценарий 2: Добавление книги с некорректными данными
	newBooks, err = book.AddBook(newBooks, "", "Автор", "2000", 100)
	if err == nil {
		t.Error("Ожидалась ошибка при добавлении книги с некорректными данными, но ее не было")
	}

	newBooks, err = book.AddBook(newBooks, "Название", "", "2000", 100)
	if err == nil {
		t.Error("Ожидалась ошибка при добавлении книги с некорректными данными, но ее не было")
	}

	newBooks, err = book.AddBook(newBooks, "Название", "Автор", "2000", 0)
	if err == nil {
		t.Error("Ожидалась ошибка при добавлении книги с некорректными данными, но ее не было")
	}
}

func TestLogFileCreation(t *testing.T) {
	// Создаем временную директорию для теста
	tempDir, err := ioutil.TempDir("", "testlogs")
	if err != nil {
		t.Fatalf("Ошибка создания временной директории: %v", err)
	}
	defer os.RemoveAll(tempDir) // Удаляем временную директорию после теста

	logFilePath := tempDir + "/app.log"
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatalf("Ошибка открытия лог файла: %v", err)
	}
	defer logFile.Close()
}

func TestLogMessage(t *testing.T) {
	// Перенаправляем лог в переменную
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)

	// Пример записи в лог
	log.Println("Тестовое сообщение")

	// Проверка, что лог содержит ожидаемое сообщение
	if !strings.Contains(logBuf.String(), "Тестовое сообщение") {
		t.Errorf("Лог не содержит ожидаемое сообщение")
	}
}

func TestLogError(t *testing.T) {
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)

	// Имитация функции, вызывающей ошибку
	logError("Ошибка при выполнении операции")

	if !strings.Contains(logBuf.String(), "Ошибка при выполнении операции") {
		t.Errorf("Лог не содержит ожидаемую ошибку")
	}
}

func logError(message string) {
	log.Printf("Ошибка: %s", message)
}

func TestAddBookLogs(t *testing.T) {
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)

	// Добавление книги
	_, err := book.AddBook([]book.Book{}, "Title", "Author", "2000", 100)

	if err != nil {
		t.Fatalf("Ошибка при добавлении книги: %v", err)
	}

	// Проверка, что лог содержит запись о добавлении книги
	if !strings.Contains(logBuf.String(), "Книга добавлена:") {
		t.Errorf("Лог не содержит ожидаемую запись о добавлении книги")
	}
}
