package pipeline_test

// import (
// 	"testing"

// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/require"

// 	"github.com/smartcontractkit/chainlink/core/services/pipeline"
// )

// func TestJSONParseTask(t *testing.T) {
// 	tests := []struct {
// 		name            string
// 		input           string
// 		path            []string
// 		wantData        interface{}
// 		wantResultError bool
// 	}{
// 		{"existing path", `{"high":"11850.00","last":"11779.99"}`, []string{"last"},
// 			"11779.99", false},
// 		{"nonexistent path", `{"high":"11850.00","last":"11779.99"}`, []string{"doesnotexist"},
// 			nil, false},
// 		{"double nonexistent path", `{"high":"11850.00","last":"11779.99"}`, []string{"no", "really"},
// 			nil, true},
// 		{"array index path", `{"data":[{"availability":"0.99991"}]}`, []string{"data", "0", "availability"},
// 			"0.99991", false},
// 		{"float result", `{"availability":0.99991}`, []string{"availability"},
// 			0.99991, false},
// 		{
// 			"index array",
// 			`{"data": [0, 1]}`,
// 			[]string{"data", "0"},
// 			float64(0),
// 			false,
// 		},
// 		{
// 			"index array of array",
// 			`{"data": [[0, 1]]}`,
// 			[]string{"data", "0", "0"},
// 			float64(0),
// 			false,
// 		},
// 		{
// 			"index of negative one",
// 			`{"data": [0, 1]}`,
// 			[]string{"data", "-1"},
// 			float64(1),
// 			false,
// 		},
// 		{
// 			"index of negative array length",
// 			`{"data": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]}`,
// 			[]string{"data", "-10"},
// 			float64(0),
// 			false,
// 		},
// 		{
// 			"index of negative array length minus one",
// 			`{"data": [0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55]}`,
// 			[]string{"data", "-12"},
// 			nil,
// 			false,
// 		},
// 		{
// 			"maximum index array",
// 			`{"data": [0, 1]}`,
// 			[]string{"data", "18446744073709551615"},
// 			nil,
// 			false,
// 		},
// 		{
// 			"overflow index array",
// 			`{"data": [0, 1]}`,
// 			[]string{"data", "18446744073709551616"},
// 			nil,
// 			false,
// 		},
// 		{
// 			"return array",
// 			`{"data": [[0, 1]]}`,
// 			[]string{"data", "0"},
// 			[]interface{}{float64(0), float64(1)},
// 			false,
// 		},
// 		{
// 			"return false",
// 			`{"data": false}`,
// 			[]string{"data"},
// 			false,
// 			false,
// 		},
// 		{
// 			"return true",
// 			`{"data": true}`,
// 			[]string{"data"},
// 			true,
// 			false,
// 		},
// 		{
// 			"regression test: keys in the path have dots",
// 			`{
//                 "Realtime Currency Exchange Rate": {
//                     "1. From_Currency Code": "LEND",
//                     "2. From_Currency Name": "EthLend",
//                     "3. To_Currency Code": "ETH",
//                     "4. To_Currency Name": "Ethereum",
//                     "5. Exchange Rate": "0.00058217",
//                     "6. Last Refreshed": "2020-06-22 19:14:04",
//                     "7. Time Zone": "UTC",
//                     "8. Bid Price": "0.00058217",
//                     "9. Ask Price": "0.00058217"
//                 }
//             }`,
// 			[]string{
// 				"Realtime Currency Exchange Rate",
// 				"5. Exchange Rate",
// 			},
// 			"0.00058217",
// 			false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		test := tt
// 		t.Run(test.name, func(t *testing.T) {
// 			transformer := job.JSONParseTask{Path: test.path}
// 			result, err := transformer.Transform(test.input)
// 			require.Equal(t, test.wantData, result)

// 			if test.wantResultError {
// 				require.Error(t, err)
// 			} else {
// 				require.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestMultiplyTask_Happy(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name  string
// 		input interface{}
// 		times decimal.Decimal
// 		want  decimal.Decimal
// 	}{
// 		{"string, by 100", "1.23", *mustDecimal(t, "100"), *mustDecimal(t, "123")},
// 		{"string, negative", "1.23", *mustDecimal(t, "-5"), *mustDecimal(t, "-6.15")},
// 		{"string, no times parameter", "1.23", *mustDecimal(t, "1"), *mustDecimal(t, "1.23")},
// 		{"string, zero", "1.23", *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"string, large value", "1.23", *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "1230000000000000000")},

