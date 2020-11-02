package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID          uint      `gorm:"primarykey;autoincrement"`
	ISBN        string    `json:"isbn"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

//TODO validation

func (b *Book) GetAll(db *gorm.DB) (*[]Book, error) {
	var err error
	var books []Book

	err = db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return &books, nil
}

func (b *Book) Save(db *gorm.DB) error {
	var err error

	err = db.Create(b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) GetByID(db *gorm.DB, id uint) error {
	var err error

	err = db.Where(&Book{ID: id}).First(b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) GetByISBN(db *gorm.DB, isbn string) error {
	var err error

	err = db.Where(&Book{ISBN: isbn}).First(b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Update(db *gorm.DB, id int) error {
	var err error

	err = db.Model(b).Where(id).Updates(b).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *Book) Delete(db *gorm.DB, id int) error {
	var err error

	err = db.First(b, id).Delete(b).Error
	if err != nil {
		return err
	}
	return nil
}
