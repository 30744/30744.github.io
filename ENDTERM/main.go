package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type server struct {
	db *sql.DB
}

func dbConnect() server {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}

	s := server{db: db}

	return s
}

func (s *server) selectBooks() []Book {
	rows, err := s.db.Query("SELECT id, title, author, description FROM books")
	if err != nil {
		log.Fatal(err)
	}

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	return books
}

func (s *server) selectBookByID(id int) Book {
	row := s.db.QueryRow("SELECT id, title, author, description FROM books WHERE id=?", id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		log.Fatal(err)
	}

	return book
}

func (s *server) createBook(title string, author string, description string) int64 {
	res, err := s.db.Exec("INSERT INTO books (title, author, description) VALUES (?, ?, ?)", title, author, description)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (s *server) updateBookByID(id int, title string, author string, description string) int64 {
	res, err := s.db.Exec("UPDATE books SET title=?, author=?, description=? WHERE id=?", title, author, description, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}

func (s *server) deleteBookByID(id int) int64 {
	res, err := s.db.Exec("DELETE FROM books WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}

func (s *server) allBooksHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/books.html")
	if err != nil {
		log.Fatal(err)
	}

	books := s.selectBooks()
	err = t.Execute(w, books)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) createBookHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	id := s.createBook(title, author, description)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) editBookHandle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	book := s.selectBookByID(id)

	t, err := template.ParseFiles("./templates/edit-book.html")
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(w, book)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *server) updateBookHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	description := r.FormValue("description")

	s.updateBookByID(id, title, author, description)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (s *server) deleteBookHandle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	s.deleteBookByID(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	s := dbConnect()

	http.HandleFunc("/", s.allBooksHandle)
	http.HandleFunc("/create", s.createBookHandle)
	http.HandleFunc("/edit", s.editBookHandle)
	http.HandleFunc("/update", s.updateBookHandle)
	http.HandleFunc("/delete", s.deleteBookHandle)

	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	ID       int
	Title    string
	Author   string
	Category string
}

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT, category TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/add", addBook)
	http.HandleFunc("/update", updateBook)
	http.HandleFunc("/delete", deleteBook)

	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Category)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	err = tmpl.Execute(w, books)
	if err != nil {
		log.Fatal(err)
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	title := r.FormValue("title")
	author := r.FormValue("author")
	category := r.FormValue("category")

	_, err = db.Exec("INSERT INTO books(title, author, category) VALUES (?, ?, ?)", title, author, category)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	category := r.FormValue("category")

	_, err = db.Exec("UPDATE books SET title = ?, author = ?, category = ? WHERE id = ?", title, author, category, id)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
*/
