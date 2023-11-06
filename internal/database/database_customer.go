package database

import (
	"context"
	"errors"

	"github.com/SantiagoBedoya/linkedin-go-ms/internal/dberrors"
	"github.com/SantiagoBedoya/linkedin-go-ms/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllCustomers(ctx context.Context, email string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).Where(models.Customer{Email: email}).Find(&customers)
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) GetCustomerByID(ctx context.Context, ID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).Where(&models.Customer{CustomerID: ID}).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{
				Entity: "customer",
				ID:     ID,
			}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where(&models.Customer{CustomerID: customer.CustomerID}).
		Updates(&models.Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Address:   customer.Address,
		})

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{
			Entity: "customer",
			ID:     customer.CustomerID,
		}
	}
	if err := result.Scan(&customers).Error; err != nil {
		return nil, err
	}
	return &customers[0], nil
}

func (c Client) DeleteCustomer(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Customer{CustomerID: ID}).Error
}
