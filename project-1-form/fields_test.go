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
				Name string
			}{},
			want: field{
				Label:       "Name",
				Name:        "Name",
				Type:        "text",
				Placeholder: "Name",
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if got != tc.want {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})
	}
}
