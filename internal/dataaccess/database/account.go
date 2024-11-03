package database

import (
	"context"

	"gorm.io/gorm"
)

type Accounts struct {
	Id          uint64 `gorm:"column:id; primaryKey; index; AUTO_INCREMENT;"`
	AccountName string `gorm:"column:account_name; NOT NULL; check:length(account_name)>0"`
	Email       string `gorm:"column:email; unique"`
	Password    string `gorm:"column:password; NOT NULL; check:length(password)>0"`
}

type AccountDataAccessor interface {
	GetAccountByID(ctx context.Context, id uint64) (*Accounts, error)
	CreateAccount(ctx context.Context, account *Accounts) (*Accounts, error)
	GetAccountByEmail(ctx context.Context, email string) (*Accounts, error)
	DeleteAll(ctx context.Context) error
}

type accountDataAccessor struct {
	database *gorm.DB
}

func (a accountDataAccessor) GetAccountByID(ctx context.Context, id uint64) (*Accounts, error) {
	var result *Accounts = nil
	tx := a.database.WithContext(ctx).Raw("select * from accounts where id = ?", id).Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}

func (a accountDataAccessor) CreateAccount(ctx context.Context, account *Accounts) (*Accounts, error) {
	res := a.database.WithContext(ctx).Create(&account)
	if res.Error != nil {
		return nil, res.Error
	}
	return account, nil
}

func (a accountDataAccessor) GetAccountByEmail(ctx context.Context, email string) (*Accounts, error) {
	var result *Accounts = nil
	tx := a.database.WithContext(ctx).Raw("select * from accounts where email = ?", email).Scan(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}

func (a accountDataAccessor) DeleteAll(ctx context.Context) error {
	// difference between truncate and delete
	tx := a.database.WithContext(ctx).Where("1 = 1").Delete(&Accounts{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func InitializeAccountDataAccessor(db *gorm.DB) AccountDataAccessor {
	return &accountDataAccessor{
		database: db,
	}
}
