package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLiteドライバ
)

// NewDB は新しいデータベース接続を作成します。
func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
