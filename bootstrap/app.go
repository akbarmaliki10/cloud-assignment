package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env   *Env
	MySql *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = newEnv()
	app.MySql = NewMySqlDatabase(app.Env)

	return *app
}
