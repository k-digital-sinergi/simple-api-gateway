package model

type Customer struct {
	ID      int64   `json:"id" uri:"id" binding:"required,min=1"`
	Phone   string  `json:"phone"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type ListResponse struct {
	Data []Customer `json:"data"`
}

type GetRequest struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}

type GetResponse struct {
	Data *Customer `json:"data"`
}

type CreateRequest struct {
	Phone   string  `json:"phone"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type CreateResponse struct {
	ResponseMessage
	CustID int64 `json:"cust_id"`
}

type UpdateRequest struct {
	Customer
}

type UpdateResponse struct {
	ResponseMessage
}

type DeleteRequest struct {
	CustID int64 `json:"cust_id" uri:"id" binding:"required,min=1"`
}

type DeleteResponse struct {
	ResponseMessage
}
