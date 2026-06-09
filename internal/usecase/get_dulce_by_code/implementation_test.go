package getdulcebycode

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/repositories/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	useCase               Implementation
	mockDulceProvider     *mocks.MockDulceProvider
	mockCategoriaProvider *mocks.MockCategoriaProvider
)

const (
	errInternalServer = "database.InternalServerError"
	errNotFound       = "database.NotFoundError"
)

func initialize() {
	mockDulceProvider = new(mocks.MockDulceProvider)
	mockCategoriaProvider = new(mocks.MockCategoriaProvider)
	useCase = Implementation{
		DulcesProvider:     mockDulceProvider,
		CategoriasProvider: mockCategoriaProvider,
	}
}

func TestWhenSuccesfullReturnDulce(t *testing.T) {
	initialize()

	partialResponse := getPartialResponse()
	mockCategorias := getMockCategorias()

	expectedResponse := partialResponse
	expectedResponse.Categorias = mockCategorias

	mockDulceProvider.On("GetByCode", partialResponse.Codigo).Return(partialResponse, nil)
	mockCategoriaProvider.On("GetCategoriasByDulceID", partialResponse.ID).Return(mockCategorias, nil)

	dulce, err := useCase.Execute(partialResponse.Codigo)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, dulce)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 1)
}

func TestWhenGetDulceWhentWrongShouldReturnInternalServerError(t *testing.T) {
	initialize()
	expectedDulce := getPartialResponse()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(responses.DetalleDulce{}, database.NewInterlServerError("error"))

	dulce, err := useCase.Execute(expectedDulce.Codigo)

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, dulce)
	assert.Equal(t, errInternalServer, errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhendulceBotFoundThenShouldReturnNotFoundError(t *testing.T) {
	initialize()
	expectedDulce := getPartialResponse()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(responses.DetalleDulce{}, database.NewNotFoundError("error"))

	dulce, err := useCase.Execute(expectedDulce.Codigo)

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, dulce)
	assert.Equal(t, errNotFound, errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenGetCategoriasWentWrongShouldReturnInternalServer(t *testing.T) {
	initialize()
	partialResponse := getPartialResponse()
	mockCategorias := getMockCategorias()

	expectedResponse := partialResponse
	expectedResponse.Categorias = mockCategorias

	mockDulceProvider.On("GetByCode", partialResponse.Codigo).Return(partialResponse, nil)
	mockCategoriaProvider.On("GetCategoriasByDulceID", partialResponse.ID).Return([]entities.Categoria{}, database.NewInterlServerError("error"))

	response, err := useCase.Execute(partialResponse.Codigo)
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, response)
	assert.Equal(t, errInternalServer, errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 1)
}

func getPartialResponse() responses.DetalleDulce {
	return responses.DetalleDulce{
		ID:           2,
		Nombre:       "Chocolatina",
		Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Codigo:       "2",
		Presentacion: entities.Presentacion{
			ID:     1,
			Nombre: "Empaque",
		},
		Marca: entities.Marca{
			ID:     2,
			Nombre: "Jet",
		},
	}
}

func getMockCategorias() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     1,
			Nombre: "Gomitas",
		},
		{
			ID:     2,
			Nombre: "Chocolates",
		},
	}
	return
}
