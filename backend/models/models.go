package models

import (
	"log"
	"time"

	"github.com/tobratland/wishlist/backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type Wishlist struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      string `gorm:"type:uuid;not null" json:"user_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	CreatedAt   string `gorm:"autoCreateTime" json:"created_at"`
	Items       []Item `json:"items"`
}

type Item struct {
	ID          string     `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WishlistID  string     `gorm:"type:uuid;not null" json:"wishlist_id"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `json:"description"`
	Purchased   bool       `gorm:"default:false" json:"purchased"`
	CreatedAt   string     `gorm:"autoCreateTime" json:"created_at"`
	Purchases   []Purchase `json:"-"`
}

type Purchase struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	ItemID      string `gorm:"type:uuid;not null" json:"item_id"`
	UserID      string `gorm:"type:uuid;not null" json:"user_id"`
	PurchasedAt string `gorm:"autoCreateTime" json:"purchased_at"`
}

type Share struct {
	ID         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	WishlistID string    `gorm:"type:uuid;not null;index" json:"wishlist_id"`
	Token      string    `gorm:"unique;not null" json:"token"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt  time.Time `json:"expires_at"` // Optional
}

func ConnectDatabase() {
	dsn := config.GetDBConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
}

func Migrate() {
	err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Fatal("Failed to create uuid extension:", err)
	}

	err = DB.AutoMigrate(&User{}, &Wishlist{}, &Item{}, &Purchase{}, &Share{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
