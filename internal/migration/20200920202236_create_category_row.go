package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateCategoryRow, downCreateCategoryRow)
}

func upCreateCategoryRow(tx *sql.Tx) error {
	_, err := tx.Exec(`insert into categories (id, created_at, updated_at, deleted_at, name) values ('2e233f0e-82f5-400a-b561-8dc58cfdf164', '2020-09-20 20:26:49.313000', '2020-09-20 20:26:51.637000', null, 'Chichi');`)
	if err != nil {
		return err
	}
	return nil
}

func downCreateCategoryRow(tx *sql.Tx) error {
	_, err := tx.Exec(`delete from categories where id='2e233f0e-82f5-400a-b561-8dc58cfdf164'`)
	if err != nil {
		return err
	}
	return nil
}
