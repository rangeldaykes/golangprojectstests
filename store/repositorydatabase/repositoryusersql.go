package repositorydatabase

func GetUserSql() string {

	configIsFirebird := true

	if configIsFirebird {
		return `
		select u.id, u.name
		  from users u
	     where u.id = :id
		`
	} else {
		return ``
	}
}
