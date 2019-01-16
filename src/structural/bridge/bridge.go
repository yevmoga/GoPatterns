package bridge


type formatter interface {
	ToFormat() string
}

type Apper interface {
	GetData() string
}

type App struct {
	Format formatter
}

func AppConstructor(format formatter) Apper {
	return &App{
		Format: format,
	}
}

func (app *App) GetData() string {
	return app.Format.ToFormat()
}
