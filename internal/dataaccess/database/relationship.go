package database

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type RelationshipDataAccessor interface {
	// CreateNode(ctx context.Context, node *database.Accounts) error
	CreateRelationship(context.Context, uint64, uint64) error
	RemoveRelationship(context.Context, uint64, uint64) error
	GetListRequests(context.Context, uint64) ([]*uint64, error)
	GetListPendingRequesters(context.Context, uint64) ([]*uint64, error)
	UpdateRelationship(context.Context, uint64, uint64) error
	GetListFriends(context.Context, uint64) ([]*uint64, error)
}

type mysqlDataAccessor struct {
	databaseName string
	driver       *gorm.DB
}

func InitializeMysqlDataAccessor(driver *gorm.DB) RelationshipDataAccessor {
	return &mysqlDataAccessor{
		databaseName: "mysql",
		driver:       driver,
	}
}

type UserFriend struct {
	id     uint64 `gorm:"primaryKey; index; AUTO_INCREMENT"`
	uid1   uint64
	uid2   uint64
	status string
}

func (r mysqlDataAccessor) CreateRelationship(ctx context.Context, from uint64, to uint64) error {
	fmt.Println("Create relationship...")
	var parsedData = make(map[string]interface{})
	if from < to {
		parsedData["uid1"] = from
		parsedData["uid2"] = to
		parsedData["status"] = "REQ_UID1"
	} else {
		parsedData["uid1"] = to
		parsedData["uid2"] = from
		parsedData["status"] = "REQ_UID2"
	}
	tx := r.driver.WithContext(ctx).Raw("insert into user_friend (uid1, uid2, status) values (?, ?, ?)", parsedData["uid1"], parsedData["uid1"], parsedData["status"])
	return tx.Error
}

func (r mysqlDataAccessor) RemoveRelationship(ctx context.Context, from uint64, to uint64) error {
	fmt.Println("Removing relationship...")

	var parsedData = make(map[string]interface{})
	if from < to {
		parsedData["uid1"] = from
		parsedData["uid2"] = to
	} else {
		parsedData["uid1"] = to
		parsedData["uid2"] = from
	}

	tx := r.driver.WithContext(ctx).Where(&UserFriend{
		uid1: parsedData["uid1"].(uint64),
		uid2: parsedData["uid2"].(uint64),
	})
	return tx.Error
}

func (r mysqlDataAccessor) GetListRequests(ctx context.Context, id uint64) ([]*uint64, error) {
	fmt.Println("Retrieving requests...")
	var results []*UserFriend
	tx := r.driver.WithContext(ctx).Raw("select * from user_friend where (uid1 = ? status = 'REQ_UID1') OR (uid2 = ? status = 'REQ_UID2')", id, id).Scan(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var ids []*uint64

	for _, result := range results {
		if result.uid1 == id {
			ids = append(ids, &result.uid2)
		} else {
			ids = append(ids, &result.uid1)
		}
	}
	return ids, nil
}

func (r mysqlDataAccessor) GetListPendingRequesters(ctx context.Context, id uint64) ([]*uint64, error) {
	fmt.Println("Retreving pending requests...")

	var results []*UserFriend
	tx := r.driver.WithContext(ctx).Raw("select * from user_friend where (uid1 = ? status = 'REQ_UID2') OR (uid2 = ? status = 'REQ_UID1')", id, id).Scan(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}

	var ids []*uint64

	for _, result := range results {
		if result.uid1 == id {
			ids = append(ids, &result.uid2)
		} else {
			ids = append(ids, &result.uid1)
		}
	}
	return ids, nil
}

func (r mysqlDataAccessor) UpdateRelationship(ctx context.Context, from uint64, to uint64) error {
	fmt.Println("Update requests...")

	var parsedData = make(map[string]interface{})
	if from < to {
		parsedData["uid1"] = from
		parsedData["uid2"] = to
	} else {
		parsedData["uid1"] = to
		parsedData["uid2"] = from
	}

	tx := r.driver.WithContext(ctx).Raw("update user_friend set status = 'FRIEND' where uid1 = ? and uid2 = ?", parsedData["uid1"], parsedData["uid2"])

	return tx.Error
}

func (r mysqlDataAccessor) GetListFriends(ctx context.Context, from uint64) ([]*uint64, error) {
	fmt.Println("Retrieving friends...")

	var results []*UserFriend
	tx := r.driver.WithContext(ctx).Raw("select * from user_friend where (uid1 = ? or uid2 = ?) AND status = 'FRIENDS'", from, from).Scan(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var ids []*uint64

	for _, result := range results {
		if result.uid1 == from {
			ids = append(ids, &result.uid2)
		} else {
			ids = append(ids, &result.uid1)
		}
	}
	return ids, nil
}
