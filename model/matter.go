package model

import (
	"time"
)

type MatterModel struct {
	ID        uint32    `json:"id" gorm:"column:id"`
	Title     string    `json:"title" gorm:"column:title"`
	Content   string    `json:"content" gorm:"column:content"`
	Priority  int8      `json:"priority" gorm:"column:priority"`     // 优先级
	State     int8      `json:"state" gorm:"column:state"`           // 状态：0/1 => 未完成/已完成
	Date      string    `json:"date" gorm:"column:date"`             // 设定的日期：2020-11-01
	Time      string    `json:"time" gorm:"column:time"`             // 设定的时间：12:20
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"` // 创建时间
	UserID    uint32    `json:"user_id" gorm:"column:user_id"`
}

func (*MatterModel) TableName() string {
	return "matter"
}

func (m *MatterModel) Create() error {
	return DB.Self.Create(m).Error
}

func (m *MatterModel) Update() error {
	return DB.Self.Save(m).Error
}

func DeleteMatter(id uint32) error {
	return DB.Self.Where("id = ?", id).Delete(MatterModel{}).Error
}

func GetMatterByID(id uint32) (*MatterModel, error) {
	matter := &MatterModel{ID: id}
	d := DB.Self.First(matter)
	return matter, d.Error
}

func ListMatters(userID uint32) ([]*MatterModel, error) {
	var list []*MatterModel
	d := DB.Self.Where("user_id = ?", userID).Find(&list)
	return list, d.Error
}

func MatterIsExisted(id uint32) (bool, error) {
	count := 0
	d := DB.Self.Table("matter").Where("id = ?", id).Count(&count)
	return count != 0, d.Error
}
