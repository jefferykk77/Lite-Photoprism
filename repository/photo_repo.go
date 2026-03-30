package repository

import (
	"github.com/jmoiron/sqlx"
)

type RepoManager struct {
	db *sqlx.DB //为什么要加一层sqlx，而不直接用sqlite
}

func NewRepoManager(db *sqlx.DB) *RepoManager {
	return &RepoManager{db: db}
}

func (RM *RepoManager) Create(path string, content string) error {
	_, err := RM.db.Exec("INSERT INTO Photos (PATH,CONTENT) VALUES(?,?)", path, content)
	return err
}

func (RM *RepoManager) ReadByID(id int) (Photo, error) {
	var photo Photo
	err := RM.db.Get(&photo, "SELECT * FROM Photos WHERE ID =?", id)
	return photo, err
}
