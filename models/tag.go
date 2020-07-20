package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(offset int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(offset).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) (bool, error) {
	var (
		count int
		err   error
	)
	err = db.Model(&Tag{}).Select("count(*)").Where("name = ?", name).Count(&count).Error
	if err == nil {
		if count > 0 {
			return true, err
		}
	}
	return false, err
}

func ExistTagById(id int) (bool, error) {
	var (
		count int
		err   error
	)
	err = db.Model(&Tag{}).Select("count(*)").Where("id = ?", id).Count(&count).Error
	if err == nil {
		if count > 0 {
			return true, err
		}
	}
	return false, err
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_on", time.Now().Unix())
	scope.SetColumn("modified_on", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("modified_on", time.Now().Unix())
	return nil
}

func AddTag(name string, state int, createdBy string) error {
	err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
	return err
}

func EditTag(id int, data interface{}) error {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	return err
}

func DeleteTag(id int) error {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	return err
}
