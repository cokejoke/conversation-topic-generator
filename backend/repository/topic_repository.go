package repository

import (
	"conversation-topic/database"
	"conversation-topic/model"
	"github.com/jinzhu/gorm"
	"math/rand"
)

func InitTopicRepository() *TopicRepository {
	database.DBConn.AutoMigrate(&model.Topic{})

	return &TopicRepository{DB: database.DBConn}
}

type TopicRepository struct {
	DB *gorm.DB
}

func (repo TopicRepository) ExistsById(id int) bool {
	var topic model.Topic
	result := repo.DB.First(&topic, id)
	return result.RowsAffected != 0
}

func (repo TopicRepository) FindById(id int) *model.Topic {
	var topic model.Topic

	result := repo.DB.Model(&model.Topic{}).
		First(&topic, id)

	if result.RowsAffected == 0 {
		return nil
	}
	return &topic
}

func (repo TopicRepository) Count() int64 {
	var count int64
	repo.DB.Model(&model.Topic{}).Count(&count)
	return count
}

func (repo TopicRepository) List() []model.Topic {
	var topics []model.Topic

	repo.DB.Model(&model.Topic{}).
		Find(&topics)

	return topics
}

func (repo TopicRepository) ListPaged(page int, limit int) []model.Topic {
	var topics []model.Topic

	repo.DB.Model(&model.Topic{}).
		Offset((page - 1) * limit).
		Limit(limit).Find(&topics)

	return topics
}

func (repo TopicRepository) Search(search string) []model.Topic {
	var topics []model.Topic

	repo.DB.Model(&model.Topic{}).
		Where("topic LIKE ?", search+"%").
		Find(&topics)

	return topics
}

func (repo TopicRepository) Create(topic *model.Topic) {
	if repo.ExistsById(topic.ID) {
		return
	}
	repo.DB.Create(topic)
}

func (repo TopicRepository) Update(topic *model.Topic) {
	repo.DB.Save(topic)
}

func (repo TopicRepository) DeleteById(id int) {
	repo.DB.Delete(&model.Topic{}, id)
}

func (repo TopicRepository) RandomTopic() model.Topic {
	var topic model.Topic

	var count int
	repo.DB.Model(&model.Topic{}).Count(&count)
	randomOffset := 0
	if count > 1 {
		randomOffset = rand.Intn(count)
	}

	repo.DB.Model(&model.Topic{}).
		Offset(randomOffset).
		Limit(1).First(&topic)
	return topic
}
