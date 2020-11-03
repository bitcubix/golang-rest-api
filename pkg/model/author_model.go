package model

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint      `gorm:"primarykey;autoincrement"`
	Name      string    `json:"name"`
	Books     *[]Book   `json:"books" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Author) GetAll(db *gorm.DB) (*[]Author, error) {
	var err error
	var book Book
	var authors []Author

	err = db.Find(&authors).Error
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		ab, _ := book.GetByAuthorID(db, int(author.ID))
		author.Books = ab
	}

	return &authors, nil
}

func (a *Author) Save(db *gorm.DB) error {
	var err error

	err = db.Create(a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Author) GetByID(db *gorm.DB, id uint) error {
	var err error

	err = db.Where(&Author{ID: id}).First(a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Author) Update(db *gorm.DB, id int) error {
	var err error

	a.ID = uint(id)
	a.UpdatedAt = time.Now()
	err = db.Model(a).Where(id).Updates(a).Error
	if err != nil {
		return err
	}
	return nil
}

func (a *Author) Delete(db *gorm.DB, id int) error {
	var err error

	err = db.First(a, id).Delete(a).Error
	if err != nil {
		return err
	}
	return nil
}
