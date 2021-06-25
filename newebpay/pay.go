package newebpay

import (
	"github.com/google/go-querystring/query"
)

func (c *Client) PreparePayRequest(info *TradeInfo) (*PayRequest, error) {
	v, err := query.Values(info)
	if err != nil {
		return nil, err
	}
	tradeInfoStr := v.Encode()

	resp := &PayRequest{
		Host:       c.endpoint.String() + "MPG/mpg_gateway",
		MerchantID: info.MerchantID,
		TradeInfo:  TradeInfoEncrypt(tradeInfoStr, c.key, c.iv),
		TradeSha:   TradeInfoHash(tradeInfoStr, c.key, c.iv),
		Version:    info.Version,
	}
	return resp, nil
}

type PayRequest struct {
	Host       string `json:"Host"`
	MerchantID string `json:"MerchantID"`
	TradeInfo  string `json:"TradeInfo"`
	TradeSha   string `json:"TradeSha"`
	Version    string `json:"Version"`
}

type TradeInfo struct {
	MerchantID      string `url:"MerchantID"`              // 商店代號
	RespondType     string `url:"RespondType"`             // 回傳格式 (請帶 JSON)
	TimeStamp       int64  `url:"TimeStamp"`               // 時間戳記
	Version         string `url:"Version"`                 // 串接程式版本 (請帶 1.6)
	LangType        string `url:"LangType,omitempty"`      // 繁體中文版參數為 zh-tw
	MerchantOrderNo string `url:"MerchantOrderNo"`         // 商店訂單編號 限英、數字、"_"格 式，30字元，同一商店中此編號不可重覆
	Amt             int64  `url:"Amt"`                     // 訂單金額 新台幣純數字
	ItemDesc        string `url:"ItemDesc"`                // 商品資訊 50 字
	TradeLimit      int64  `url:"TradeLimit,omitempty"`    // 限制交易的秒數
	ExpireDate      string `url:"ExpireDate,omitempty"`    // 繳費有效期限(適用於非即時交易)
	ReturnURL       string `url:"ReturnURL,omitempty"`     // 支付完成 返回商店網址
	NotifyURL       string `url:"NotifyURL,omitempty"`     // 支付通知網址
	CustomerURL     string `url:"CustomerURL,omitempty"`   // 商店取號網址
	ClientBackURL   string `url:"ClientBackURL,omitempty"` // 返回商店網址
	Email           string `url:"Email"`                   // 付款人電子信箱
	EmailModify     int    `url:"EmailModify,omitempty"`   // 付款人電子信箱 是否開放修改 0 or 1
	LoginType       int    `url:"LoginType"`               // 藍新金流會員 0 = 不須登入 1 = 須要登入藍新金流會員
	OrderComment    string `url:"OrderComment,omitempty"`  // 商店備註
	CREDIT          int    `url:"CREDIT,omitempty"`        // 信用卡 一次付清啟用
	ANDROIDPAY      int    `url:"ANDROIDPAY,omitempty"`    // Google Pay 啟用
	SAMSUNGPAY      int    `url:"SAMSUNGPAY,omitempty"`    // Samsung Pay 啟用
	LINEPAY         int    `url:"LINEPAY,omitempty"`       // LINE Pay
	ImageUrl        string `url:"ImageUrl,omitempty"`      // LINE PAY 產品圖檔連結 網址
	InstFlag        int    `url:"InstFlag,omitempty"`      // 信用卡 分期付款啟用
	CreditRed       int    `url:"CreditRed,omitempty"`     // 信用卡 紅利啟用
	CREDITAE        int    `url:"CREDITAE,omitempty"`      // 信用卡 美國運通卡啟用
	UNIONPAY        int    `url:"UNIONPAY,omitempty"`      // 信用卡 銀聯卡啟用
	WEBATM          int    `url:"WEBATM,omitempty"`        // WebATM 啟用
	VACC            int    `url:"VACC,omitempty"`          // ATM 轉帳啟用
	CVS             int    `url:"CVS,omitempty"`           // 超商代碼繳費 啟用
	BARCODE         int    `url:"BARCODE,omitempty"`       // 超商條碼繳費啟用
	ALIPAY          int    `url:"ALIPAY,omitempty"`        // 支付寶啟用
	P2G             int    `url:"P2G,omitempty"`           // ezPay 電子錢包
	CVSCOM          int    `url:"CVSCOM,omitempty"`        // 物流啟用 	1 = 啟用超商取貨不付款 2 = 啟用超商取貨付款 3 = 啟用超商取貨不付款及超商取貨付款 0 或者未有此參數，即代表不開啟。
}

const (
	PaymentMethodCredit  = "CREDIT"  // 信用卡
	PaymentMethodWebATM  = "WEBATM"  // WebATM
	PaymentMethodVacc    = "VACC"    // ATM 轉帳 (非即時交易)
	PaymentMethodCvs     = "CVS"     // 超商代碼繳費 (非即時交易)
	PaymentMethodBarcode = "BARCODE" // 超商條碼繳費 (非即時交易)
	PaymentMethodCvscom  = "CVSCOM"  // 超商取貨付款 (非即時交易)
	PaymentMethodAlipay  = "ALIPAY"  // 支付寶 (非即時交易)
	PaymentMethodP2Geacc = "P2GEACC" // ezPay 電子錢包
	PaymentMethodLinePay = "LINEPAY" // LINE Pay
)

