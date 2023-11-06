package database

import (
	"context"
	"dsystem/backend/pck/helper"
	"dsystem/backend/pck/types"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {
	dbUrl := "postgres://TodoManager:password@127.0.0.1:5432/TodoDB"
	newpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		fmt.Println("Error while creating db pool")
		return
	}
	Pool = newpool
}

func GetItem(id string) types.TodoItem {
	var item types.TodoItem
	cmd := "SELECT * FROM \"Items\" WHERE \"Id\"=$1"
	row, err := Pool.Query(context.Background(), cmd, id)
	if err != nil {
		return item
	}
	for row.Next() {
		row.Scan(&item.Id, &item.Title, &item.Description, &item.Created, &item.IsDone, &item.Image, &item.Color)
	}
	return item
}

func GetAllItemsShort() []types.TodoItemShort {
	var items []types.TodoItemShort
	cmd := "SELECT \"Id\", \"Title\", \"Color\" FROM \"Items\" ORDER BY \"IsDone\" ASC"
	row, err := Pool.Query(context.Background(), cmd)
	if err != nil {
		return items
	}
	for row.Next() {
		var item types.TodoItemShort
		row.Scan(&item.Id, &item.Title, &item.Color)
		items = append(items, item)
	}
	return items
}

func InsertItem(c *gin.Context) bool {
	var item types.TodoItem
	item.Id = helper.GenerateRandomId(32)
	item.Title = c.PostForm("title")
	item.Description = c.PostForm("desc")
	item.Created = time.Now()
	item.IsDone = false
	item.Image = c.PostForm("img")
	item.Color = c.PostForm("color")
	if len(item.Image) <= 2 {
		item.Image = ""
	}
	return insertIntoDatabase(item)
}

func insertIntoDatabase(item types.TodoItem) bool {
	cmd := "INSERT INTO \"Items\"(\"Id\", \"Title\", \"Description\", \"Created\", \"IsDone\", \"Image\", \"Color\") VALUES($1,$2,$3,$4,$5,$6,$7)"
	_, err := Pool.Exec(context.Background(), cmd, item.Id, item.Title, item.Description, item.Created, item.IsDone, item.Image, item.Color)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func DeleteItem(c *gin.Context) bool {
	id := c.Param("id")
	cmd := "DELETE FROM \"Items\" WHERE \"Id\"=$1"
	_, err := Pool.Exec(context.Background(), cmd, id)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
