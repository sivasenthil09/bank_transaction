package controllers

import (
	"context"

	pro "github.com/sivasenthil09/bank_transaction/banktrans"
	"github.com/sivasenthil09/bank_transaction/interfaces"
	"github.com/sivasenthil09/bank_transaction/models"
)

type RPCserver struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCserver) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{
		CustomerID: req.CustomerID}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerID: result.CustomerID,
		}
		return responseCustomer, nil
	}
}
