package newebpay

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) PostCreditCardClose(ctx context.Context, merchantID string, request *CreditCardCloseRequest) (*CreditCardCloseResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/API/CreditCard/Close", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(CreditCardCloseResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

func (c *Client) PostCreditCardCancel(ctx context.Context, merchantID string, request *CreditCardCancelRequest) (*CreditCardCancelResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/API/CreditCard/Cancel", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(CreditCardCancelResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

func (c *Client) CreditCardCancelCheckCodeHash(merchantID string, request *CreditCardCancelRequest) string {
	hashStr := "HashIV=" + c.iv +
		"&Amt=" + strconv.FormatInt(request.Amt, 10) +
		"&MerchantID=" + merchantID +
		"&MerchantOrderNo=" + request.MerchantOrderNo +
		"&TradeNo=" + request.TradeNo +
		"&HashKey=" + c.key
	h := sha256.New()
	h.Write([]byte(hashStr))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

type CreditCardCloseRequest struct {
	RespondType     string `url:"RespondType"`      // 回傳格式
	Version         string `url:"Version"`          // 串接程式版本
	Amt             int64  `url:"Amt"`              // 取消授權金額
	MerchantOrderNo string `url:"MerchantOrderNo"`  // 商店訂單編號
	TimeStamp       int64  `url:"TimeStamp"`        // 時間戳記
	IndexType       int    `url:"IndexType"`        // 單號類別
	TradeNo         string `url:"TradeNo"`          // 藍新金流交易序號
	CloseType       int    `url:"CloseType"`        // 請款或退款
	Cancel          int    `url:"Cancel,omitempty"` // 取消請款或退款
}

type CreditCardCloseResponse struct {
	Status  string                 `json:"Status"`  // 回傳狀態
	Message string                 `json:"Message"` // 回傳訊息
	Result  *CreditCardCloseResult `json:"Result"`  // 回傳資料
}

type CreditCardCloseResult struct {
	MerchantID      string `json:"MerchantID"`      // 商店代號
	Amt             int64  `json:"Amt"`             // 交易金額
	TradeNo         string `json:"TradeNo"`         // 藍新金流交易序號
	MerchantOrderNo string `json:"MerchantOrderNo"` // 商店訂單編號
}

type CreditCardCancelRequest struct {
	RespondType     string `url:"RespondType"`               // 回傳格式
	Version         string `url:"Version"`                   // 串接程式版本
	Amt             int64  `url:"Amt"`                       // 取消授權金額
	MerchantOrderNo string `url:"MerchantOrderNo,omitempty"` // 商店訂單編號
	TradeNo         string `url:"TradeNo,omitempty"`         // 藍新金流交易序號
	IndexType       int    `url:"IndexType"`                 // 單號類別
	TimeStamp       int64  `url:"TimeStamp"`                 // 時間戳記
}

type CreditCardCancelResponse struct {
	Status  string                    `json:"Status"`  // 回傳狀態
	Message string                    `json:"Message"` // 回傳訊息
	Result  *[]CreditCardCancelResult `json:"Result"`  // 回傳資料
}

type CreditCardCancelResult struct {
	MerchantID      string `json:"MerchantID"`      // 商店代號
	TradeNo         string `json:"TradeNo"`         // 藍新金流交易序號
	Amt             int64  `json:"Amt"`             // 交易金額
	MerchantOrderNo string `json:"MerchantOrderNo"` // 商店訂單編號
	CheckCode       string `json:"CheckCode"`       // 檢核碼
}
