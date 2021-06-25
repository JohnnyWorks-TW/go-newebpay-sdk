# 藍新金流 SDK for Golang

根據藍新文件所實作的給 Golang 使用的 SDK

- 文件：
  https://www.newebpay.com/website/Page/content/download_api

目前有實作的

- 多功能收款 MPG V1.6

## 使用方式

藍新金流測試平台：https://cwww.newebpay.com/  
依照文件建立好測試站帳號，

得到

- `MerchantID` 商店代號
- `Key` API 串接金鑰
- `IV` API 串接金鑰

> 請參照文件 **藍新金流Newebpay_MPG串接手冊_MPG_1.1.0** 的  
「三、 測試環境串接與作業流程」 與 「四、 正式環境串接與作業流程」 章節

### 轉跳信用卡刷卡範例

以下是轉跳至信用卡刷卡頁面的範例

```go
merchantID := os.Getenv("NEWEBPAY_MERCHANT_ID")
key := os.Getenv("NEWEBPAY_MERCHANT_KEY")
iv := os.Getenv("NEWEBPAY_MERCHANT_IV")

pay, err := newebpay.New(key, iv, newebpay.WithSandbox())
if err != nil {
	panic(err)
	return
}

// TODO 需替換掉
email := "test@example.com"
itemDesc := "This is my test product"
orderNo := "testOrderNo"
amt := 100

serverHost := os.Getenv("SERVER_HOST")

info := &newebpay.TradeInfo{
	MerchantID:      merchantID,
	Amt:             amt,
	Email:           email,
	ItemDesc:        itemDesc,
	LoginType:       0,
	MerchantOrderNo: orderNo,
	RespondType:     "JSON",
	TimeStamp:       time.Now().Unix(),
	Version:         "1.6",
	NotifyURL:       serverHost + "/NotifyURL",
	ReturnURL:       serverHost + "/ReturnURL",
	CREDIT:          1,
}

request, errR := pay.PreparePayRequest(info)
if errR != nil {
	panic(errR)
	return
}

print(request)
```

設定好必填欄位

- `MerchantID` 商店代號
- `RespondType` 回傳格式 (這邊固定帶 JSON)
- `TimeStamp` 時間戳記
- `Version` 串接程式版本 (這邊固定帶 1.6)
- `MerchantOrderNo` 商店自訂訂單編號  
  (限英、數字、 "_" 格式，長度限制為 30 字元，不可重複。)
- `Amt` 訂單金額 (純數字，幣別：新台幣)
- `ItemDesc` 商品資訊 (長度限制為 50 字元，不可帶特殊符號)
- `Email` 付款人電子信箱
- `LoginType` 藍新金流會員 (0 =  不須登入藍新金流會員, 1 = 須要登入藍新金流會員)

> 關於欄位說明細節，請參照文件 **藍新金流Newebpay_MPG串接手冊_MPG_1.1.0** 的  
「五、MPG 參數設定說明」 章節

得到所需 Post 資料

- `MerchantID` 商店代號
- `TradeInfo` 交易資料（AES 加密）
- `TradeSha` 交易資料（SHA256 加密）
- `Version` 串接程式版本

做 Post 轉跳，進入刷卡頁面

藍新將由 `NotifyURL` 所設定的值，背景接收藍新的 Callback  
藍新將由 `ReturnURL` 所設定的值，做 Post 轉跳頁面  

#### 背景接收 Notify 範例

以下是背景接收藍新的 Callback 的範例

```go
func PayNotify(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("NEWEBPAY_MERCHANT_KEY")
	iv := os.Getenv("NEWEBPAY_MERCHANT_IV")

	pay, err := newebpay.New(key, iv, newebpay.WithSandbox())
	if err != nil {
		return
	}
	notify, err := pay.ParseNotifyRequest(r)
	if err != nil {
		return
	}

	// TODO Handle post request
	log.Println(json.Marshal(notify))
}
```

收到 Post 請求之後，呼叫 `ParseNotifyRequest()` 可得到解密後的 json 物件 `NotifyResult`

> 關於欄位說明細節，請參照文件 **藍新金流Newebpay_MPG串接手冊_MPG_1.1.0** 的  
「七、 交易支付系統回傳參數說明」與「八、 取號完成系統回傳參數說明」章節

其餘未寫出的細節，請參考文件。


## License

MIT