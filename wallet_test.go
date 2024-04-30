package wallet

import (
	"testing"
	"time"
)

func TestSign(t *testing.T) {

	w, _ := NewWallet("your_secret_api_key_sYIpNypce5sls6Ik", time.Duration(time.Second))
	param := WebhookRequestHeader{
		Signature: "MGfJzeEprADZbihhRcGcCY5pYTI/IEJ91ejyA+XOWAs=",
		Timestamp: "168824905680291",
	}
	//`[{"eventDateTime":"2023-07-28T10:20:17.681338Z","eventId":10030477545046017,"type":"ORDER_PAID","payload":{"id":10030467668508673,"number":"XYTNJP2O","customData":"in exercitation culpa","externalId":"JDF23NN","orderAmount":{"amount":"0.100000340","currencyCode":"TON"},"selectedPaymentOption":{"amount":{"amount":"0.132653","currencyCode":"USDT"},"amountFee":{"amount":"0.001327","currencyCode":"USDT"},"amountNet":{"amount":"0.131326","currencyCode":"USDT"},"exchangeRate":"1.3265247467314987"},"orderCompletedDateTime":"2023-07-28T10:20:17.628946Z"}}]`
	body := make([]WebhookRequestBody, 0)
	body = append(body, WebhookRequestBody{
		EventDateTime: "2023-07-28T10:20:17.681338Z",
		EventId:       10030477545046017,
		Type:          "ORDER_PAID",
		Payload: WebhookRequestBodyPayLoad{
			Id:         10030467668508673,
			Number:     "XYTNJP2O",
			CustomData: "in exercitation culpa",
			ExternalId: "JDF23NN",
			OrderAmount: Amount{
				Amount:       "0.100000340",
				CurrencyCode: "TON",
			},
			SelectedPaymentOption: SelectedPaymentOption{
				Amount: &Amount{
					Amount:       "0.132653",
					CurrencyCode: "USDT",
				},
				AmountFee: &Amount{
					Amount:       "0.001327",
					CurrencyCode: "USDT",
				},
				AmountNet: &Amount{
					Amount:       "0.131326",
					CurrencyCode: "USDT",
				},
				ExchangeRate: "1.3265247467314987",
			},
			OrderCompletedDateTime: "2023-07-28T10:20:17.628946Z",
		},
	})

	if !w.VerifyingWebhook("/webhook/", "POST", param, body) {
		t.Errorf("Test Verifying Webhook not pass")
	}
}
