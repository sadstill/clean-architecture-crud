package storage

type Author struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
