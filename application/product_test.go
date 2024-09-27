package application_test

import (
	"testing"

	"github.com/LuiFeAn/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestApplicationProductEnable(t *testing.T) {

	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())

}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())
}

func TestProductId(t *testing.T) {
	generatedUuid := uuid.NewV4().String()
	product := application.Product{}
	product.ID = generatedUuid
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, generatedUuid, product.ID)
}

func TestProductStatusDisabled(t *testing.T) {
	generatedUuid := uuid.NewV4().String()
	product := application.Product{}
	product.ID = generatedUuid
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, application.DISABLED, product.Status)
}

func TestProductStatusEnabled(t *testing.T) {
	generatedUuid := uuid.NewV4().String()
	product := application.Product{}
	product.ID = generatedUuid
	product.Name = "hello"
	product.Status = application.ENABLED
	product.Price = 10

	require.Equal(t, application.ENABLED, product.Status)
}

func TestProductName(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, "hello", product.Name)
}

func TestProductPrice(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, float64(10), product.Price)
}

func TestProductID(t *testing.T) {

	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = application.DISABLED
	product.Price = 10

	require.Equal(t, "hello", product.Name)
}

func TestProductDisable(t *testing.T) {

	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Equal(t, "The price must be zero in order to have disabled", err.Error())

}
