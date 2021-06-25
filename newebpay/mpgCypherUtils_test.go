package newebpay_test

import (
	"github.com/stretchr/testify/assert"
	"go-newebpay-sdk/newebpay"
	"testing"
)

func TestMPG_CypherUtils(t *testing.T) {

	t.Run("TradeInfoEncrypt", func(t *testing.T) {
		merKey := "12345678901234567890123456789012"
		merIv := "1234567890123456"
		plainText := "MerchantID=3430112&RespondType=JSON&TimeStamp=1485232229&Version=1.4&MerchantOrderNo=S_1485232229&Amt=40&ItemDesc=UnitTest"
		expectEncrypted := "ff91c8aa01379e4de621a44e5f11f72e4d25bdb1a18242db6cef9ef07d80b0165e476fd1d9acaa53170272c82d122961e1a0700a7427cfa1cf90db7f6d6593bbc93102a4d4b9b66d9974c13c31a7ab4bba1d4e0790f0cbbbd7ad64c6d3c8012a601ceaa808bff70f94a8efa5a4f984b9d41304ffd879612177c622f75f4214fa"

		encryptedString := newebpay.TradeInfoEncrypt(plainText, merKey, merIv)

		assert.Equal(t, encryptedString, expectEncrypted)
	})

	t.Run("TradeInfoDecrypt", func(t *testing.T) {
		merKey := "12345678901234567890123456789012"
		merIv := "1234567890123456"
		expectPlainText := "MerchantID=3430112&RespondType=JSON&TimeStamp=1485232229&Version=1.4&MerchantOrderNo=S_1485232229&Amt=40&ItemDesc=UnitTest"
		encrypted := "ff91c8aa01379e4de621a44e5f11f72e4d25bdb1a18242db6cef9ef07d80b0165e476fd1d9acaa53170272c82d122961e1a0700a7427cfa1cf90db7f6d6593bbc93102a4d4b9b66d9974c13c31a7ab4bba1d4e0790f0cbbbd7ad64c6d3c8012a601ceaa808bff70f94a8efa5a4f984b9d41304ffd879612177c622f75f4214fa"

		decryptedText := newebpay.TradeInfoDecrypt(encrypted, merKey, merIv)

		assert.Equal(t, decryptedText, expectPlainText)
	})

	t.Run("TradeInfoHash", func(t *testing.T) {
		merKey := "12345678901234567890123456789012"
		merIv := "1234567890123456"
		plainText := "MerchantID=3430112&RespondType=JSON&TimeStamp=1485232229&Version=1.4&MerchantOrderNo=S_1485232229&Amt=40&ItemDesc=UnitTest"
		exceptHash := "EA0A6CC37F40C1EA5692E7CBB8AE097653DF3E91365E6A9CD7E91312413C7BB8"

		result := newebpay.TradeInfoHash(plainText, merKey, merIv)

		assert.Equal(t, result, exceptHash)
	})

}
