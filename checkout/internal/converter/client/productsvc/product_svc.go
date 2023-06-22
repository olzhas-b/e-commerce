package productsvc

import (
	"route256/checkout/internal/model"
	"route256/checkout/pkg/productsvc_v1"
)

func GetProductRequest(sku uint32, token string) *productsvc_v1.GetProductRequest {
	return &productsvc_v1.GetProductRequest{
		Token: token,
		Sku:   sku,
	}
}

func ProductFromResp(p *productsvc_v1.GetProductResponse) model.Product {
	return model.Product{
		Name:  p.Name,
		Price: p.Price,
	}
}
