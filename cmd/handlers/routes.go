package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *Application) Routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/user/getId/:id", app.getUserHandler)

	router.HandlerFunc(http.MethodPut, "/user/add/:id/:user_cash", app.adddepositUserHandler)

	router.HandlerFunc(http.MethodPut, "/user/transfer/:id/:ToId/:user_cash", app.transferUserHandler)

	router.HandlerFunc(http.MethodPut, "/user/withdrawal/:id/:user_cash", app.withdrawalepositUserHandler)

	router.HandlerFunc(http.MethodPut, "/user/reserv/:id_user/:id_service", app.depositReservUserHandler)

	router.HandlerFunc(http.MethodGet, "/service/get/:id", app.serviceGetHandler)

	router.HandlerFunc(http.MethodPost, "/service/create/:name/:price", app.createServiceHandler)

	router.HandlerFunc(http.MethodGet, "/reporttemp/get/:id", app.getPaymentTempHandler)
	//передается ид из темп репорт
	router.HandlerFunc(http.MethodPut, "/report/create/:id", app.SuccsesPaymentHandler)
	//передается ид из темп репорт
	router.HandlerFunc(http.MethodPut, "/report/notcreate/:id", app.NotSuccsesPaymentHandler)

	router.HandlerFunc(http.MethodGet, "/report/getID/:id", app.getPaymentHandler)

	router.HandlerFunc(http.MethodGet, "/report/get/:year/:month", app.getReportHandler)

	router.HandlerFunc(http.MethodGet, "/transaction/get/:id", app.getTransactionHandler)

	router.HandlerFunc(http.MethodGet, "/file/:filename", app.getFileOfReportHandler)

	router.HandlerFunc(http.MethodGet, "/history/transactions/:id", app.getTransactionHistoryHandler)

	router.HandlerFunc(http.MethodGet, "/history/user/:id", app.getUserReportHistoryHandler)

	return router
}
