package database

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Messages struct {
	ID          uint64    `gorm:"primaryKey; column:message_id; index; AUTO_INCREMENT"`
	MessageFrom uint64    `gorm:"column:message_from"`
	MessageTo   uint64    `gorm:"column:message_to"`
	Content     string    `gorm:"column:content"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdateAt    time.Time `gorm:"column:created_at"`
}

type MessageDataAccessor interface {
	CreateMessage(ctx context.Context, message *Messages) (*Messages, error)
	EditMessage(ctx context.Context, message *Messages) (*Messages, error)
	DeleteMessage(ctx context.Context, id uint64) error
}

type messageDataAccessor struct {
	database *gorm.DB
}

func (m messageDataAccessor) CreateMessage(ctx context.Context, message *Messages) (*Messages, error) {
	tx := m.database.Create(message)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return message, nil
}

func (m messageDataAccessor) DeleteMessage(ctx context.Context, id uint64) error {
	tx := m.database.Raw("update messages set message_from = null where ID = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m messageDataAccessor) EditMessage(ctx context.Context, message *Messages) (*Messages, error) {
	// tx := m.database.Raw(
	// 	"update messages set content = ?, update_at = ? where id = ?",
	// 	message.Content,
	// 	time.Now(),
	// 	message.ID,
	// )

	tx := m.database.Save(&message)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return message, nil
}

func InitializeMessageDataAccessor(database *gorm.DB) MessageDataAccessor {
	return &messageDataAccessor{
		database: database,
	}
}
