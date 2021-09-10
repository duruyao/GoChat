package data

import (
	"fmt"
	"github.com/duruyao/gochat/server/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

var db *sqlx.DB

// DbPath returns '$HOME/.GoChat/data/gochat.sqlite'.
func DbPath() string {
	return fmt.Sprintf("%s/.GoChat/data/gochat.sqlite", util.UserHomeDir())
}

// DbDir returns '$HOME/.GoChat/data'.
func DbDir() string {
	return filepath.Dir(DbPath())
}

// IsNotExist returns true if the file '$HOME/.GoChat/data/gochat.sqlite' doesn't exists, otherwise false.
func IsNotExist() bool {
	if _, err := os.Stat(DbPath()); os.IsNotExist(err) {
		return true
	}
	return false
}

// createDb creates a new path '$HOME/.GoChat/data/gochat.sqlite'.
func createDb() (err error) {
	var file *os.File
	if _, e := os.Stat(DbDir()); os.IsNotExist(e) {
		if err = os.MkdirAll(DbDir(), os.ModePerm); err != nil {
			return err
		}
	}
	file, err = os.Create(DbPath())
	if err != nil {
		return
	}
	if err = file.Close(); err != nil {
		return
	}
	if err = openDb(); err != nil {
		return
	}
	err = createTables()
	return
}

//
func openDb() (err error) {
	db, err = sqlx.Open("sqlite3", DbPath())
	return
}

//
func closeDb() error {
	return db.Close()
}

func createTables() (err error) {
	cmd := `DROP TABLE IF EXISTS USERS;

CREATE TABLE USERS
(
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    UUID       VARCHAR(63) NOT NULL UNIQUE,
    NAME       VARCHAR(63) NOT NULL UNIQUE,
    PASSWORD   TEXT,
    MAX_ROLE   INTEGER     NOT NULL DEFAULT 0,
    CREATED_AT TIMESTAMP   NOT NULL DEFAULT (DATETIME(CURRENT_TIMESTAMP, 'LOCALTIME'))
);

DROP TABLE IF EXISTS SESSIONS;

CREATE TABLE SESSIONS
(
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    UUID       VARCHAR(63) NOT NULL UNIQUE,
    USER_ID    INTEGER     NOT NULL,
    CREATED_AT TIMESTAMP   NOT NULL DEFAULT (DATETIME(CURRENT_TIMESTAMP, 'LOCALTIME')),
    FOREIGN KEY (USER_ID) REFERENCES USERS (ID)
);

DROP TABLE IF EXISTS GROUPS;

CREATE TABLE GROUPS
(
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    UUID       VARCHAR(63) NOT NULL UNIQUE,
    NAME       VARCHAR(63) NOT NULL UNIQUE,
    ADMIN_ID   INTEGER     NOT NULL,
    TOKEN      VARCHAR(63),
    CREATED_AT TIMESTAMP   NOT NULL DEFAULT (DATETIME(CURRENT_TIMESTAMP, 'LOCALTIME')),
    FOREIGN KEY (ADMIN_ID) REFERENCES USERS (ID)
);

DROP TABLE IF EXISTS MEMBERS;

CREATE TABLE MEMBERS
(
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    UUID       VARCHAR(63) NOT NULL UNIQUE,
    GROUP_ID   INTEGER     NOT NULL,
    USER_ID    INTEGER     NOT NULL,
    CREATED_AT TIMESTAMP   NOT NULL DEFAULT (DATETIME(CURRENT_TIMESTAMP, 'LOCALTIME')),
    FOREIGN KEY (GROUP_ID) REFERENCES GROUPS (ID),
    FOREIGN KEY (USER_ID) REFERENCES USERS (ID)
);`
	_, err = db.Exec(cmd)
	return
}
