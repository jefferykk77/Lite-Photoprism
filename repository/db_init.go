package repository

func InitSchema() string {
	schema := `CREATE TABLE IF NOT EXISTS Photos (
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	PATH TEXT NOT NULL,
	CONTENT TEXT NOT NULL) `
	return schema
}

type Photo struct {
	ID      int    `db:"ID"`
	Path    string `db:"PATH"`
	Content string `db:"CONTENT"`
}
