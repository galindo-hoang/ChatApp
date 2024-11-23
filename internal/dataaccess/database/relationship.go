package database

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type RelationshipDataAccessor interface {
	CreateNode(ctx context.Context, node *Accounts) error
	CreateRelationship(ctx context.Context, from uint64, to uint64) error
	RemoveRelationship(ctx context.Context, from uint64, to uint64) error
}

var databaseName = "neo4j"

type relationshipDataAccessor struct {
	driver neo4j.DriverWithContext
}

func InitializeRelationshipDataAccessor(driver neo4j.DriverWithContext) RelationshipDataAccessor {
	return &relationshipDataAccessor{driver: driver}
}

func (a relationshipDataAccessor) CreateNode(ctx context.Context, node *Accounts) error {
	session := a.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: databaseName})
	defer session.Close(ctx)
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		r, err := tx.Run(ctx,
			"CREATE (node:Person {id: $id, name: $name, email: $email})",
			map[string]any{
				"id":    node.Id,
				"name":  node.AccountName,
				"email": node.Email,
			},
		)
		return r, err
	})

	return err
}

func (a relationshipDataAccessor) CreateRelationship(ctx context.Context, from uint64, to uint64) error {
	session := a.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: databaseName})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		r, err := tx.Run(ctx,
			"MATCH (from:Person {id: $from}), (to:Person {id: $to}) MERGE (from)-[:Following]->(to)",
			map[string]any{
				"from": from,
				"to":   to,
			},
		)
		return r, err
	})

	return err
}

func (a relationshipDataAccessor) RemoveRelationship(ctx context.Context, from uint64, to uint64) error {
	session := a.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: databaseName})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		r, err := tx.Run(ctx,
			"MATCH (from:Person {id: $from})-[f:Following]->(to:Person {id: $to}) DELETE f",
			map[string]any{
				"from": from,
				"to":   to,
			},
		)
		return r, err
	})

	return err
}
