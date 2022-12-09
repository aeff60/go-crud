// CRUD API with mysql database

package main

import (
	"encoding/json" // ใช้สำหรับแปลงค่าที่เราส่งกลับไปให้เป็น json
	"fmt"           // ใช้สำหรับแสดงผล
	"log"           // ใช้สำหรับแสดง log
	"net/http"      // ใช้สำหรับสร้าง server
	"strconv"       // ใช้สำหรับแปลงค่า string เป็น int

	_ "github.com/go-sql-driver/mysql" // ใช้สำหรับเชื่อมต่อกับ mysql
	"github.com/gorilla/mux"           // ใช้สำหรับสร้าง router
)

// Book struct (Model)
type Book struct { // สร้าง struct ขึ้นมาเพื่อเก็บค่าที่จะส่งกลับไปให้ client
	ID     int     `json:"id"`     // สร้าง field ขึ้นมาเพื่อเก็บค่า id และกำหนดให้เป็น json
	Isbn   string  `json:"isbn"`   // สร้าง field ขึ้นมาเพื่อเก็บค่า isbn และกำหนดให้เป็น json
	Title  string  `json:"title"`  // สร้าง field ขึ้นมาเพื่อเก็บค่า title และกำหนดให้เป็น json
	Author *Author `json:"author"` // สร้าง field ขึ้นมาเพื่อเก็บค่า author และกำหนดให้เป็น json และกำหนดให้เป็น pointer เพื่อเก็บค่าที่อยู่ใน memory แทนที่จะเก็บค่าที่อยู่ใน stack และเป็นการเรียกใช้ struct ที่อยู่ใน struct
}

// Author struct
type Author struct { // สร้าง struct ขึ้นมาเพื่อเก็บค่าที่จะส่งกลับไปให้ client
	Firstname string `json:"firstname"` // สร้าง field ขึ้นมาเพื่อเก็บค่า firstname และกำหนดให้เป็น json
	Lastname  string `json:"lastname"`  // สร้าง field ขึ้นมาเพื่อเก็บค่า lastname และกำหนดให้เป็น json
}

