package internal

import (
	"log"
	"readinglistapp/initialisers"
	"readinglistapp/model"
	"readinglistapp/view"
)

type IApp interface {
	view.IView
	model.IModel
	initialisers.IDB
	GetView() *view.View
	GetModel() *model.Model
	GetDB() *initialisers.DB
}

type App struct {
	View  *view.View
	Model *model.Model
	DB    *initialisers.DB
}

func (a App) GetView() *view.View {
	return a.View
}

func (a App) GetModel() *model.Model {
	return a.Model
}

func (a App) GetDB() *initialisers.DB {
	return a.DB
}

func (a App) NewView() *view.View {
	if a.View == nil {
		a.View = view.NewView()
	}

	return a.View
}

func (a App) NewModel() *model.Model {
	if a.Model == nil {
		a.Model = model.NewModel()
	}

	return a.Model
}

func (a App) NewDB() *initialisers.DB {
	if a.DB == nil {
		var err error
		a.DB, err = initialisers.NewDB()
		if err != nil {
			log.Fatal("err")
		}
	}
	return a.DB
}
