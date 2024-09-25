package model

type Department struct {
	ID       uint         `gorm:"primaryKey;autoIncrement"`
	Name     string       `gorm:"not null" json:"name"`
	ParentID *uint        `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"parentId,omitempty"`
	Parent   *Department  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Department `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}
