package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Firstname string         `json:"firstname" gorm:"not null"`
	Lastname  string         `json:"lastname" gorm:"not null"`
	Username  string         `json:"username" gorm:"not null,unique"`
	Email     string         `json:"email" gorm:"not null,unique"`
	Password  string         `json:"-" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdateAt  time.Time      `json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Account struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	UserID    int            `json:"user_id" gorm:"references:users.id"`
	Balance   float64        `json:"balance" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdateAt  time.Time      `json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Product struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	CategoryID  int            `json:"category_id" gorm:"references:categories.id"`
	IsActive    bool           `json:"is_active" gorm:"not null"`
	CreatedAt   time.Time      `json:"-"`
	UpdateAt    time.Time      `json:"-"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Category struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	IsActive    bool           `json:"is_active" gorm:"not null"`
	CreatedAt   time.Time      `json:"-"`
	UpdateAt    time.Time      `json:"-"`
	DeleteAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

type Transaction struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	From      int            `json:"from_id"`
	To        int            `json:"to_id"`
	Comment   string         `json:"comment" gorm:"not null"`
	Amount    float64        `json:"amount"`
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"created_at"`
	UpdateAt  time.Time      `json:"updated_at"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Report struct {
	ID        int            `json:"id" gorm:"not null"`
	From      int            `json:"from_id,omitempty"`
	To        int            `json:"to_id,omitempty"`
	Limit     int            `json:"limit,omitempty"`
	Page      int            `json:"page,omitempty"`
	Type      string         `json:"type,omitempty"`
	CreatedAt time.Time      `json:"-"`
	UpdateAt  time.Time      `json:"-"`
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
