package newebpay

import (
	"encoding/json"
	"net/http"
)

type NotifyResultEncrypted struct {
	Status     string `url:"Status"`      // 回傳狀態
	MerchantID string `json:"MerchantID"` // 回傳訊息
	TradeInfo  string `json:"TradeInfo"`  // 交易資料（AES 加密）
	TradeSha   string `json:"TradeSha"`   // 交易資料（SHA256 加密）
	Version    string `json:"Version"`    // 串接程式版本
}

type NotifyResult struct {
	Status  string             `json:"Status"`  // 回傳狀態
	Message string             `json:"Message"` // 回傳訊息
	Result  *PayResponseResult `json:"Result"`  // 回傳內容
}

func (c *Client) ParseNotifyRequest(r *http.Request) (*NotifyResult, error) {
	result := new(NotifyResultEncrypted)
	result.Status = r.PostFormValue("Status")
	result.MerchantID = r.PostFormValue("MerchantID")
	result.TradeInfo = r.PostFormValue("TradeInfo")
	result.TradeSha = r.PostFormValue("TradeSha")
	result.Version = r.PostFormValue("Version")

	notify := new(NotifyResult)
	info := TradeInfoDecrypt(result.TradeInfo, c.key, c.iv)

	err := json.Unmarshal([]byte(info), notify)
	if err != nil {
		return nil, err
	}
	return notify, nil
}
