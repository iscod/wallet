package wallet

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const OrderListMinCount = 1
const OrderListMaxCount = 10000

type Wallet struct {
	apiKey string
	Client *http.Client
}

func (w *Wallet) getHeaders() map[string]string {
	return map[string]string{
		"Wpay-Store-Api-Key": w.apiKey,
		"Content-Type":       "application/json",
	}
}

const HOST = "https://pay.wallet.tg"

func NewWallet(apiKey string, timeoutSeconds time.Duration) (*Wallet, error) {
	return &Wallet{apiKey: apiKey, Client: &http.Client{Timeout: timeoutSeconds}}, nil
}

func (w *Wallet) Create(params *CreateParams) (*CreateResponse, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	//https://pay.wallet.tg/wpay/store-api/v1/order
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/wpay/store-api/v1/order", HOST), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range w.getHeaders() {
		req.Header.Set(k, v)
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))

	res := &CreateResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return res, nil
	} else {
		log.Printf("Response error StatusCode %d : %s\n", resp.StatusCode, res.Message)
		return nil, errors.New(res.Status)
	}
}

func (w *Wallet) GetPreview(id string) (*PreviewResponse, error) {
	// https://pay.wallet.tg/wpay/store-api/v1/order/preview&id=id
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wpay/store-api/v1/order/preview", HOST), strings.NewReader(fmt.Sprintf("id=%s", id)))
	if err != nil {
		return nil, err
	}

	for k, v := range w.getHeaders() {
		req.Header.Set(k, v)
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		res := &PreviewResponse{}
		err = json.Unmarshal(body, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return nil, errors.New(fmt.Sprintf("[%v] %s", resp.StatusCode, body))
	}

}

func (w *Wallet) GetGetOrderList(offset, count int) (*OrderListResponse, error) {
	// https://pay.wallet.tg/wpay/store-api/v1/reconciliation/order-list
	if offset < 0 {
		return nil, errors.New("offset must be less than or equal zero")
	}

	if count < OrderListMinCount || count >= OrderListMaxCount {
		return nil, errors.New(fmt.Sprintf("count must be less than zero and greater than %d", OrderListMaxCount))
	}

	//req, _ := NewRequest("POST", "http://www.google.com/", strings.NewReader("z=post"))
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wpay/store-api/v1/reconciliation/order-list", HOST), strings.NewReader(fmt.Sprintf("offset=%d&count=%d", offset, count)))
	if err != nil {
		return nil, err
	}

	for k, v := range w.getHeaders() {
		req.Header.Set(k, v)
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		res := &OrderListResponse{}
		err = json.Unmarshal(body, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		return nil, errors.New(fmt.Sprintf("[%v] %s", resp.StatusCode, body))
	}

}

func (w *Wallet) GetGetOrderAmount(offset, count int) (int, error) {
	// https://pay.wallet.tg/wpay/store-api/v1/reconciliation/order-amount
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/wpay/store-api/v1/reconciliation/order-amount", HOST), nil)
	if err != nil {
		return 0, err
	}

	for k, v := range w.getHeaders() {
		req.Header.Set(k, v)
	}

	resp, err := w.Client.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	res := &OrderAmountResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode == 200 {
		return res.Data.TotalAmount, nil
	} else {
		log.Printf("Response error StatusCode %d : %s\n", resp.StatusCode, res.Message)
		return 0, errors.New(fmt.Sprintf("[%v] %s", resp.StatusCode, body))
	}
}

func (w *Wallet) VerifyingWebhook(uriPath, method string, webhookSign WebhookRequestHeader, webhookBody []WebhookRequestBody) bool {
	str, _ := json.Marshal(webhookBody)
	return w.ComputeSignature(w.apiKey, method, uriPath, webhookSign.Timestamp, string(str)) == webhookSign.Signature
}

func (w *Wallet) ComputeSignature(wpayStoreApiKey string, httpMethod string, uriPath string, timestamp string, body string) string {
	base64Body := base64.StdEncoding.EncodeToString([]byte(body))
	stringToSign := httpMethod + "." + uriPath + "." + timestamp + "." + base64Body
	mac := hmac.New(sha256.New, []byte(wpayStoreApiKey))
	mac.Write([]byte(stringToSign))
	byteArraySignature := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(byteArraySignature)
}
