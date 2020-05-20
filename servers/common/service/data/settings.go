package data

import "github.com/cjburchell/settings-go"

const defaultMongoUrl = "localhost"

type MongoSettings struct {
	Address string
	UserName string
	Password string
}

type Settings struct {
	Mongo MongoSettings
}

func GetSettings(settings settings.ISettings) Settings  {
	return Settings{
		Mongo:MongoSettings{
			Address:  settings.GetSection("Mongo").Get("Address", defaultMongoUrl),
			UserName: settings.GetSection("Mongo").Get("UserName", ""),
			Password: settings.GetSection("Mongo").Get("Password", ""),
		},
	}
}