const (
	BankNameHNCB      = "HNCB"      // 華南銀行
	BankNameEsun      = "Esun"      // 玉山銀行
	BankNameTaishin   = "Taishin"   // 台新銀行
	BankNameCTBC      = "CTBC"      // 中國信託銀行
	BankNameNCCC      = "NCCC"      // 聯合信用卡中心
	BankNameCathayBK  = "CathayBK"  // 國泰世華銀行
	BankNameCitibank  = "Citibank"  // 花旗銀行
	BankNameUBOT      = "UBOT"      // 聯邦銀行
	BankNameSKBank    = "SKBank"    // 新光銀行
	BankNameFubon     = "Fubon"     // 富邦銀行
	BankNameFirstBank = "FirstBank" // 第一銀行
)

// PayResponse type
type PayResponse struct {
	Status    string `json:"Status"`
	Message   string `json:"Message"`
	ResultRaw string `json:"Result"`
	result    *PayResponseResult
}

type PayResponseResult struct {
	// 所有支付方式共同回傳參數
	MerchantID      string      `json:"MerchantID,omitempty"`      // 商店代號
	Amt             int64       `json:"Amt,omitempty"`             // 交易金額 新台幣純數字
	TradeNo         string      `json:"TradeNo,omitempty"`         // 藍新金流交易序號
	MerchantOrderNo string      `json:"MerchantOrderNo,omitempty"` // 商店訂單編號
	PaymentType     string      `json:"PaymentType,omitempty"`     // 支付方式
	RespondType     string      `json:"RespondType,omitempty"`     // 回傳格式 JSON
	PayTime         JsonPayTime `json:"PayTime,omitempty"`         // 支付完成時間
	IP              string      `json:"IP,omitempty"`              // IP
	EscrowBank      string      `json:"EscrowBank,omitempty"`      // 款項保管銀行
	// 信用卡支付回傳參數(一次付清、Google Pay、Samaung Pay、AE、國民旅遊卡、銀聯)
	AuthBank          string    `json:"AuthBank,omitempty"`          // 收單金融機構 [Esun]: 玉山銀行 [Taishin]: 台新銀行 [CTBC]: 中國信託銀行 [NCCC]: 聯合信用卡中心 [CathayBK]: 國泰世華銀行 [Citibank]:花旗銀行 [UBOT]:聯邦銀行 [SKBank]:新光銀行 [Fubon]:富邦銀行 [FirstBank]:第一銀行
	RespondCode       string    `json:"RespondCode,omitempty"`       // 金融機構回應碼
	Auth              string    `json:"Auth,omitempty"`              // 授權碼
	Card6No           string    `json:"Card6No,omitempty"`           // 卡號前六碼
	Card4No           string    `json:"Card4No,omitempty"`           // 卡號末四碼
	Inst              int64     `json:"Inst,omitempty"`              // 分期-期別
	InstFirst         int64     `json:"InstFirst,omitempty"`         // 分期-首期金額
	InstEach          int64     `json:"InstEach,omitempty"`          // 分期-每期金額
	ECI               string    `json:"ECI,omitempty"`               // ECI 值
	TokenUseStatus    StringInt `json:"TokenUseStatus,omitempty"`    // 信用卡快速結帳使用狀態
	RedAmt            int64     `json:"RedAmt,omitempty"`            // 紅利折抵後實際金額
	PaymentMethod     string    `json:"PaymentMethod,omitempty"`     // 交易類別
	DCC_Amt           float64   `json:"DCC_Amt,omitempty"`           //外幣金額
	DCC_Rate          float64   `json:"DCC_Rate,omitempty"`          // 匯率
	DCC_Markup        float64   `json:"DCC_Markup,omitempty"`        // 風險匯率
	DCC_Currency      string    `json:"DCC_Currency,omitempty"`      // 幣別
	DCC_Currency_Code int64     `json:"DCC_Currency_Code,omitempty"` // 幣別代碼
	// WEBATM、ATM 繳費回傳參數
	PayBankCode       string `json:"PayBankCode,omitempty"`       // 付款人金融機構代碼
	PayerAccount5Code string `json:"PayerAccount5Code,omitempty"` // 付款人金融機構帳號末五碼
	// 超商代碼繳費回傳參數
	CodeNo  string `json:"CodeNo,omitempty"`  // 繳費代碼
	StoreID string `json:"StoreID,omitempty"` // 繳費門市代號
	// 超商條碼繳費回傳參數
	Barcode_1 string `json:"Barcode_1,omitempty"` // 第一段條碼
	Barcode_2 string `json:"Barcode_2,omitempty"` // 第二段條碼
	Barcode_3 string `json:"Barcode_3,omitempty"` // 第三段條碼
	PayStore  string `json:"PayStore,omitempty"`  // 繳費超商
	// ezPay 電子錢包回傳參數
	P2GTradeNo     string `json:"P2GTradeNo,omitempty"`     // ezPay 交易序號
	P2GPaymentType string `json:"P2GPaymentType,omitempty"` // ezPay 支付方式
	P2GAmt         int64  `json:"P2GAmt,omitempty"`         // ezPay 交易金額
	// 超商物流回傳參數
	StoreCode   string `json:"StoreCode,omitempty"`   // 超商門市編號
	StoreName   string `json:"StoreName,omitempty"`   // 超商門市名稱
	StoreType   int64  `json:"StoreType,omitempty"`   // 超商類別名稱
	StoreAddr   string `json:"StoreAddr,omitempty"`   // 超商門市地址
	TradeType   int    `json:"TradeType,omitempty"`   // 取件交易方式
	CVSCOMName  string `json:"CVSCOMName,omitempty"`  // 取貨人
	CVSCOMPhone string `json:"CVSCOMPhone,omitempty"` // 取貨人手機號碼
	LgsNo       string `json:"LgsNo,omitempty"`       // 物流訂單序號
	// 跨境支付回傳參數
	ChannelID string `json:"ChannelID,omitempty"` // 跨境通路類型
	ChannelNo string `json:"ChannelNo,omitempty"` // 跨境通路交易序號
}
