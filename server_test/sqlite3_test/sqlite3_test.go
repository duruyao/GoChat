package sqlite3_test_test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"testing"
)

const adminsCreateSQL = `DROP TABLE IF EXISTS ADMINS_TB;
CREATE TABLE ADMINS_TB
(
    UID  VARCHAR(32) NOT NULL UNIQUE,
    PWD  TEXT        NOT NULL,
    TIME TEXT        DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO ADMINS_TB (UID, PWD)
VALUES ('root', '971213'),
       ('admin', '971114');`

const adminsQuerySQL = `SELECT UID, PWD, TIME FROM ADMINS_TB`

func TestSQLite3(t *testing.T) {
	path := "/home/dry/.GoChat/db/gochat.db"
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		t.Fatal(err)
	}
	file, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	_ = file.Close()

	var db *sql.DB
	db, _ = sql.Open("sqlite3", path)
	defer func() { _ = db.Close() }()

	if _, err = db.Exec(adminsCreateSQL); err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query(adminsQuerySQL)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var admin struct {
			Uid  string
			Pwd  string
			Time string
		}
		if err := rows.Scan(&admin.Uid, &admin.Pwd, &admin.Time); err != nil {
			t.Fatal(err)
		}
		t.Log(admin)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}
}
