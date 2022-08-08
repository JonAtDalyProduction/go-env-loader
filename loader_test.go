package envloader

import (
	"testing"
)

//Test_createEnvMap tests creating a map of env variables from a .env file
func Test_createEnvMap(t *testing.T) {
	testMap := make(map[string]string)
	testMap["TEST_ONE"] = "1"
	testMap["TEST_TWO"] = "test_two"
	testMap["TEST_THREE"] = "test_three=test_three"
	testMap["TEST_INLINE_COMMENT"] = "test_inline_comment"
	got := createEnvMap("sample.env")
	for k, v := range testMap {
		if v != got[k] {
			t.Fatalf("createEnvMap unexpected value; want %q; got %q", testMap[k], got[k])
		}
	}
}
