package main

import (
	"errors"
	"fmt"

	// "log"
	"net/http"
	"strconv"
	"time"
	"viettran2/config"
	"viettran2/controller"
	"viettran2/helper"
	"viettran2/model"
	"viettran2/repository"
	"viettran2/router"
	"viettran2/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	// "github.com/go-playground/validator"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type ToDoItem struct {
	ID          int       `json:"id" gorm:"column:id"`
	Description string    `json:"description" gorm:"column:description"`
	Title       string    `json:"title" gorm:"column:title"`
	Status      string    `json:"status" gorm:"column:status"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"column:updated_at"`
}
func (ToDoItem) TableName() string{return "todo_items"}
type ToDoItemCreation struct {
    ID          int    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description" gorm:"column:description;"`
	Title       string    `json:"title" gorm:"column:title;"`
	Status      string    `json:"status" gorm:"column:status;"`

}
func (ToDoItemCreation) TableName() string{return ToDoItem{}.TableName()}
// ... (import statements and ToDoItem struct definition)
type ToDoItemUpdate struct {

	Description string    `json:"description" gorm:"column:description;"`
	Title       string    `json:"title" gorm:"column:title;"`
	Status      string    `json:"status" gorm:"column:status;"`

}
func (ToDoItemUpdate) TableName() string{return ToDoItem{}.TableName()}
func main() {
	log.Info().Msg("Started Server!")
    db := config.DatabaseMySqlConnection()
    
	validate := validator.New()
 
    db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)

	// Router
	routes := router.NewRouter(tagsController)
    
 

    // r := gin.Default()
    // v1 := r.Group("/v1")
    // v1.GET("/items/:id", getItem(db))   
    // v1.PATCH("/items/:id", updateItem(db))   
    // v1.GET("/items", listItems(db))   
    // v1.POST("/items", createItem(db)) 
   
    // r.Run(":8080")
    server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}



func createItem(db *gorm.DB) func(c *gin.Context) {
    return func(c *gin.Context) {
        var data ToDoItemCreation
        if err := c.ShouldBindJSON(&data); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

   
        // Create the item in the database
        if err := db.Create(&data).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item in the database"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Item created successfully", "data": data.ID})
    }
}


func listItems(db *gorm.DB) func(c *gin.Context) {
    return func(c *gin.Context) {
        var items []ToDoItem

        // Retrieve all items from the database
        if err := db.Find(&items).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items from the database"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Items fetched successfully", "data": items})
    }
}
func getItem(db *gorm.DB) func(c *gin.Context) {
    return func(c *gin.Context) {
        var item ToDoItem

        // Get the item ID from the request parameters
        itemIDStr := c.Param("id")

        // Convert the itemID string to an integer
        itemID, err := strconv.Atoi(itemIDStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
            return
        }

        // Retrieve the item from the database using the ID
        if err := db.First(&item, itemID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Item fetched successfully", "data": item})
    }
}
func updateItem(db *gorm.DB) func(c *gin.Context) {
    return func(c *gin.Context) {
        var item ToDoItemUpdate

        // Get the item ID from the request parameters
        itemIDStr := c.Param("id")

        // Convert the itemID string to an integer
        itemID, err := strconv.Atoi(itemIDStr)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
            return
        }
        if err := c.ShouldBindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        // Retrieve the item from the database using the ID
        // if err := db.Where(query: "itemID = ?", itemID).Update(&item).Error; err != nil {
        //     c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        //     return
        // }


// test ------
fmt.Printf("this is item update 1111111")
if err := db.Where("itemID = ?", itemID).First(&item).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        return
    }
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve item"})
    return
}
fmt.Printf("this is item update ")
fmt.Printf("%+v\n", item)
// ------------
        // Bind the JSON data from the request to the item struct
        if err := c.BindJSON(&item); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
            return
        }

        // Update the item in the database
        if err := db.Save(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Item updated successfully", "data": item})
    }
}


// func updateItem(c *gin.Context) {
// 	id := c.Param("id")
// 	// Convert id to integer and find the corresponding item
// 	// (Error handling omitted for brevity)
// 	itemID := convertToInt(id)
// 	for i := range todoItems {
// 		if todoItems[i].ID == itemID {
// 			var updatedItem ToDoItem
// 			if err := c.ShouldBindJSON(&updatedItem); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			updatedItem.ID = itemID
// 			updatedItem.CreatedAt = todoItems[i].CreatedAt
// 			updatedItem.UpdatedAt = time.Now()
// 			todoItems[i] = updatedItem

// 			c.JSON(http.StatusOK, updatedItem)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// }

// func deleteItem(c *gin.Context) {
// 	id := c.Param("id")
// 	// Convert id to integer and find the index of the corresponding item
// 	// (Error handling omitted for brevity)
// 	itemID := convertToInt(id)
// 	for i, item := range todoItems {
// 		if item.ID == itemID {
// 			todoItems = append(todoItems[:i], todoItems[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
// }

func convertToInt(s string) int {
	// Convert a string to an integer (Error handling omitted for brevity)
	return 0
}

// func generateNewID() int {
// 	// Generate a new unique ID for a new item (Error handling omitted for brevity)
// 	return len(todoItems) + 1
// }