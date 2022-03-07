package model

type Customer struct {
	ID      int64   `json:"id" uri:"id" binding:"required,min=1"`
	Phone   string  `json:"phone"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type GetRequest struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}

type CreateRequest struct {
	Phone   string  `json:"phone"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type UpdateRequest struct {
	Customer
}

type DeleteRequest struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}
