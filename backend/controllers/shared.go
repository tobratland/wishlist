package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobratland/wishlist/backend/models"
	"gorm.io/gorm"
)

// SharedWishlistResponse defines the structure of the shared wishlist response
type SharedWishlistResponse struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	CreatedAt   string       `json:"created_at"`
	Items       []SharedItem `json:"items"`
}

// SharedItem defines the structure of each item in the shared wishlist
type SharedItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Purchased   bool   `json:"purchased"`
	// Optionally, include who purchased if the requester is the purchaser themselves
	PurchasedBy string `json:"purchased_by,omitempty"` // Only set if the requester is the purchaser
}

// GetSharedWishlist handles accessing a wishlist via a shareable link
func GetSharedWishlist(c *gin.Context) {
	shareToken := c.Param("token")

	var share models.Share
	if err := models.DB.Where("token = ?", shareToken).First(&share).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invalid or expired share link"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching share link"})
		return
	}

	// Optionally, check if the share link has expired
	// if time.Now().After(share.ExpiresAt) {
	//     c.JSON(http.StatusGone, gin.H{"error": "Share link has expired"})
	//     return
	// }

	var wishlist models.Wishlist
	if err := models.DB.Preload("Items.Purchases").Where("id = ?", share.WishlistID).First(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching wishlist"})
		return
	}

	// Determine if the requester is authenticated
	var userID string
	var isAuthenticated bool
	if uid, exists := c.Get("userID"); exists {
		userID = uid.(string)
		isAuthenticated = true
	}

	// Prepare the response
	response := SharedWishlistResponse{
		ID:          wishlist.ID,
		Title:       wishlist.Title,
		Description: wishlist.Description,
		CreatedAt:   wishlist.CreatedAt,
		Items:       []SharedItem{},
	}

	for _, item := range wishlist.Items {
		sharedItem := SharedItem{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Purchased:   item.Purchased,
		}

		if isAuthenticated {
			// Check if the authenticated user has purchased this item
			for _, purchase := range item.Purchases {
				if purchase.UserID == userID {
					sharedItem.PurchasedBy = "You"
					break
				}
			}
		}

		response.Items = append(response.Items, sharedItem)
	}

	c.JSON(http.StatusOK, gin.H{"wishlist": response})
}
