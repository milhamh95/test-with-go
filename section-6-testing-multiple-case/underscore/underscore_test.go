package underscore

import (
	"fmt"
	"testing"
)

type testCase struct {
	arg  string
	want string
}

// func TestCamel(t *testing.T) {
// 	testCases := []testCase{
// 		{
// 			"thisIsACamelCaseString",
// 			"this_is_a_camel_case_string",
// 		},
// 		{
// 			"with a space",
// 			"with a space",
// 		},
// 		{
// 			"endsWithA",
// 			"ends_with_a",
// 		},
// 	}
// 	for _, tt := range testCases {
// 		got := Camel(tt.arg)
// 		if got != tt.want {
// 			t.Errorf("Camel(%q) = %q; want %q", tt.arg, got, tt.want)
// 		}
// 	}

// }

func TestCamel(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"simple", args{"thisIsACamelCaseString"}, "this_is_a_camel_case_string"},
		{"spaces", args{"with a space"}, "with a space "},
		{"ends with capital", args{"endsWithA"}, "ends_with_a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Camel(tt.args.str); got != tt.want {
				t.Fatalf("Camel(%v) = %v, want %v", tt.args.str, got, tt.want)
			}
			fmt.Println("this won't print if it fails")
		})
	}
}
