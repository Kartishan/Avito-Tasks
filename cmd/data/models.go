package data

import "database/sql"

type Transaction struct {
	TransactionId   int64   `json:"transaction_id"`
	SenderId        int64   `json:"sender_id"`
	ReceiverId      int64   `json:"receiver_id"`
	OperationId     int64   `json:"operation_id"`
	TransactionTime string  `json:"transaction_time"`
	Price           float64 `json:"price"`
}
type TransactionModel struct {
	DB *sql.DB
}
type TransactionHistory struct {
	TransactionId    int64   `json:"transaction_id"`
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type TransactionHistoryModel struct {
	DB *sql.DB
}

type TransactionHistoryResponse struct {
	SenderId         int64   `json:"sender_id"`
	ReceiverId       int64   `json:"receiver_id"`
	TransactionTime  string  `json:"transaction_time"`
	TransactionPrice float64 `json:"transaction_price"`
	OperationType    string  `json:"operation_type"`
}

type User struct {
	UserId           int64   `json:"user_id"`
	UserCash         float64 `json:"user_cash"`
	UserReservedCash float64 `json:"user_reserved_cash"`
}

type UserModel struct {
	DB *sql.DB
}
type Service struct {
	ServiceId    int64   `json:"service_id"`
	ServiceName  string  `json:"service_name"`
	ServicePrice float64 `json:"service_price"`
}
type ServiceModel struct {
	DB *sql.DB
}
type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

type Report struct {
	UserId     int64  `json:"user_id"`
	ReportId   int64  `json:"report_id"`
	ServiceId  int64  `json:"service_id"`
	ReportTime string `json:"report_time"`
}

type ReportModel struct {
	DB *sql.DB
}

type ReportResult struct {
	ServiceId int64  `json:"service_id"`
	Revenue   string `json:"report_revenue"`
}

type UserReportHistoryResponse struct {
	UserId       int64   `json:"user_id"`
	ReportTime   string  `json:"report_time"`
	ServicePrice float64 `json:"service_price"`
	ServiceName  string  `json:"service_name"`
}

type Models struct {
	User               UserModel
	Report             ReportModel
	TransactionHistory TransactionHistoryModel
	Transaction        TransactionModel
	Service            ServiceModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		User:               UserModel{DB: db},
		Report:             ReportModel{DB: db},
		TransactionHistory: TransactionHistoryModel{DB: db},
		Transaction:        TransactionModel{DB: db},
		Service:            ServiceModel{DB: db},
	}
}
