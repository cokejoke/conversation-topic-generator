package repository

import (
	"conversation-topic/database"
	"conversation-topic/model"
	"github.com/jinzhu/gorm"
)

func InitPackRepository() *PackRepository {
	database.DBConn.AutoMigrate(&model.Pack{})
	database.DBConn.AutoMigrate(&model.Topic{})

	return &PackRepository{DB: database.DBConn}
}

type PackRepository struct {
	DB *gorm.DB
}

func (repo PackRepository) ExistsById(id int) bool {
	var pack model.Pack
	result := repo.DB.First(&pack, id)
	return result.RowsAffected != 0
}

func (repo PackRepository) FindById(id int) *model.Pack {
	var pack model.Pack

	result := repo.DB.Model(&model.Pack{}).
		Preload("Topics").
		First(&pack, id)

	if result.RowsAffected == 0 {
		return nil
	}
	return &pack
}

func (repo PackRepository) Count() int64 {
	var count int64
	repo.DB.Model(&model.Pack{}).Count(&count)
	return count
}

func (repo PackRepository) List() []model.Pack {
	var packs []model.Pack

	repo.DB.Model(&model.Pack{}).
		Find(&packs)

	return packs
}

func (repo PackRepository) ListPaged(page int, limit int) []model.Pack {
	var packs []model.Pack

	repo.DB.Model(&model.Pack{}).
		Offset((page - 1) * limit).
		Limit(limit).Find(&packs)

	return packs
}

func (repo PackRepository) Search(search string) []model.Pack {
	var packs []model.Pack

	repo.DB.Model(&model.Pack{}).
		Where("name LIKE ?", search+"%").
		Find(&packs)

	return packs
}

func (repo PackRepository) Create(pack *model.Pack) {
	if repo.ExistsById(pack.ID) {
		return
	}

	repo.DB.Create(pack)
}

func (repo PackRepository) Update(pack *model.Pack) {
	repo.DB.Save(pack)
}

func (repo PackRepository) DeleteById(id int) {
	repo.DB.Delete(&model.Pack{}, id)
}
