package handlers

import (
	"Avito-Tasks/cmd/data"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Получение временного платежа
func (app *Application) getPaymentTempHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.Logger.Println(err)
		return
	}

	payment, err := app.Models.Report.GetTemp(id)
	if err != nil {
		app.Logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"temppayment": payment}, nil)

	if err != nil {
		app.Logger.Println(err)
	}
}

// Совершение успешного платежа
func (app *Application) SuccsesPaymentHandler(w http.ResponseWriter, r *http.Request) {
	TempPaymentID, err := app.readIdParam(r)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}

	TempPayment, err := app.Models.Report.GetTemp(TempPaymentID)
	if err != nil {
		app.Logger.Println(err)
	}

	Service, err := app.Models.Service.Get(TempPayment.ServiceId)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}
	User, err := app.Models.User.Get(TempPayment.UserId)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}
	app.Models.Report.CreatePayment(TempPayment.UserId, TempPayment.ServiceId, TempPaymentID, User.UserReservedCash-Service.ServicePrice)
	if err != nil {
		app.Logger.Println(err)
	}
}

func (app *Application) NotSuccsesPaymentHandler(w http.ResponseWriter, r *http.Request) {
	TempPaymentID, err := app.readIdParam(r)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}

	TempPayment, err := app.Models.Report.GetTemp(TempPaymentID)
	if err != nil {
		app.Logger.Println(err)
	}

	Service, err := app.Models.Service.Get(TempPayment.ServiceId)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}
	User, err := app.Models.User.Get(TempPayment.UserId)
	if err != nil || TempPaymentID < 1 {
		app.Logger.Println(err)
		return
	}
	app.Models.Report.NotCreatePayment(TempPayment.UserId, TempPaymentID, User.UserReservedCash-Service.ServicePrice, User.UserCash+Service.ServicePrice)
	if err != nil {
		app.Logger.Println(err)
	}
}

// функция получения платежа
func (app *Application) getPaymentHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIdParam(r)

	if err != nil || id < 1 {
		app.Logger.Println(err)
		return
	}

	payment, err := app.Models.Report.Get(id)
	if err != nil {
		app.Logger.Println(err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": payment}, nil)

	if err != nil {
		app.Logger.Println(err)
	}
}

// Фунцкция получения данных из таблицы репорт и запись в csv
func (app *Application) getReportHandler(w http.ResponseWriter, r *http.Request) {

	year, err := app.readYearParam(r)
	if err != nil {
		/*app.notFoundResponse(w, r)*/
		log.Println(err)
		return
	}

	month, err := app.readMonthParam(r)
	if err != nil {
		/*app.notFoundResponse(w, r)*/
		log.Println(err)
		return
	}

	reports, err := app.Models.Report.GetMonthlyReport(year, month)

	if err != nil {
		switch {
		/*case errors.Is(err, data.ErrRecordNotFound):
		app.notFoundResponse(w, r)*/
		default:
			log.Println(err)
		}
		return
	}

	//empData := [][]string{
	//	{"Name", "City", "Skills"},
	//	{"Smith", "Newyork", "Java"},
	//	{"William", "Paris", "Golang"},
	//	{"Rose", "London", "PHP"},
	//}
	fileName := fmt.Sprintf("report%d-%d.csv", year, month)

	//filePath := fmt.Sprintf("reports/report%d-%d.csv", year, month)
	//filePath := "./reports/" + fileName
	filePath := "reports/" + fileName
	//fmt.Println(filePath)
	fmt.Println("filePath" + filePath)
	csvFile, err := os.Create(filePath)

	if err != nil {
		//fmt.Println("asf")
		//log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	err = csvwriter.WriteAll(reports)
	if err != nil {
		log.Println("writing to CSV file error")
		http.Error(w, "cannot write to CSV file", http.StatusInternalServerError)
		return
	}
	//var report data.ReportResult

	csvwriter.Flush()

	link := fmt.Sprintf("http://localhost:4000/file/%s", fileName)
	fmt.Println(link)
	//err = app.writeJSON(w, http.StatusOK, envelope{"reports": reports}, nil)
	err = app.writeJSON(w, http.StatusOK, envelope{"link": link}, nil)
	//TODO link

	//app.getFileOfReportHandler(w, r)
	if err != nil {
		//app.logger.Println(err)
		//http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		//app.serverErrorResponse(w, r, err)
		log.Println(err)
	}

}

// получение файлов репорта
func (app *Application) getFileOfReportHandler(w http.ResponseWriter, r *http.Request) {

	fileName, _ := app.readFileNameParam(r)
	//fmt.Println("filename: " + fileName)
	filePath := "reports/" + fileName
	//fmt.Println(filePath)
	http.ServeFile(w, r, filePath)
}

// Получение истории юзера
func (app *Application) getUserReportHistoryHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		data.Filters
	}

	qs := r.URL.Query()
	var err error
	input.Filters.Page = app.readIntFromQuery(qs, "page", 1)
	input.Filters.PageSize = app.readIntFromQuery(qs, "page_size", 5)
	input.Filters.Sort, err = app.readStringFromQuery(qs, "sort", "")
	if err != nil {
		err = app.writeJSON(w, http.StatusBadRequest, envelope{"error": error.Error(err)}, nil)
		return
	}

	id, err := app.readIdParam(r)
	if err != nil || id < 1 {
		log.Println(err)
		return
	}

	userReportHistory, metadata, err := app.Models.Report.GetUserHistory(id, input.Filters)

	if err != nil {
		switch {
		default:
			log.Println(err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK,
		envelope{"metadata": metadata}, nil)
	err = app.writeJSON(w, http.StatusOK,
		envelope{"User Buying Services History": userReportHistory}, nil)
	if err != nil {
		log.Println(err)
	}

}
