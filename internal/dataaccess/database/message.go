package database

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Messages struct {
	ID          uint64    `gorm:"primaryKey; column:message_id; index; AUTO_INCREMENT"`
	MessageFrom *uint64   `gorm:"column:message_from"`
	MessageTo   uint64    `gorm:"column:message_to"`
	Content     string    `gorm:"column:content; NOT NULL; check:length(content)>0"`
	CreatedAt   time.Time `gorm:"column:created_at; NOT NULL"`
	UpdateAt    time.Time `gorm:"column:update_at; NOT NULL"`

	// foreign key
	AccountFrom *Accounts `gorm:"foreignKey:message_from; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AccountTo   *Accounts `gorm:"foreignKey:message_to"`
}

type MessageDataAccessor interface {
	DeleteAll(ctx context.Context) error
	DeleteMessage(ctx context.Context, id uint64) error
	getMessageByID(ctx context.Context, id uint64) (*Messages, error)
	EditMessage(ctx context.Context, message *Messages) (*Messages, error)
	CreateMessage(ctx context.Context, message *Messages) (*Messages, error)
	GetMessages(ctx context.Context, messageFrom, messageTo, offSet, limit uint64) ([]*Messages, error)
}

type messageDataAccessor struct {
	database *gorm.DB
}

func (m messageDataAccessor) CreateMessage(ctx context.Context, message *Messages) (*Messages, error) {
	if message.AccountFrom == nil {
		return nil, errors.New("message from cannot be nil")
	}
	curTime := time.Now()
	message.CreatedAt = curTime
	message.UpdateAt = curTime
	tx := m.database.WithContext(ctx).Create(message)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return message, nil
}

func (m messageDataAccessor) DeleteMessage(ctx context.Context, id uint64) error {
	tx := m.database.WithContext(ctx).Exec("update messages set message_from = null where message_id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m messageDataAccessor) getMessageByID(ctx context.Context, id uint64) (*Messages, error) {
	var res Messages
	tx := m.database.WithContext(ctx).Raw("select * from messages where message_id = ?", id).Scan(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &res, nil
}

func (m messageDataAccessor) EditMessage(ctx context.Context, message *Messages) (*Messages, error) {
	// tx := m.database.Raw(
	// 	"update messages set content = ?, update_at = ? where id = ?",
	// 	message.Content,
	// 	time.Now(),
	// 	message.ID,
	// )
	message.UpdateAt = time.Now()

	tx := m.database.WithContext(ctx).Save(&message)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return message, nil
}

func (m messageDataAccessor) DeleteAll(ctx context.Context) error {
	tx := m.database.WithContext(ctx).Where("1 = 1").Delete(&Messages{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (m messageDataAccessor) GetMessages(ctx context.Context, messageFrom, messageTo, offSet, limit uint64) ([]*Messages, error) {
	var messages []*Messages
	tx := m.database.WithContext(ctx).Raw("select * from messages where message_from = ? and message_to = ? limit ? offset ?", messageFrom, messageTo, limit, offSet).Scan(&messages)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return messages, nil
}

func InitializeMessageDataAccessor(database *gorm.DB) MessageDataAccessor {
	return &messageDataAccessor{
		database: database,
	}
}
