package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tobratland/wishlist/backend/models"
	"gorm.io/gorm"
)

// AddItemInput defines the input for adding a new item
type AddItemInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// AddItem handles adding a new item to a wishlist
func AddItem(c *gin.Context) {
	wishlistID := c.Param("id")
	var input AddItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Verify that the user owns the wishlist
	var wishlist models.Wishlist
	if err := models.DB.Where("id = ?", wishlistID).First(&wishlist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching wishlist"})
		return
	}

	if wishlist.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only the creator can add items"})
		return
	}

	item := models.Item{
		ID:          uuid.New().String(),
		WishlistID:  wishlistID,
		Name:        input.Name,
		Description: input.Description,
	}

	if err := models.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

// PurchaseItem handles marking an item as purchased
func PurchaseItem(c *gin.Context) {
	itemID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var item models.Item
	if err := models.DB.Where("id = ?", itemID).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching item"})
		return
	}

	// Check if the item is already purchased by this user
	var existingPurchase models.Purchase
	if err := models.DB.Where("item_id = ? AND user_id = ?", itemID, userID.(string)).First(&existingPurchase).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item already purchased by you"})
		return
	}

	// Create a purchase record
	purchase := models.Purchase{
		ID:          uuid.New().String(),
		ItemID:      itemID,
		UserID:      userID.(string),
		PurchasedAt: time.Now().Format(time.RFC3339),
	}

	if err := models.DB.Create(&purchase).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error recording purchase"})
		return
	}

	// Update the item's purchased status
	// Depending on requirements, you might want to allow multiple purchases by different users
	// Here, we'll set purchased to true if at least one purchase exists
	models.DB.Model(&item).Update("purchased", true)

	c.JSON(http.StatusOK, gin.H{"message": "Item marked as purchased"})
}
