## Requirements
#### 1. Your device must have `make`
#### 2. Your device must have`docker`


## Quick Start:
```
make run-all
```

# LOMS (Logistics and Order Management System)

The service is responsible for order accounting and logistics.

### createOrder

Creates a new order for the user from the list of transferred products.
Goods must be reserved at the warehouse.

Request
```
{
     user int64
     items[]{
         sku uint32
         count uint16
     }
}
```

Response
```
{
     orderID int64
}
```

## listOrder

Shows order information

Request
```
{
     orderID int64
}
```

Response
```
{
     status string // (new | awaiting payment | failed | paid | cancelled)
     user int64
     items[]{
         sku uint32
         count uint16
     }
}
```

### orderPayed

Marks the order as paid. Reserved items should change to purchased status.

Request
```
{
     orderID int64
}
```

Response
```
{}
```

### cancelOrder

Cancels the order, removes the reserve from all products in the order.

Request
```
{
     orderID int64
}
```

Response
```
{}
```

### stocks

Returns the number of items that can be purchased from different warehouses. If the product was reserved from someone in the order and is waiting for payment, it cannot be bought.

Request
```
{
     sku uint32
}
```

Response
```
{
     stocks[]{
         warehouseID int64
         count uint64
     }
}
```

# check out

The service is responsible for the shopping cart and checkout.

### addToCart

Add an item to a specific user's shopping cart. In this case, you need to check the availability of goods through LOMS.stocks

Request
```
{
     user int64
     sku uint32
     count uint16
}
```

Response
```
{}
```

### deleteFromCart

Remove an item from a specific user's shopping cart.

Request
```
{
     user int64
     sku uint32
     count uint16
}
```

Response
```
{}
```

### listCart

Show a list of products in the cart with names and prices (they must be received in real time from the ProductService)

Request
```
{
     user int64
}
```

Response
```
{
     items[]{
         sku uint32
         count uint16
         name string
         price uint32
     }
     totalPrice uint32
}
```

### purchases

Place an order for all items in the cart. Calls createOrder on LOMS.

Request
```
{
     user int64
}
```

Response
```
{
     orderID int64
}
```

# Notifications

Will listen to Kafka and send notifications, there is no external API.

#ProductService

Swagger is deployed at:
http://route256.pavl.uk:8080/docs/

GRPC deployed at:
route256.pavl.uk:8082

## get_product

Request
```
{
     token string
     sku uint32
}
```

Response
```
{
     name string
     price uint32
}
```

## list_skus

Request
```
{
     token string
     startAfterSku uint32
     count uint32
}
```

Response
```
{
     skus[]uint32
}
```


# Way of buying goods

- Checkout.addToCart
     + add to cart and check what is in stock)
- We can remove from the basket
- We can receive a list of goods in the basket
     + title and price are pulled from ProductService.get_product
- We purchase goods through Checkout.purchase
     + go to LOMS.createOrder and create an order
     + The status of the order is new
     + LOMS reserves the required number of items
     + If it was not possible to reserve, the order falls into the failed status
     + If successful, fall into the awaiting payment status
- We pay for the order
     + Call LOMS.orderPayed
     + Reserves are transferred to the write-off of goods from the warehouse
     + The order goes to the paid status
- You can cancel the order before payment
     + Call LOMS.cancelOrder
     + All reservations for the order are canceled, the goods are again available to other users
     + Order goes into canceled status
     + LOMS should itself cancel orders by timeout if not paid within 10 minutes
