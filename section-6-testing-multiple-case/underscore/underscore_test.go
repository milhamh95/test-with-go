package underscore

import "testing"

type testCase struct {
	arg  string
	want string
}

func TestCamel(t *testing.T) {
	testCases := []testCase{
		{
			"thisIsACamelCaseString",
			"this_is_a_camel_case_string",
		},
		{
			"with a space",
			"with a space",
		},
		{
			"endsWithA",
			"ends_with_a",
		},
	}
	for _, tt := range testCases {
		got := Camel(tt.arg)
		if got != tt.want {
			t.Errorf("Camel(%q) = %q; want %q", tt.arg, got, tt.want)
		}
	}

}
