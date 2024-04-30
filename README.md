# wallet

> Package for using Wallet Pay API https://docs.wallet.tg/pay/

[![Build status](https://ci.appveyor.com/api/projects/status/v5lt859vmjm3v9i5?svg=true)](https://ci.appveyor.com/project/iscod/wallet)


## Supported endpoints

Create order [/wpay/store-api/v1/order][createOrderEndpoint]

Get order preview [/wpay/store-api/v1/order/preview][getPreview]

Get order list [/wpay/store-api/v1/reconciliation/order-list][getOrderList]

Get order amount [/wpay/store-api/v1/reconciliation/order-amount][getOrderAmount]

## Installation

```sh
go get github.com/iscod/wallet
```

## Usage examples

### init wallet

```golang
w, err := NewWallet("secret_api_key", time.Duration(time.Second))

if err != nil {
	fmt.printf("Error creating wallet: %s\n", err)
}
```

### create new order
[/wpay/store-api/v1/order][createOrderEndpoint]

```go
w, err := NewWallet("secret_api_key", time.Duration(time.Second))

if err != nil {
    fmt.printf("Error creating wallet: %s\n", err)
}

createResponse, err := w.Create(&wallet.CreateParams{
    Description:            "VPN for 1 month",
    Amount:                 wallet.Amount{Amount: "1.00", CurrencyCode: "TON"},
    ReturnUrl:              "https://t.me/wallet/start?startapp",
    TimeoutSeconds:         10800,
    CustomerTelegramUserId: 0,
})

var order *wallet.Order = createResponse.Data

fmt.Println(order.Id)
```

### get order preview
[/wpay/store-api/v1/order/preview][getPreview]

```go
w, err := NewWallet("secret_api_key", time.Duration(time.Second))

if err != nil {
    fmt.printf("Error creating wallet: %s\n", err)
}

previewResponse, err := w.GetPreview("2703383946854401")

var order *wallet.Order = previewResponse.Data

fmt.Println(order.Id)
```
