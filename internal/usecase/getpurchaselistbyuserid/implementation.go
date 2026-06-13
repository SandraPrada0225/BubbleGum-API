package getpurchaselistbyuserid

import (
	"bubblegum-api/internal/domain/dto/responses"
	"bubblegum-api/internal/repositories/providers"
)

type Implementation struct {
	VentasProvider providers.VentasProvider
}

func (useCase Implementation) Execute(UserID uint64) (responses.GetPurchaseList, error) {
	return useCase.VentasProvider.GetListByUserID(UserID)
}
