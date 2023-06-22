package server

import (
	"route256/loms/internal/model"
	"route256/loms/pkg/loms_v1"
)

func OrderFromReq(in *loms_v1.CreateOrderRequest) model.Order {
	var items []model.Item
	for _, item := range in.Order.Items {
		items = append(items, itemFromReq(item))
	}
	return model.Order{
		User:  in.Order.UserId,
		Items: items,
	}
}

func OrderToResp(order model.Order) *loms_v1.Order {
	return &loms_v1.Order{
		UserId: order.User,
		Status: order.Status,
		Items:  ItemsToReq(order.Items),
	}
}

func OrderToListResp(order model.Order) *loms_v1.ListOrderResponse {
	var response loms_v1.ListOrderResponse
	response.Order = OrderToResp(order)
	return &response
}

func OrderIDToResp(orderID int64) *loms_v1.CreateOrderResponse {
	return &loms_v1.CreateOrderResponse{
		OrderId: orderID,
	}
}
