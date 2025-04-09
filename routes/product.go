package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"olxkz/config"
	"olxkz/models"
	"strconv"
)

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 100
	}
	return i
}

func RegisterProductRoutes(r *gin.Engine) {
	r.GET("/products", GetProducts)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)    // Обновление продукта
	r.DELETE("/products/:id", DeleteProduct) // Удаление продукта
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	// Получение query параметров
	categoryID := c.Query("category_id")
	limitParam := c.DefaultQuery("limit", "10")
	pageParam := c.DefaultQuery("page", "1")

	limit := toInt(limitParam)
	page := toInt(pageParam)

	if limit <= 0 {
		limit = 2
	}
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	query := config.DB

	// Фильтрация по категории
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// Применение лимита и offset (пагинация)
	query = query.Limit(limit).Offset(offset).Find(&products)

	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	var existingProduct models.Product
	if err := config.DB.First(&existingProduct, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден!"})
		return
	}

	// Обновление
	config.DB.Model(&existingProduct).Updates(product)
	c.JSON(http.StatusOK, existingProduct)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Продукт не найден!"})
		return
	}

	// Удаление
	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Продукт удален"})
}
