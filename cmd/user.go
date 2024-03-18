package cmd

type User struct {
	name string
}

func (u User) createUser(username string) {
	u.name = username
}
