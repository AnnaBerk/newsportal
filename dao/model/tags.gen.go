// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	TagID    int32  `gorm:"column:tagId;primaryKey;autoIncrement:true" json:"tagId"`
	Title    string `gorm:"column:title;not null" json:"title"`
	StatusID int32  `gorm:"column:statusId;not null" json:"statusId"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
