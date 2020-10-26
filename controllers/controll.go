package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koko2824/goSample/db"
	"net/http"
	"strconv"
)

type Controller struct {}

// index
func (Controller) Index(c *gin.Context)  {
	todos := db.GetAllTodo()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"todos":todos,
	})
}

// addForm
func (Controller) AddForm(c *gin.Context)  {
	c.HTML(http.StatusOK, "add.html", gin.H{})
}

// create
func (Controller) Create(c *gin.Context)  {
	title := c.PostForm("title")
	text := c.PostForm("text")
	status := c.PostForm("status")
	db.AddTodo(title, text, status)
	c.Redirect(http.StatusFound, "/")
}

// detail
func (Controller) Detail(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("controllError: Detail()")
	}
	todo := db.GetOneTodo(id)
	c.HTML(http.StatusOK, "detail.html", gin.H{"todo":todo})
}

// update 作りかけ
func (Controller) Update(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil{
		panic(err)
	}
	todo := db.GetOneTodo(id)
	c.HTML(http.StatusOK, "update.html", gin.H{"todo":todo})
}

// Delete
func (Controller) Delete(c *gin.Context)  {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic("DeleteError")
	}
	db.DeleteTodo(id)
	c.Redirect(http.StatusFound, "/")
}