package model

import "time"

type MatterModel struct {
	ID        uint32    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Priority  int8      `json:"priority"`   // 优先级
	Date      string    `json:"date"`       // 设定的日期：2020-11-01
	Time      string    `json:"time"`       // 设定的时间：12:20
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UserID    uint32    `json:"user_id"`
}

func (m *MatterModel) Create() error {
	return DB.Self.Create(m).Error
}

func (m *MatterModel) Update() error {
	return DB.Self.Save(m).Error
}

func DeleteMatter(id uint32) error {
	matter := &MatterModel{ID: id}
	return DB.Self.Delete(matter).Error
}

func GetMatterByID(id uint32) (*MatterModel, error) {
	matter := &MatterModel{ID: id}
	d := DB.Self.First(matter)
	return matter, d.Error
}

func ListMatters(userID uint32) ([]*MatterModel, error) {
	var list []*MatterModel
	d := DB.Self.Where("user_id", userID).Find(&list)
	return list, d.Error
}
