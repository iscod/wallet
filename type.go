package wallet

import "time"

type ResponseStatus string

type ResponseStatusCode string

const (
	ResponseStatusCodeSuccess        ResponseStatusCode = "SUCCESS"
	ResponseStatusCodeALREADY        ResponseStatusCode = "ALREADY"
	ResponseStatusCodeConflict       ResponseStatusCode = "CONFLICT"
	ResponseStatusCodeAccessDenied   ResponseStatusCode = "ACCESS_DENIED"
	ResponseStatusCodeInvalidRequest ResponseStatusCode = "ACCESS_DENIED"
	ResponseStatusCodeInternalError  ResponseStatusCode = "INTERNAL_ERROR"
)

type Amount struct {
	Amount       string `json:"amount"`       // 1.000
	CurrencyCode string `json:"currencyCode"` //USDT,TON,BTC
}

type SelectedPaymentOption struct {
	Amount       *Amount `json:"amount"`
	AmountFee    *Amount `json:"amountFee"`
	AmountNet    *Amount `json:"amountNet"`
	ExchangeRate string  `json:"exchangeRate"`
}

type Order struct {
	Id                     string                 `json:"id"`     //2703383946854401
	Status                 string                 `json:"status"` // "ACTIVE" "EXPIRED" "PAID" "CANCELLED"
	Number                 string                 `json:"number"` // 9aeb581c
	Amount                 *Amount                `json:"amount"`
	AutoConversionCurrency string                 `json:"autoConversionCurrency"` // "USDT"
	ExternalId             string                 `json:"externalId"`             //"ORD-5023-4E89"
	CustomerTelegramUserId int                    `json:"customerTelegramUserId"` //"0"
	CreatedDateTime        time.Time              `json:"createdDateTime"`
	ExpirationDateTime     time.Time              `json:"expirationDateTime"`
	CompletedDateTime      time.Time              `json:"completedDateTime"`
	PaymentDateTime        time.Time              `json:"paymentDateTime"`
	PayLink                string                 `json:"payLink"`
	DirectPayLink          string                 `json:"directPayLink"`
	SelectedPaymentOption  *SelectedPaymentOption `json:"selectedPaymentOption"`
}

type Items struct {
	Items []Order `json:"items"`
}

type CreateParams struct {
	Amount                 Amount `json:"amount"`
	AutoConversionCurrency string `json:"autoConversionCurrency,omitempty"` //"USDT"
	Description            string `json:"description"`                      //"VPN for 1 month"
	ReturnUrl              string `json:"returnUrl"`                        //"https://t.me/wallet/start?startapp"
	FailReturnUrl          string `json:"failReturnUrl"`                    //"https://t.me/wallet"
	CustomData             string `json:"customData,omitempty"`             //"client_ref=4E89"
	ExternalId             string `json:"externalId"`                       //"ORD-5023-4E89"
	TimeoutSeconds         int    `json:"timeoutSeconds"`                   //"10800"
	CustomerTelegramUserId int    `json:"customerTelegramUserId"`           //"0"
}

type CreateResponse struct {
	Status  string `json:"status"` //"SUCCESS" "ALREADY" "CONFLICT" "ACCESS_DENIED" "INVALID_REQUEST" "INTERNAL_ERROR"
	Message string `json:"message"`
	Data    *Order `json:"data"`
}

//type PreviewParams struct {
//	Id string `json:"id"`
//}

type PreviewResponse struct {
	Status  string `json:"status"` //"SUCCESS" "INVALID_REQUEST" "INTERNAL_ERROR"
	Message string `json:"message"`
	Data    *Order `json:"data"`
}

//type OrderListParams struct {
//	Offset int64 `json:"offset"`
//	Count  int32 `json:"count"`
//}

type OrderListResponse struct {
	Status  string `json:"status"` //"SUCCESS" "ALREADY" "CONFLICT" "ACCESS_DENIED" "INVALID_REQUEST" "INTERNAL_ERROR"
	Message string `json:"message"`
	Data    *Items `json:"data"`
}

type TotalAmount struct {
	TotalAmount int `json:"totalAmount"`
}

type OrderAmountResponse struct {
	Status  string      `json:"status"` //"SUCCESS" "ALREADY" "CONFLICT" "ACCESS_DENIED" "INVALID_REQUEST" "INTERNAL_ERROR"
	Message string      `json:"message"`
	Data    TotalAmount `json:"data"`
}

type WebhookRequestHeader struct {
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

type WebhookRequestBody struct {
	EventDateTime string                    `json:"eventDateTime"`
	EventId       int                       `json:"eventId"`
	Type          string                    `json:"type"`
	Payload       WebhookRequestBodyPayLoad `json:"payload"`
}

type WebhookRequestBodyPayLoad struct {
	Id                     int64                 `json:"id"`
	Number                 string                `json:"number"`
	Status                 string                `json:"status,omitempty"`
	CustomData             string                `json:"customData,omitempty"`
	ExternalId             string                `json:"externalId"`
	OrderAmount            Amount                `json:"orderAmount"`
	SelectedPaymentOption  SelectedPaymentOption `json:"selectedPaymentOption,omitempty"`
	OrderCompletedDateTime string                `json:"orderCompletedDateTime,omitempty"`
}
