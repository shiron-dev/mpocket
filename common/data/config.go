package data

type Config struct {
	UserData UserData
}

type UserData struct {
	FromGit bool
	Name    string
	Email   string
}

func getDefaultConfig() Config {
	return Config{
		UserData: UserData{
			FromGit: true,
		},
	}
}