// Init books var as a slice Book struct
var books []Book // สร้างตัวแปร books ขึ้นมาเพื่อเก็บค่าที่จะส่งกลับไปให้ client

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) { // สร้าง function ขึ้นมาเพื่อเรียกใช้งาน และส่งค่ากลับไปให้ client โดยใช้ http.ResponseWriter และ http.Request เป็น parameter ของ function นี้
	w.Header().Set("Content-Type", "application/json") // เรียกใช้ function Set ของ Header ของ http.ResponseWriter และกำหนดให้เป็น json
	json.NewEncoder(w).Encode(books)                   // ใช้  NewEncoder เพื่อเข้ารหัส json และส่งค่ากลับไปให้ client
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) { // สร้าง function ขึ้นมาเพื่อเรียกใช้งาน และส่งค่ากลับไปให้ client โดยใช้ http.ResponseWriter และ http.Request เป็น parameter ของ function นี้
	w.Header().Set("Content-Type", "application/json") // เรียกใช้ function Set ของ Header ของ http.ResponseWriter และกำหนดให้เป็น json
	params := mux.Vars(r)                              // สร้างตัวแปร params ขึ้นมาเพื่อเก็บค่าที่ส่งมาจาก client โดยใช้ mux.Vars และ http.Request เป็น parameter ของ function นี้
	// โดย mux.Vars จะเป็นการเรียกใช้ค่าที่ส่งมาจาก client โดยใช้ key ของค่าที่ส่งมา
	fmt.Println(params["id"])    // แสดงค่าที่ส่งมาจาก client โดยใช้ key ของค่าที่ส่งมา
	for _, item := range books { // ใช้ for loop เพื่อเรียกใช้ค่าที่ส่งมาจาก client โดยใช้ key ของค่าที่ส่งมา
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID { // ใช้ if statement เพื่อเช็คว่าค่าที่ส่งมาจาก client ตรงกับค่าใน database หรือไม่ โดยใช้ key ของค่าที่ส่งมา และเปลี่ยนค่าที่ส่งมาจาก client ให้เป็น int ก่อน
			json.NewEncoder(w).Encode(item) // ใช้  NewEncoder เพื่อเข้ารหัส json และส่งค่ากลับไปให้ client
			return                          // ใช้ return เพื่อออกจาก function
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // ใช้  NewEncoder เพื่อเข้ารหัส json และส่งค่ากลับไปให้ client โดยใช้ค่าว่าง และเป็นการส่งค่ากลับไปให้ client ว่าไม่พบค่าที่ต้องการ โดยใช้ key ของค่าที่ส่งมา
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) { // สร้าง function ขึ้นมาเพื่อเรียกใช้งาน และส่งค่ากลับไปให้ client โดยใช้ http.ResponseWriter และ http.Request เป็น parameter ของ function นี้

	w.Header().Set("Content-Type", "application/json") // ตั้งค่า header ให้เป็น json

	var book Book // สร้างตัวแปร book ขึ้นมาเพื่อเก็บค่าที่ส่งมาจาก client

	_ = json.NewDecoder(r.Body).Decode(&book) // แปลงค่าที่ส่งมาจาก client ให้เป็น json แล้วเก็บไว้ในตัวแปร book

	books = append(books, book) // เพิ่มค่าที่ได้จาก client ไปเก็บไว้ในตัวแปร books ที่เป็น slice ของ Book struct ที่เราสร้างไว้ แล้วส่งค่ากลับไปให้ client ที่เรียกใช้ api นี้

	json.NewEncoder(w).Encode(book) // ส่งค่ากลับไปให้ client ที่เรียกใช้ api นี้ โดยแปลงค่าที่เราส่งกลับไปให้เป็น json ก่อน แล้วส่งไป
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) { // สร้าง function ขึ้นมาเพื่อเรียกใช้งาน และส่งค่ากลับไปให้ client โดยใช้ http.ResponseWriter และ http.Request เป็น parameter ของ function นี้
	w.Header().Set("Content-Type", "application/json") // ตั้งค่า header ให้เป็น json
	params := mux.Vars(r)                              // สร้างตัวแปร params ขึ้นมาเพื่อเก็บค่าที่ส่งมาจาก client โดยใช้ mux.Vars(r) และส่งค่าที่ส่งมาจาก client ไปเก็บไว้ในตัวแปร params
	for index, item := range books {                   // วนลูปเพื่อเช็คค่าที่ส่งมาจาก client ว่ามีค่าที่ตรงกับค่าในตัวแปร books หรือไม่
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID { // ถ้ามีค่าที่ตรงกัน ให้เก็บค่าที่ส่งมาจาก client ไว้ในตัวแปร item.ID แล้วเช็คว่าค่าที่ส่งมาจาก client ตรงกับค่าในตัวแปร item.ID หรือไม่
			books = append(books[:index], books[index+1:]...) // ถ้ามีค่าที่ตรงกัน ให้เอาค่าที่ตรงกันออกจากตัวแปร books โดยใช้ append และเก็บค่าที่เหลือไว้ในตัวแปร books
			var book Book                                     // สร้างตัวแปร book ขึ้นมาเพื่อเก็บค่าที่ส่งมาจาก client
			_ = json.NewDecoder(r.Body).Decode(&book)         // แปลงค่าที่ส่งมาจาก client ให้เป็น json แล้วเก็บไว้ในตัวแปร book
			book.ID = item.ID                                 // เก็บค่า id ที่ส่งมาจาก client ไว้ในตัวแปร book.ID
			books = append(books, book)                       // เก็บค่าที่ส่งมาจาก client ไว้ในตัวแปร books
			json.NewEncoder(w).Encode(book)                   // ส่งค่ากลับไปให้ client ที่เรียกใช้ api นี้ โดยแปลงค่าที่เราส่งกลับไปให้เป็น json ก่อน แล้วส่งไป
			return
		}
	}
	json.NewEncoder(w).Encode(books) // ส่งค่ากลับไปให้ client ที่เรียกใช้ api นี้ โดยแปลงค่าที่เราส่งกลับไปให้เป็น json ก่อน แล้วส่งไป
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // กำหนดให้ค่าที่ส่งกลับไปเป็น json
	params := mux.Vars(r)                              // สร้างตัวแปร params เพื่อเก็บค่าที่ส่งมาจาก client
	for index, item := range books {                   // วนลูปเพื่อเช็คค่าที่ส่งมาจาก client ว่ามีค่าที่ตรงกับค่าในตัวแปร books หรือไม่
		if item.ID, _ = strconv.Atoi(params["id"]); item.ID == item.ID { // ถ้ามีค่าที่ตรงกัน ให้เก็บค่าที่ส่งมาจาก client ไว้ในตัวแปร item.ID แล้วเช็คว่าค่าที่ส่งมาจาก client ตรงกับค่าในตัวแปร item.ID หรือไม่
			books = append(books[:index], books[index+1:]...) // ถ้ามีค่าที่ตรงกัน ให้เอาค่าที่ตรงกันออกจากตัวแปร books โดยใช้ append และเก็บค่าที่เหลือไว้ในตัวแปร books
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: 1, Isbn: "448743", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: 2, Isbn: "448744", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
