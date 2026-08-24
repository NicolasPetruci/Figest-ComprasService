package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/NicolasPetruci/Figest-ComprasService/internal/database"
	"github.com/NicolasPetruci/Figest-ComprasService/internal/models"
	"github.com/gin-gonic/gin"
)

func CreatePurchase(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var purchase models.Purchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase.TotalPrice = float64(purchase.Quantity) * purchase.UnitPrice
	purchase.UserID = userID

	if err := database.DB.Create(&purchase).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create purchase"})
		return
	}

	c.JSON(http.StatusCreated, purchase)
}

func GetPurchases(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var purchases []models.Purchase
	if err := database.DB.Where("user_id = ?", userID).Find(&purchases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch purchases"})
		return
	}

	c.JSON(http.StatusOK, purchases)
}

func GetPurchase(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := c.Param("id")

	var purchase models.Purchase
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&purchase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func UpdatePurchase(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := c.Param("id")

	var purchase models.Purchase
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&purchase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		return
	}

	var updateData models.Purchase
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateData.Quantity != 0 || updateData.UnitPrice != 0 {
		qty := purchase.Quantity
		if updateData.Quantity != 0 {
			qty = updateData.Quantity
		}
		price := purchase.UnitPrice
		if updateData.UnitPrice != 0 {
			price = updateData.UnitPrice
		}
		updateData.TotalPrice = float64(qty) * price
	}

	if err := database.DB.Model(&purchase).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update purchase"})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func DeletePurchase(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	id := c.Param("id")

	var purchase models.Purchase
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&purchase).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		return
	}

	if err := database.DB.Delete(&purchase).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete purchase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase deleted"})
}

func GetPurchasesSummary(c *gin.Context) {
	userID := getUserID(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	monthStr := c.Query("month")
	yearStr := c.Query("year")

	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	endOfMonth = time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, time.UTC)

	type SummaryResult struct {
		SupplierID uint    `json:"supplier_id"`
		TotalSpent float64 `json:"total_spent"`
	}
	var results []SummaryResult

	if err := database.DB.Model(&models.Purchase{}).
		Select("supplier_id, sum(total_price) as total_spent").
		Where("user_id = ? AND purchase_date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
		Group("supplier_id").
		Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch summary"})
		return
	}

	c.JSON(http.StatusOK, results)
}
