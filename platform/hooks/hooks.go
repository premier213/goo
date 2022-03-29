package hook

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {

	if uuid, err := uuid.NewUUID(); err == nil {

		tx.Statement.SetColumn("ID", uuid)
	}

	return err
}
