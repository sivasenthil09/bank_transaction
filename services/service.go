package services

import (
	"context"

	"github.com/sivasenthil09/bank_transaction/interfaces"
	"github.com/sivasenthil09/bank_transaction/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{
		CustomerCollection: collection,
		ctx: ctx,
	}
}

func (p *CustomerService) CreateCustomer(customer *models.Customer) (*models.DBResponse, error) {
	// Insert the customer into the MongoDB collection
	_, err := p.CustomerCollection.InsertOne(p.ctx, customer)
	if err != nil {
		return nil, err
	}

	// Construct the response with customer ID and CreatedAt
	response := &models.DBResponse{
		CustomerID: customer.CustomerID,
		CreatedAt:  customer.CreatedAt,
	}


	return response, nil
}
func(p*CustomerService)Transaction(name string)(*models.DBResponse,error){
	

}

