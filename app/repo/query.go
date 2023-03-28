package repo

import (
	"book-gorm/app/entity"
	"book-gorm/pkg/database"
	"errors"
	"time"

	"gorm.io/gorm"
)

func GetSemuaData() ([]entity.BookGorm, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	books := []entity.BookGorm{}
	err = db.Find(&books).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.BookGorm{}, nil
		}
		return nil, err
	}

	return books, nil
}

func GetDataBerdasarkanID(idparam int) ([]entity.BookGorm, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	books := []entity.BookGorm{}
	err = db.First(&books, "id = $1", idparam).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.BookGorm{}, nil
		}
		return nil, err
	}

	return books, nil

}

func CreateBuku(title, author string) ([]entity.BookGorm, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	books := entity.BookGorm{
		Title:  title,
		Author: author,
	}

	err = db.Create(&books).Error
	if err != nil {
		return nil, err
	}

	return []entity.BookGorm{books}, nil

}

func UpdateBuku(id int, title string, author string) ([]entity.BookGorm, error) {
	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	books := entity.BookGorm{}
	var createdAt time.Time

	db.Select("created_at").First(&books, "id = ?", id).Scan(&createdAt)

	books.ID = id
	books.CreatedAt = createdAt
	err = db.Model(&books).Where("id = ?", id).Updates(entity.BookGorm{Title: title, Author: author}).Error
	if err != nil {
		return nil, err
	}

	return []entity.BookGorm{books}, nil

}

func DeleteBuku(idparam int) (string, error) {
	db, err := database.Connect()

	if err != nil {
		return "Gagal", err
	}

	err = db.Where("id = ?", idparam).Delete(&entity.BookGorm{}).Error
	if err != nil {
		return "Gagal", err
	}

	return "Berhasil", nil
}
