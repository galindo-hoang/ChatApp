package database

import (
	"embed"
)

var (
	//go:embed migrations/mysql/*
	migrationDirectoryMySQL embed.FS
)

//
//type Migrator interface {
//	Up(context.Context) error
//	Down(context.Context) error
//	// redo() error
//	// status() error
//	// new() error
//}
//
//type migrator struct {
//	db     *sql.DB
//	logger *zap.Logger
//}
//
//func NewMigrator(db *sql.DB /*, logger *zap.Logger*/) Migrator {
//	return &migrator{db: db, logger: nil}
//}
//
//func (m migrator) migrate(ctx context.Context, direction migrate.MigrationDirection) error {
//	if _, err := migrate.ExecContext(ctx, m.db, "mysql", migrate.EmbedFileSystemMigrationSource{
//		FileSystem: migrationDirectoryMySQL,
//		Root:       "migrations/mysql",
//	}, direction); err != nil {
//		fmt.Printf("=============== %s ==================", err.Error())
//		return fmt.Errorf("failted to execute %v migration: %v", direction, err)
//	}
//	return nil
//}
//
//func (m migrator) Up(ctx context.Context) error {
//	return m.migrate(ctx, migrate.Up)
//}
//func (m migrator) Down(ctx context.Context) error {
//	return m.migrate(ctx, migrate.Down)
//}
