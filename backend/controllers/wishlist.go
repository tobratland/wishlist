package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tobratland/wishlist/backend/models"
	"gorm.io/gorm"
)

// CreateWishlistInput defines the input for creating a wishlist
type CreateWishlistInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// CreateWishlist handles the creation of a new wishlist
func CreateWishlist(c *gin.Context) {
	var input CreateWishlistInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	wishlist := models.Wishlist{
		ID:          uuid.New().String(),
		UserID:      userID.(string),
		Title:       input.Title,
		Description: input.Description,
	}

	if err := models.DB.Create(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating wishlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"wishlist": wishlist})
}

// GetWishlist retrieves the details of a specific wishlist
func GetWishlist(c *gin.Context) {
	wishlistID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var wishlist models.Wishlist
	if err := models.DB.Preload("Items").Where("id = ?", wishlistID).First(&wishlist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching wishlist"})
		return
	}

	if wishlist.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Prepare response without purchaser identities
	items := make([]gin.H, len(wishlist.Items))
	for i, item := range wishlist.Items {
		items[i] = gin.H{
			"id":          item.ID,
			"name":        item.Name,
			"description": item.Description,
			"purchased":   item.Purchased,
			"created_at":  item.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          wishlist.ID,
		"title":       wishlist.Title,
		"description": wishlist.Description,
		"created_at":  wishlist.CreatedAt,
		"items":       items,
	})
}

// ShareWishlistInput defines the input for sharing a wishlist
type ShareWishlistInput struct {
	// Additional fields can be added if needed
}

// ShareWishlist generates a shareable link for the wishlist
func ShareWishlist(c *gin.Context) {
	wishlistID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

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
		c.JSON(http.StatusForbidden, gin.H{"error": "Only the creator can share the wishlist"})
		return
	}

	// Generate a unique share token
	shareToken := uuid.New().String()

	// Save the share token with association to the wishlist
	// For simplicity, assume we have a Share model (not previously defined)
	share := models.Share{
		ID:         uuid.New().String(),
		WishlistID: wishlistID,
		Token:      shareToken,
	}

	if err := models.DB.Create(&share).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating shareable link"})
		return
	}

	// Construct the shareable link
	shareableLink := "http://yourfrontend.com/shared/" + shareToken

	c.JSON(http.StatusOK, gin.H{"shareable_link": shareableLink})
}
