package form

import (
	"fmt"
	"testing"
)

func TestFields(t *testing.T) {
	tests := []struct {
		strct interface{}
		want  field
	}{
		{
			// value in field
			strct: struct {
				FullName string
			}{},
			want: field{
				Label:       "FullName",
				Name:        "FullName",
				Type:        "text",
				Placeholder: "FullName",
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if got != tc.want {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})
	}
}
