package repository

import (
	"conversation-topic/database"
	"conversation-topic/model"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func InitUserRepository() *UserRepository {
	database.DBConn.AutoMigrate(&model.User{})

	var count int64
	database.DBConn.Model(&model.User{}).Count(&count)

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("kcolrehs"), 8)
		if err == nil {
			database.DBConn.Create(&model.User{
				Name:     "Eurus Holmes",
				Username: "admin",
				Password: string(hashedPassword),
			})
		}
	}

	return &UserRepository{DB: database.DBConn}
}

type UserRepository struct {
	DB *gorm.DB
}

func (repo UserRepository) ExistsById(id int) bool {
	var user model.User
	result := repo.DB.First(&user, id)
	return result.RowsAffected != 0
}

func (repo UserRepository) FindById(id int) *model.User {
	var user model.User

	result := repo.DB.Model(&model.User{}).
		First(&user, id)

	if result.RowsAffected == 0 {
		return nil
	}
	return &user
}

func (repo UserRepository) FindByUsername(username string) *model.User {
	var user model.User

	result := repo.DB.Model(&model.User{}).
		Where("username = ?", username).
		First(&user)

	if result.RowsAffected == 0 {
		return nil
	}
	return &user
}

func (repo UserRepository) Count() int64 {
	var count int64
	repo.DB.Model(&model.User{}).Count(&count)
	return count
}

func (repo UserRepository) List() []model.User {
	var users []model.User

	repo.DB.Model(&model.User{}).
		Find(&users)

	return users
}

func (repo UserRepository) ListPaged(page int, limit int) []model.User {
	var users []model.User

	repo.DB.Model(&model.User{}).
		Offset((page - 1) * limit).
		Limit(limit).Find(&users)

	return users
}

func (repo UserRepository) Search(search string) []model.User {
	var users []model.User

	repo.DB.Model(&model.User{}).
		Where("username LIKE ?", search+"%").
		Find(&users)

	return users
}

func (repo UserRepository) Create(user *model.User) {
	if repo.ExistsById(user.ID) {
		return
	}
	repo.DB.Create(user)
}

func (repo UserRepository) Update(user *model.User) {
	repo.DB.Save(user)
}

func (repo UserRepository) DeleteById(id int) {
	repo.DB.Delete(&model.User{}, id)
}
