package form

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFields(t *testing.T) {
	tests := []struct {
		strct interface{}
		want  []field
	}{
		{
			// value in field
			strct: struct {
				Name string
			}{},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "",
				},
			},
		},
		{
			// value in field
			strct: struct {
				Name  string
				Email string
				Age   int
			}{
				Name:  "Jon Calhoun",
				Email: "jon@calhoun.io",
				Age:   123,
			},
			want: []field{
				{
					Label:       "Name",
					Name:        "Name",
					Type:        "text",
					Placeholder: "Name",
					Value:       "Jon Calhoun",
				},
				{
					Label:       "Email",
					Name:        "Email",
					Type:        "text",
					Placeholder: "Email",
					Value:       "jon@calhoun.io",
				},
				{
					Label:       "Age",
					Name:        "Age",
					Type:        "text",
					Placeholder: "Age",
					Value:       123,
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%T", tc.strct), func(t *testing.T) {
			got := fields(tc.strct)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("fields() = %v; want %v", got, tc.want)
			}
		})
	}
}
