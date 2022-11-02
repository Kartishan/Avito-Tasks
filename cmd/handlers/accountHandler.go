package handlers

import (
	"fmt"
	"net/http"
)

func (app *Application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	app.Models.User.Create()
}

func (app *Application) getUserHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		/*app.notFoundResponse(w, r)*/
		app.Logger.Println(err)
		return
	}

	user, err := app.Models.User.Get(id)
	if err != nil {
		switch {
		/*case errors.Is(err, data.ErrRecordNotFound):
		app.notFoundResponse(w, r)*/
		default:
			//app.serverErrorResponse(w, r, err)
			app.Logger.Println(err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)

	if err != nil {
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.Logger.Println(err)
	}
}

// функция перевода денег из кошелька в резерв
func (app *Application) depositReservUserHandler(w http.ResponseWriter, r *http.Request) {
	id_user, id_service, err := app.readDepositReservParam(r)

	if err != nil {
		app.Logger.Println(err)
		//app.notFoundResponse(w, r)
		return
	}

	user, err := app.Models.User.Get(id_user)
	service, err := app.Models.Service.Get(id_service)
	var price = service.ServicePrice
	if err != nil {
		app.Logger.Println(err)

		return
	}
	if user.UserId == 0 || service.ServiceId == 0 {
		app.Logger.Println("Service or user not found")
		return
	}
	user.UserId = id_user
	if user.UserCash-price < 0 {
		app.Logger.Println("not enough money")
		return
	}
	user.UserCash -= price
	user.UserReservedCash += price

	err = app.Models.User.UpdateFull(user)
	if err != nil {
		app.Logger.Println(err)
	}

	print("перед createTemp")
	app.Models.Report.CreateTemp(id_user, id_service)
	if err != nil {
		app.Logger.Println("Can not create temp_payment")
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.Logger.Println(err)
	}
}

// функция добавления денег на аккаунт
func (app *Application) adddepositUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	cash, err := app.readCashParam(r)

	if err != nil {
		app.Logger.Println(err)
		return
	}

	user, err := app.Models.User.Get(id)
	if err != nil || user == nil {
		app.Models.User.CreateId(id)
		app.Logger.Println(err)
	}
	user, err = app.Models.User.Get(id)

	user.UserId = id
	user.UserCash += cash

	err = app.Models.User.Update(user)
	if err != nil {
		app.Logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.Logger.Println(err)
	}
	app.Models.Transaction.Create(user.UserId, user.UserId, 1, cash)
}

// функция перевода денег с аккаунта на аккаунт
func (app *Application) transferUserHandler(w http.ResponseWriter, r *http.Request) {
	FromId, err := app.readIdParam(r)
	ToId, err := app.readToIdParam(r)
	cash, err := app.readCashParam(r)
	if cash < 0 {
		return
	}

	userFrom, err := app.Models.User.Get(FromId)
	if err != nil {
		switch {
		default:
			app.Logger.Println(err)
		}
		return
	}

	userTo, err := app.Models.User.Get(ToId)
	if err != nil {
		switch {
		default:
			app.Logger.Println(err)
		}
		return
	}
	var newCash float64 = userFrom.UserCash - cash
	if newCash < 0 {
		return
	} else {
		fmt.Printf("%2f", newCash)
		userFrom.UserCash = newCash
		newCash = userTo.UserCash + cash
		userTo.UserCash = newCash
		app.Models.User.Update(userFrom)
		app.Models.User.Update(userTo)
		app.Models.Transaction.Create(userFrom.UserId, userTo.UserId, 3, cash)
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"userFrom": userFrom}, nil)
	err = app.writeJSON(w, http.StatusOK, envelope{"userTo": userTo}, nil)
}

func (app *Application) withdrawalepositUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	cash, err := app.readCashParam(r)

	if err != nil {
		app.Logger.Println(err)
		return
	}

	user, err := app.Models.User.Get(id)
	if err != nil {
		switch {
		default:
			app.Logger.Println(err)
		}
		return
	}
	if user.UserId == 0 {
		app.Logger.Println("user not found")
		return
	}
	user.UserId = id
	user.UserCash -= cash

	err = app.Models.User.Update(user)
	if err != nil {
		app.Logger.Println(err)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.Logger.Println(err)
	}
	app.Models.Transaction.Create(user.UserId, user.UserId, 2, cash)
}
