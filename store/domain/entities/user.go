package entities

type User struct {
	ID   int    `db:"ID"`
	Name string `db:"NAME"`
}