// 		{"int, by 100", int(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"int, negative", int(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"int, no times parameter", int(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"int, zero", int(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"int, large value", int(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"int8, by 100", int8(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"int8, negative", int8(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"int8, no times parameter", int8(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"int8, zero", int8(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"int8, large value", int8(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"int16, by 100", int16(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"int16, negative", int16(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"int16, no times parameter", int16(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"int16, zero", int16(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"int16, large value", int16(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"int32, by 100", int32(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"int32, negative", int32(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"int32, no times parameter", int32(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"int32, zero", int32(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"int32, large value", int32(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"int64, by 100", int64(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"int64, negative", int64(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"int64, no times parameter", int64(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"int64, zero", int64(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"int64, large value", int64(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"uint, by 100", uint(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"uint, negative", uint(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"uint, no times parameter", uint(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"uint, zero", uint(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"uint, large value", uint(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"uint8, by 100", uint8(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"uint8, negative", uint8(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"uint8, no times parameter", uint8(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"uint8, zero", uint8(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"uint8, large value", uint8(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"uint16, by 100", uint16(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"uint16, negative", uint16(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"uint16, no times parameter", uint16(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"uint16, zero", uint16(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"uint16, large value", uint16(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"uint32, by 100", uint32(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"uint32, negative", uint32(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"uint32, no times parameter", uint32(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"uint32, zero", uint32(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"uint32, large value", uint32(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"uint64, by 100", uint64(2), *mustDecimal(t, "100"), *mustDecimal(t, "200")},
// 		{"uint64, negative", uint64(2), *mustDecimal(t, "-5"), *mustDecimal(t, "-10")},
// 		{"uint64, no times parameter", uint64(2), *mustDecimal(t, "1"), *mustDecimal(t, "2")},
// 		{"uint64, zero", uint64(2), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"uint64, large value", uint64(2), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "2000000000000000000")},

// 		{"float32, by 100", float32(1.23), *mustDecimal(t, "10"), *mustDecimal(t, "12.3")},
// 		{"float32, negative", float32(1.23), *mustDecimal(t, "-5"), *mustDecimal(t, "-6.15")},
// 		{"float32, no times parameter", float32(1.23), *mustDecimal(t, "1"), *mustDecimal(t, "1.23")},
// 		{"float32, zero", float32(1.23), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"float32, large value", float32(1.23), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "1230000000000000000")},

// 		{"float64, by 100", float64(1.23), *mustDecimal(t, "10"), *mustDecimal(t, "12.3")},
// 		{"float64, negative", float64(1.23), *mustDecimal(t, "-5"), *mustDecimal(t, "-6.15")},
// 		{"float64, no times parameter", float64(1.23), *mustDecimal(t, "1"), *mustDecimal(t, "1.23")},
// 		{"float64, zero", float64(1.23), *mustDecimal(t, "0"), *mustDecimal(t, "0")},
// 		{"float64, large value", float64(1.23), *mustDecimal(t, "1000000000000000000"), *mustDecimal(t, "1230000000000000000")},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		t.Run(test.name, func(t *testing.T) {
// 			transformer := job.MultiplyTask{Times: test.times}
// 			result, err := transformer.Transform(test.input)
// 			require.NoError(t, err)
// 			require.Equal(t, test.want.String(), result.(decimal.Decimal).String())
// 		})
// 	}
// }

// func TestMultiplyTask_Unhappy(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		times decimal.Decimal
// 		input interface{}
// 	}{
// 		{"map", *mustDecimal(t, "100"), map[string]interface{}{"chain": "link"}},
// 		{"slice", *mustDecimal(t, "100"), []interface{}{"chain", "link"}},
// 	}

// 	for _, tt := range tests {
// 		test := tt
// 		t.Run(test.name, func(t *testing.T) {
// 			transformer := job.MultiplyTask{Times: test.times}
// 			_, err := transformer.Transform(test.input)
// 			require.Error(t, err)
// 		})
// 	}
// }

// func mustDecimal(t *testing.T, arg string) *decimal.Decimal {
// 	ret, err := decimal.NewFromString(arg)
// 	require.NoError(t, err)
// 	return &ret
// }
