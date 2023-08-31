package interfaces

import "github.com/sivasenthil09/bank_transaction/models"

type ICustomer interface {
	CreateCustomer(customer *models.Customer) (*models.DBResponse, error)
	Transaction(name string)(*models.DBResponse,error)
}
