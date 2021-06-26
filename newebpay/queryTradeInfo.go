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
	"time"
)

func (c *Client) PostQueryTradeInfo(ctx context.Context, merchantID string, request *QueryTradeInfoRequest) (*QueryTradeInfoResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID", request.MerchantID)
	form.Add("Version", request.Version)
	form.Add("RespondType", request.RespondType)
	form.Add("CheckValue", c.QueryTradeInfoCheckValueHash(request))
	form.Add("TimeStamp", request.TimeStamp)
	form.Add("MerchantOrderNo", request.MerchantOrderNo)
	form.Add("Amt", strconv.FormatInt(request.Amt, 10))
	form.Add("Gateway", request.Gateway)

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/API/QueryTradeInfo", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(QueryTradeInfoResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

func (c *Client) QueryTradeInfoCheckValueHash(request *QueryTradeInfoRequest) string {
	hashStr := "HashIV=" + c.iv +
		"&Amt=" + strconv.FormatInt(request.Amt, 10) +
		"&MerchantID=" + request.MerchantID +
		"&MerchantOrderNo=" + request.MerchantOrderNo +
		"&HashKey=" + c.key
	h := sha256.New()
	h.Write([]byte(hashStr))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

func (c *Client) QueryTradeInfoCheckCodeHash(result *QueryTradeInfoResult) string {
	hashStr := "HashIV=" + c.iv +
		"&Amt=" + strconv.FormatInt(result.Amt, 10) +
		"&MerchantID=" + result.MerchantID +
		"&MerchantOrderNo=" + result.MerchantOrderNo +
		"&TradeNo=" + result.TradeNo +
		"&HashKey=" + c.key
	h := sha256.New()
	h.Write([]byte(hashStr))
	return strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
}

type QueryTradeInfoRequest struct {
	MerchantID      string `url:"MerchantID"`        // 商店代號
	Version         string `url:"Version"`           // 串接程式版本
	RespondType     string `url:"RespondType"`       // 回傳格式
	CheckValue      string `url:"CheckValue"`        // 檢查碼
	TimeStamp       string `url:"TimeStamp"`         // 時間戳記
	MerchantOrderNo string `url:"MerchantOrderNo"`   // 商店訂單編號
	Amt             int64  `url:"Amt"`               // 訂單金額
	Gateway         string `url:"Gateway,omitempty"` // 資料來源
}

type QueryTradeInfoResponse struct {
	Status  string                `json:"Status"`  // 回傳狀態
	Message string                `json:"Message"` // 回傳訊息
	Result  *QueryTradeInfoResult `json:"Result"`  // 回傳資料
}

type QueryTradeInfoResult struct {
	MerchantID      string    `json:"MerchantID"`      // 商店代號
	Amt             int64     `json:"Amt"`             // 交易金額
	TradeNo         string    `json:"TradeNo"`         // 藍新金流交易序號
	MerchantOrderNo string    `json:"MerchantOrderNo"` // 商店訂單編號
	TradeStatus     int       `json:"TradeStatus"`     // 支付狀態
	PaymentType     string    `json:"PaymentType"`     // 支付方式
	CreateTime      time.Time `json:"CreateTime"`      // 交易建立時間
	PayTime         time.Time `json:"PayTime"`         // 支付完成時間
	CheckCode       string    `json:"CheckCode"`       // 檢核碼
	FundTime        time.Time `json:"FundTime"`        // 預計撥款日
	ShopMerchantID  string    `json:"ShopMerchantID"`  // 實際交易商店代號
	// 信用卡專屬欄位
	RespondCode   string `json:"RespondCode,omitempty"`   // 金融機構回應碼
	Auth          string `json:"Auth,omitempty"`          // 授權碼
	ECI           string `json:"ECI,omitempty"`           // ECI
	CloseAmt      int64  `json:"CloseAmt,omitempty"`      // 請款金額
	CloseStatus   int    `json:"CloseStatus,omitempty"`   // 請款狀態
	BackBalance   int64  `json:"BackBalance,omitempty"`   // 可退款餘額
	BackStatus    int    `json:"BackStatus,omitempty"`    // 退款狀態
	RespondMsg    string `json:"RespondMsg,omitempty"`    // 授權結果訊息
	Inst          int    `json:"Inst,omitempty"`          // 分期-期別
	InstFirst     int64  `json:"InstFirst,omitempty"`     // 分期-首期金額
	InstEach      int64  `json:"InstEach,omitempty"`      // 分期-每期金額
	PaymentMethod string `json:"PaymentMethod,omitempty"` // 交易類別
	Card6No       string `json:"Card6No,omitempty"`       // 信用卡前 6 碼CheckValue
	Card4No       string `json:"Card4No,omitempty"`       // 信用卡後 4 碼
	AuthBank      string `json:"AuthBank,omitempty"`      // 收單金融機構
	// 超商代碼、超商條碼、ATM 轉帳專屬欄位
	PayInfo    string    `json:"PayInfo,omitempty"`    // 付款資訊
	ExpireDate time.Time `json:"ExpireDate,omitempty"` // 繳費有效期限
}
