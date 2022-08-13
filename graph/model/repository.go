package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

type BioRepository interface {
	CreatePerson(bioInput *AddBio) (*BioData, error)
	UpdatePerson(bioInput *AddBio, age int) error
	DeletePerson(age int) error
	GetOnePerson(age int) (*BioData, error)
	GetAllPersons() ([]*BioData, error)
}

type BioService struct {
	db *gorm.DB
}

func Database() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&BioData{})

	if err != nil {
		log.Fatal("failed to migrate")

	}
	return db
}

func (b *BioService) CreateBook(input *AddBio) (*BioData, error) {
	bio := &BioData{
		Name: input.Name,
		Age:  input.Age,
	}

	err := b.db.Create(&bio).Error

	return bio, err
}

func (b *BioService) UpdateBook(input *AddBio, age int) error {
	bio := &BioData{
		Name: input.Name,
		Age:  input.Age,
	}
	err := b.db.Model(&bio).Where("age = ?", age).Updates(bio).Error
	return err

}

func (b *BioService) GetOneBook(age int) (*BioData, error) {
	bio := &BioData{}
	err := b.db.Where("age = ?", age).First(bio).Error
	return bio, err
}

func (b *BioService) GetAllBooks() ([]*BioData, error) {
	var bio []*BioData
	err := b.db.Find(&bio).Error
	return bio, err
}

func (b *BioService) DeleteBook(age int) error {
	bio := &BioData{}
	err := b.db.Delete(bio, age).Error
	return err
}
