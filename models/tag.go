package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func TagExistsByName(name string) (flag bool) {
	var tmp Tag
	db.Select("id").Where("name = ?", name).First(&tmp)
	if tmp.ID > 0 {
		flag = true
	}
	return flag
}

func TagExistsById(id int) (flag bool) {
	var tmp Tag
	db.Select("id").Where("id = ?", id).First(&tmp)
	if tmp.ID > 0 {
		flag = true
	}
	return flag
}

func AddTag(name string, state int, createBy string) {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})
}

func EditTag(tag *Tag, id int) {
	db.Model(&tag).Where("id = ?", id).Update(&tag)
}

func DeleteTag(id int) {
	db.Where("id = ?", id).Delete(&Tag{})
}

func (tag *Tag) BeforeCreate(scpoe *gorm.Scope) error {
	scpoe.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
