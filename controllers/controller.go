package controllers

import (
	"github.com/gin-gonic/gin"
	"TodoApp/models"
	"TodoApp/repositories"
	"gorm.io/gorm"
	"strconv"
	"TodoApp/TodoSort"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
    return &TodoController{db: db}
}

func (con TodoController) Add(c *gin.Context) {
	todo := models.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		c.Status(503)
        return
	}

	repo := repositories.NewTodoRepository(con.db)
	err = repo.Add(&todo)
	if err!= nil {
		c.Status(503)
        return
    }
	c.Status(204)
}

func (con TodoController) List(c *gin.Context) {
	Qtitle := c.DefaultQuery("Title","")
	repo := repositories.NewTodoRepository(con.db)
	TodoList, err := repo.List(Qtitle)
	if err!= nil {
        c.Status(503)
        return
    }
	c.JSON(200, gin.H{
		"list": TodoList,
	})
}

func (con TodoController) SortedList(c *gin.Context) {
	Qtitle := c.DefaultQuery("Title","")
	repo := repositories.NewTodoRepository(con.db)
	TodoList, err := repo.List(Qtitle)
	if err!= nil {
        c.Status(503)
        return
    }
	sortedTodos := TodoSort.QuickSortStart(TodoList)
	c.JSON(200, gin.H{
		"sortedList": sortedTodos,
	})
}

func (con TodoController) Update(c *gin.Context) {
	todo := models.Todo{}
	err := c.ShouldBindQuery(&todo)
	if err!= nil {
        c.Status(503)
        return
    }
	repo := repositories.NewTodoRepository(con.db)

	var existingTodo models.Todo
    if err := con.db.First(&existingTodo, todo.ID).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.Status(404)
        } else {
            c.Status(503)
        }
        return
    }

	err = repo.Update(&todo)
	if err!= nil {
		c.Status(503)
        return
    }
	c.Status(204)
}

func (con TodoController) Delete(c *gin.Context) {
	id := c.Query("id")
	intId, err := strconv.Atoi(id)

	var existingTodo models.Todo
    if err := con.db.First(&existingTodo, intId).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.Status(404)
        } else {
            c.Status(503)
        }
        return
    }

    if err != nil {
		c.Status(503)
		return
	}
	
	repo := repositories.NewTodoRepository(con.db)
	err = repo.Delete(intId)
	if err!= nil {
		c.Status(503)
        return
    }
	c.Status(204)
}