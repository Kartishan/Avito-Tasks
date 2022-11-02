package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

func (r ReportModel) GetMonthlyReport(year int64, month int64) ([][]string, error) {
	if year < 0 || month > 12 || month < 1 {
		return nil, errors.New("incorrect data")
	}

	query := `
			SELECT
			service.service_id,
			SUM(service.service_price)
			FROM report, service
			WHERE (EXTRACT(YEAR FROM report_time) = $1
			AND EXTRACT(MONTH FROM report_time) = $2) AND (service.service_id = report.service_id)
			GROUP BY service.service_id;
		`
	//args := []interface{}{year, month}
	rows, err := r.DB.Query(query, year, month)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var strings = make([][]string, 0)
	strings = append(strings,
		[]string{"service_id", fmt.Sprintf("Revenue for:  year: %d month: %d", year, month)})

	var reports []ReportResult

	for rows.Next() {
		var report ReportResult
		var str = make([]string, 0)
		err := rows.Scan(&report.ServiceId, &report.Revenue)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
		str = append(str, fmt.Sprintf(`%d`, report.ServiceId), report.Revenue)
		strings = append(strings, str)
	}

	if err != nil {
		switch {
		/*case errors.Is(err, sql.ErrNoRows):
		return nil, ErrRecordNotFound*/
		default:
			errors.New("smh")
			return nil, err
		}
	}

	return strings, nil
}

func (r ReportModel) GetUserHistory(id int64, filters Filters) ([]UserReportHistoryResponse, Metadata, error) {
	if id < 1 {
		return nil, Metadata{}, errors.New("incorrect id")
	}

	query := fmt.Sprintf(`
			SELECT COUNT(*) OVER(), user_id,
			report_time, service.service_price, service.service_name
			FROM report, service
			WHERE user_id = $1 AND report.service_id = service.service_id
			ORDER BY %s %s, user_id ASC
			LIMIT $2 OFFSET $3
		`, filters.sortColumnReportQuery(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query, id, filters.limit(), filters.offset())
	if err != nil {
		return nil, Metadata{}, errors.New("No records")
	}

	defer rows.Close()

	totalRecords := 0

	var userReportHistory []UserReportHistoryResponse

	for rows.Next() {
		var userHistory UserReportHistoryResponse
		//var str = make([]string, 0)
		err := rows.Scan(&totalRecords, &userHistory.UserId, &userHistory.ReportTime,
			&userHistory.ServicePrice, &userHistory.ServiceName)
		if err != nil {
			return nil, Metadata{}, err
		}
		userReportHistory = append(userReportHistory, userHistory)
	}

	if err != nil {
		switch {
		/*case errors.Is(err, sql.ErrNoRows):
		return nil, ErrRecordNotFound*/
		default:
			errors.New("something went wrong")
			return nil, Metadata{}, err
		}
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return userReportHistory, metadata, nil
}

func (r ReportModel) CreateTemp(user_id, service_id int64) {
	query := `
		INSERT INTO temp_report (user_id, service_id)
		VALUES ($1, $2)`

	r.DB.QueryRow(query, user_id, service_id).Scan()
}

func (r ReportModel) GetTemp(id int64) (*Report, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}
	fmt.Println("GetTemp")
	query := `
			SELECT *
			FROM temp_report
			WHERE temp_report_id = $1
		`

	var report Report
	err := r.DB.QueryRow(query, id).Scan(
		&report.ReportId,
		&report.UserId,
		&report.ServiceId,
		&report.ReportTime,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("getTempError")
		default:
			return nil, err
		}
	}

	return &report, nil
}

func (r ReportModel) Get(id int64) (*Report, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}
	fmt.Println("GetTemp")
	query := `
			SELECT *
			FROM report
			WHERE report_id = $1
		`

	var report Report
	err := r.DB.QueryRow(query, id).Scan(
		&report.ReportId,
		&report.UserId,
		&report.ServiceId,
		&report.ReportTime,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("")
		default:
			return nil, err
		}
	}

	return &report, nil
}

func (r ReportModel) CreatePayment(userId, serviceId, tempReporttId int64, NEW float64) {

	query := `
		INSERT INTO report (user_id, service_id)
		VALUES ($1, $2);`
	r.DB.QueryRow(query, userId, serviceId).Scan()
	print("quari1")

	query2 := `
		UPDATE usert
			set user_reserved_cash = $2
			WHERE user_id = $1;`
	r.DB.QueryRow(query2, userId, NEW).Scan()
	print("quari2")

	query3 := `
		DELETE FROM temp_report
		    WHERE temp_report_id = $1;`
	r.DB.QueryRow(query3, tempReporttId).Scan()
	return
}
func (r ReportModel) NotCreatePayment(userId, tempReporttId int64, NEW, cash float64) {
	query2 := `
		UPDATE usert
			set (user_reserved_cash,user_cash) = ($2, $3)
			WHERE user_id = $1;`
	r.DB.QueryRow(query2, userId, NEW, cash).Scan()

	query3 := `
		DELETE FROM temp_report
		    WHERE temp_report_id = $1;`
	r.DB.QueryRow(query3, tempReporttId).Scan()
	return
}
