package handlers

import "net/http"

// создание сервиса
func (app *Application) createServiceHandler(w http.ResponseWriter, r *http.Request) {
	name, price, err := app.readServiceParam(r)

	if err != nil {
		app.Logger.Println(err)
		return
	}

	app.Models.Service.Create(name, price)

	if err != nil {
		app.Logger.Println(err)
	}

}

// получение сервиса
func (app *Application) serviceGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)

	if err != nil {
		app.Logger.Println(err)
		return
	}

	service, err := app.Models.Service.Get(id)
	if err != nil {
		app.Logger.Println(err)
		return
	}

	service.ServiceId = id

	err = app.writeJSON(w, http.StatusOK, envelope{"service": service}, nil)
	if err != nil {
		app.Logger.Println(err)
	}
}
