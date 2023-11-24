package controller

import (
	"github.com/alyzsa/FinPro4/database"
	"github.com/alyzsa/FinPro4/entity"
	"github.com/alyzsa/FinPro4/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryHandlerImpl struct{}

func NewCategoryHandlerImpl() CategoryHandler {
	return &CategoryHandlerImpl{}
}

func (s *CategoryHandlerImpl) CategoryCreate(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	category := entity.Category{}
	category.Sold_Product_Amount = 0

	if contentType == appJSON {
		c.ShouldBindJSON(&category)
	} else {
		c.ShouldBind(&category)
	}

	if err := category.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err := db.Debug().Create(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                  category.ID,
		"type":                category.Type,
		"sold_product_amount": category.Sold_Product_Amount,
		"created_at":          category.CreatedAt,
	})
}

func (s *CategoryHandlerImpl) CategoryGet(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	var Category []entity.Category

	if contentType == appJSON {
		c.ShouldBindJSON(&Category)
	} else {
		c.ShouldBind(&Category)
	}

	err := db.Preload("Product").Find(&Category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Category)
}

func (s *CategoryHandlerImpl) CategoryUpdate(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType

	category := entity.Category{}

	categoryID, _ := strconv.Atoi(c.Param("categoryID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&category)
	} else {
		c.ShouldBind(&category)
	}

	category.ID = uint(categoryID)

	err := db.Model(&category).Where("id = ?", categoryID).Update("type", category.Type).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":                  category.ID,
		"type":                category.Type,
		"sold_product_amount": category.Sold_Product_Amount,
		"created_at":          category.CreatedAt,
	})
}

func (s *CategoryHandlerImpl) CategoryDelete(c *gin.Context) {
	var db = database.GetDB()
	contentType := helper.GetContentType(c)

	category := entity.Category{}

	categoryID, _ := strconv.Atoi(c.Param("categoryID"))

	if contentType == appJSON {
		c.ShouldBindJSON(&category)
	} else {
		c.ShouldBind(&category)
	}

	category.ID = uint(categoryID)

	err := db.Model(&category).Where("id = ?", categoryID).Delete(&category).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
