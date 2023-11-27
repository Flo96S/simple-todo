package endpoint

import (
	"backend/pck/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	list := database.GetAllItemsShort()
	c.JSON(http.StatusOK, list)
}

func GetTodoById(c *gin.Context) {
	item := database.GetItem(c.Param("id"))
	c.JSON(http.StatusOK, item)
}

func AddItem(c *gin.Context) {
	success := database.InsertItem(c)
	fmt.Println("Success: ", success)
	if success {
		c.JSON(http.StatusOK, "Done")
		return
	}
	c.JSON(http.StatusBadRequest, "Not successful")
}

func UpdateItem(c *gin.Context) {
	id := c.PostForm("id")
	success := database.Update(id, true)
	if success {
		c.JSON(http.StatusOK, "")
		return
	}
	c.JSON(http.StatusBadRequest, "")
}

func DeleteItem(c *gin.Context) {
	success := database.DeleteItem(c)
	if success {
		c.JSON(http.StatusOK, "Deleted")
		return
	}
	c.JSON(http.StatusBadRequest, "Still exists")
}

func ApiRoot(c *gin.Context) {
	c.JSON(http.StatusOK, "Working api - distributedsystems")
}

/* ------------------------------ */
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success:": database.Check(),
		"error":    database.ErrorCheck(),
	})
}
