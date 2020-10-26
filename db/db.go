package db

import (
	"github.com/koko2824/goSample/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)


// DB初期化
func Init() {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("dateBaseError: db.go Init()")
	}
	db.AutoMigrate(&models.Todo{})
	defer db.Close()
}

// DB追加
func AddTodo(title string, text string, status string) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("dateBaseError: db.go AddTodo()")
	}
	db.Create(&models.Todo{Title: title, Text: text, Status: status})
	defer db.Close()
}

// 更新
func UpdateTodo(id int, title string, text string, status string) {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("dateBaseError: db.go UpdateTdo()")
	}
	var todo models.Todo
	db.First(&todo, id)
	todo.Title = title
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

// 削除
func DeleteTodo(id int)  {
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("dateBaseError: db.go DeleteTodo()")
	}
	db.First(&models.Todo{})
	db.Delete(&models.Todo{})
	db.Close()
}

// DB全取得
func GetAllTodo()  []models.Todo{
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil {
		panic("dataBaseError: db.go GetAllTodo()")
	}
	var todos []models.Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

// DB一つ取得
func GetOneTodo(id int)  models.Todo{
	db, err := gorm.Open("sqlite3", "todo.sqlite3")
	if err != nil{
		panic("dataBaseError: db.go GetOneTodo()")
	}
	var todo models.Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
