package getfiltros

import (
	"bubblegum-api/internal/domain/dto/query"
	"bubblegum-api/internal/domain/entities"
	"bubblegum-api/internal/domain/errors/database"
	"bubblegum-api/internal/repositories/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	useCase                  Implementation
	mockMarcaProvider        *mocks.MockMarcaProvider
	mockCategoriaProvider    *mocks.MockCategoriaProvider
	mockPresentacionProvider *mocks.MockPresentacionProvider
)

func initialize() {
	mockCategoriaProvider = new(mocks.MockCategoriaProvider)
	mockMarcaProvider = new(mocks.MockMarcaProvider)
	mockPresentacionProvider = new(mocks.MockPresentacionProvider)

	useCase = Implementation{
		MarcasProvider:         mockMarcaProvider,
		CategoriasProvider:     mockCategoriaProvider,
		PresentacionesProvider: mockPresentacionProvider,
	}
}

func TestWhenSuccesfullReturnAll(t *testing.T) {
	initialize()
	expectedFiltros := Getfiltros()

	mockMarcaProvider.On("GetAll").Return(expectedFiltros.Marcas, nil)
	mockCategoriaProvider.On("GetAll").Return(expectedFiltros.Categorias, nil)
	mockPresentacionProvider.On("GetAll").Return(expectedFiltros.Presentaciones, nil)

	filtros, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Equal(t, expectedFiltros, filtros)

	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 1)
}

func TestWentWrongGetMarcasRun(t *testing.T) {
	initialize()

	mockMarcaProvider.On("GetAll").Return([]entities.Marca{}, database.NewInterlServerError(""))

	filtros, err := useCase.Execute()
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 0)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 0)
}

func TestWentWrongGetCategoriasRun(t *testing.T) {
	initialize()

	mockMarcaProvider.On("GetAll").Return([]entities.Marca{}, nil)
	mockCategoriaProvider.On("GetAll").Return([]entities.Categoria{}, database.NewInterlServerError(""))

	filtros, err := useCase.Execute()
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 0)
}

func TestWentWrongGetPresentacionesRun(t *testing.T) {
	initialize()

	mockMarcaProvider.On("GetAll").Return([]entities.Marca{}, nil)
	mockCategoriaProvider.On("GetAll").Return([]entities.Categoria{}, nil)
	mockPresentacionProvider.On("GetAll").Return([]entities.Presentacion{}, database.NewInterlServerError(""))

	filtros, err := useCase.Execute()
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 1)
}

func Getfiltros() (filtros query.GetFiltros) {
	filtros = query.GetFiltros{
		Categorias: []entities.Categoria{
			{
				ID:     1,
				Nombre: "Gomitas",
			},
			{
				ID:     2,
				Nombre: "Chupetes",
			},
		},
		Marcas: []entities.Marca{
			{
				ID:     1,
				Nombre: "Trululu",
			},
			{
				ID:     2,
				Nombre: "Jet",
			},
		},
		Presentaciones: []entities.Presentacion{
			{
				ID:     1,
				Nombre: "Caja",
			},
			{
				ID:     2,
				Nombre: "Bolsa",
			},
		},
	}
	return
}
