package newebpay_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-newebpay-sdk/newebpay"
	"testing"
	"time"
)

func TestJsonPayTime(t *testing.T) {

	t.Run("create JsonPayTime", func(t *testing.T) {
		wayback := time.Date(2020, time.May, 19, 1, 2, 3, 4, time.UTC)
		var jsonTime = newebpay.JsonPayTime(wayback)
		assert.Equal(t, jsonTime.Format(time.RFC3339), wayback.Format(time.RFC3339))
	})

	t.Run("Unmarshal JsonPayTime", func(t *testing.T) {
		jsonStr := "{\"PayTime\":\"2021-06-25 18:27:10\"}"
		expectedTime := time.Date(2021, time.June, 25, 18, 27, 10, 0, time.UTC)

		var testObj struct {
			PayTime newebpay.JsonPayTime
		}
		err := json.Unmarshal([]byte(jsonStr), &testObj)
		if err != nil {
			panic(err)
		}

		result := testObj.PayTime

		assert.Equal(t, result.Format(time.RFC3339), expectedTime.Format(time.RFC3339))
	})

	t.Run("Marshal JsonPayTime", func(t *testing.T) {
		var testObj struct {
			PayTime newebpay.JsonPayTime
		}
		testObj.PayTime = newebpay.JsonPayTime(time.Date(2021, time.June, 25, 18, 27, 10, 0, time.UTC))
		expectedJson := "{\"PayTime\":\"2021-06-25 18:27:10\"}"

		jsonByte, err := json.Marshal(testObj)
		if err != nil {
			panic(err)
		}

		result := string(jsonByte)

		assert.Equal(t, result, expectedJson)
	})
}
