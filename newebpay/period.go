package newebpay

import (
	"context"
	"github.com/google/go-querystring/query"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 建立委託單
func (c *Client) PostNewPeriod(ctx context.Context, merchantID string, request *NewPeriodRequest) (*NewPeriodResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/MPG/period", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(NewPeriodResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// 修改已建立委託單狀態
func (c *Client) PostAlterPeriodStatus(ctx context.Context, merchantID string, request *AlterPeriodRequest) (*AlterPeriodResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/MPG/period/AlterStatus", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(AlterPeriodResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

// 修改已建立委託單內容(授權區間及金額)
func (c *Client) PostAlterPeriodAmt(ctx context.Context, merchantID string, request *AlterPeriodAmtRequest) (*AlterPeriodAmtResponse, *http.Response, error) {
	form := url.Values{}

	v, _ := query.Values(request)
	tradeInfoStr := v.Encode()

	form.Add("MerchantID_", merchantID)
	form.Add("PostData_", TradeInfoEncrypt(tradeInfoStr, c.key, c.iv))

	httpReq, err := c.NewRequest(http.MethodPost, "/MPG/period/AlterAmt", strings.NewReader(form.Encode()))
	if err != nil {
		return nil, nil, err
	}
	resp := new(AlterPeriodAmtResponse)
	httpResp, err := c.Do(ctx, httpReq, resp)
	if err != nil {
		return nil, httpResp, err
	}
	return resp, httpResp, nil
}

type NewPeriodRequest struct {
	RespondType     string `url:"RespondType"`               // 回傳格式 (這裡帶JSON)
	TimeStamp       string `url:"TimeStamp"`                 // 時間戳記
	Version         string `url:"Version"`                   // 串接程式版本 (這裡帶1.2)
	LangType        string `url:"LangType,omitempty"`        // 語系
	MerOrderNo      string `url:"MerOrderNo"`                // 商店訂單編號
	ProdDesc        string `url:"ProdDesc"`                  // 產品名稱
	PeriodAmt       int64  `url:"PeriodAmt"`                 // 委託金額
	PeriodType      string `url:"PeriodType"`                // 週期類別 (D=固定天期制 W=每週 M=每月 Y=每年)
	PeriodPoint     string `url:"PeriodPoint"`               // 交易週期授權時間
	PeriodStartType int    `url:"PeriodStartType"`           // 檢查卡號模式 (1=立即執行十元授權 (因部分發卡銀行會阻擋一元交易，因此調整為十元) 2=立即執行委託金額授權 3=不檢查信用卡資訊，不授權)
	PeriodTimes     int    `url:"PeriodTimes"`               // 授權期數
	PeriodFirstdate string `url:"PeriodFirstdate,omitempty"` // 第1期發動日
	ReturnURL       string `url:"ReturnURL,omitempty"`       // 返回商店網址
	PeriodMemo      string `url:"PeriodMemo,omitempty"`      // 備註說明
	PayerEmail      string `url:"PayerEmail"`                // 付款人電子信箱
	EmailModify     int    `url:"EmailModify,omitempty"`     // 付款人電子信箱是否開放修改 (1=可修改	0=不可修改)
	PaymentInfo     string `url:"PaymentInfo,omitempty"`     // 是否開啟付款人資訊 (Y=是 N=否)
	OrderInfo       string `url:"OrderInfo,omitempty"`       // 是否開啟收件人資訊 (Y=是 N=否)
	NotifyURL       string `url:"NotifyURL,omitempty"`       // 每期授權結果通知
	BackURL         string `url:"BackURL,omitempty"`         // 返回商店網址
	UNIONPAY        int    `url:"UNIONPAY,omitempty"`        // 銀聯卡啟用 (1=啟用 0 或者未有此參數=不啟用)
}

type NewPeriodResponse struct {
	Status  string           `json:"Status,omitempty"`  // 回傳狀態
	Message string           `json:"Message,omitempty"` // 回傳訊息
	Result  *NewPeriodResult `json:"Result,omitempty"`  // 回傳資料
}

type NewPeriodResult struct {
	MerchantID      string `json:"MerchantID"`              // 商店代號
	MerchantOrderNo string `json:"MerchantOrderNo"`         // 商店訂單編號
	PeriodType      string `json:"PeriodType"`              // 週期類別
	AuthTimes       int64  `json:"AuthTimes"`               // 授權次數
	AuthTime        string `json:"AuthTime,omitempty"`      // 授權時間
	DateArray       string `json:"DateArray"`               // 授權排程日期
	TradeNo         string `json:"TradeNo,omitempty"`       // 藍新金流交易序號
	CardNo          string `json:"CardNo,omitempty"`        // 卡號前六後四碼
	PeriodAmt       int64  `json:"PeriodAmt"`               // 每期金額
	AuthCode        string `json:"AuthCode,omitempty"`      // 授權碼
	RespondCode     string `json:"RespondCode,omitempty"`   // 銀行回應碼
	EscrowBank      string `json:"EscrowBank,omitempty"`    // 款項保管銀行
	AuthBank        string `json:"AuthBank,omitempty"`      // 收單機構
	PaymentMethod   string `json:"PaymentMethod,omitempty"` // 交易類別
	PeriodNo        string `json:"PeriodNo"`                // 委託單號
	Extday          string `json:"Extday"`                  // 信用卡到期日
}

type PeriodNotify struct {
	RespondCode     string    `json:"RespondCode"`             // 銀行回應碼
	MerchantID      string    `json:"MerchantID"`              // 商店代號
	MerchantOrderNo string    `json:"MerchantOrderNo"`         // 商店訂單編號
	OrderNo         string    `json:"OrderNo"`                 // 自訂單號
	TradeNo         string    `json:"TradeNo"`                 // 交易序號
	AuthDate        time.Time `json:"AuthDate"`                // 授權時間
	TotalTimes      string    `json:"TotalTimes"`              // 總期數
	AlreadyTimes    string    `json:"AlreadyTimes"`            // 已授權次數
	AuthAmt         int64     `json:"AuthAmt"`                 // 授權金額
	AuthCode        string    `json:"AuthCode"`                // 授權碼
	EscrowBank      string    `json:"EscrowBank"`              // 款項保管銀行
	AuthBank        string    `json:"AuthBank"`                // 收單機構
	PaymentMethod   string    `json:"PaymentMethod,omitempty"` // 交易類別
	NextAuthDate    time.Time `json:"NextAuthDate"`            // 下次授權日期
	PeriodNo        string    `json:"PeriodNo"`                // 委託單號
}

type AlterPeriodRequest struct {
	RespondType string `url:"RespondType"` // 回傳格式
	Version     string `url:"Version"`     // 串接程式版本
	MerOrderNo  string `url:"MerOrderNo"`  // 商店訂單編號
	PeriodNo    string `url:"PeriodNo"`    // 委託單號
	AlterType   string `url:"AlterType"`   // 委託狀態
	TimeStamp   string `url:"TimeStamp"`   // 時間戳記
}

type AlterPeriodResponse struct {
	Status  string             `json:"Status"`  // 回傳狀態
	Message string             `json:"Message"` // 回傳訊息
	Result  *AlterPeriodResult `json:"Result"`  // 回傳資料
}

type AlterPeriodResult struct {
	MerOrderNo  string `json:"MerOrderNo"`  // 商店訂單編號
	PeriodNo    string `json:"PeriodNo"`    // 委託單號
	AlterType   string `json:"AlterType"`   // 委託狀態
	NewNextTime string `json:"NewNextTime"` // 下一期授權日期
}

type AlterPeriodAmtRequest struct {
	RespondType string `url:"RespondType"`           // 回傳格式
	Version     string `url:"Version"`               // 串接程式版本
	TimeStamp   string `url:"TimeStamp"`             // 時間戳記
	MerOrderNo  string `url:"MerOrderNo"`            // 商店訂單編號
	PeriodNo    string `url:"PeriodNo"`              // 委託單號
	AlterAmt    int64  `url:"AlterAmt,omitempty"`    // 委託金額
	PeriodType  string `url:"PeriodType,omitempty"`  // 週期類別
	PeriodPoint string `url:"PeriodPoint,omitempty"` // 交易週期授權時間
	PeriodTimes string `url:"PeriodTimes,omitempty"` // 授權期數
	Extday      string `url:"Extday"`                // 信用卡到期日
}

type AlterPeriodAmtResponse struct {
	Status  string                `json:"Status"`  // 回傳狀態
	Message string                `json:"Message"` // 回傳訊息
	Result  *AlterPeriodAmtResult `json:"Result"`  // 回傳資料

}

type AlterPeriodAmtResult struct {
	MerOrderNo  string `json:"MerOrderNo"`  // 商店訂單編號
	PeriodNo    string `json:"PeriodNo"`    // 委託單號
	AlterAmt    string `json:"AlterAmt"`    // 委託金額
	PeriodType  string `json:"PeriodType"`  // 週期類別
	PeriodPoint string `json:"PeriodPoint"` // 交易週期授權時間
	NewNextAmt  string `json:"NewNextAmt"`  // 下一期授權金額
	NewNextTime string `json:"NewNextTime"` // 下一期授權時間
	PeriodTimes string `json:"PeriodTimes"` // 授權期數
}
