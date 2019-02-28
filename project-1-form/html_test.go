package form_test

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	form "test-with-go/project-1-form"
	"testing"
)

var (
	tplTypeNameValue = template.Must(template.New("").Parse(`<input type="{{.Type}}" name="{{.Name}}"{{with .Value}} value="{{.}}"{{end}}>`))
	tplAll           = template.Must(template.New("").Parse(`
	<label>{{.Label}}</label>
	<input
		type="{{.Type}}"
		name="{{.Name}}"
		placeholder="{{.Placeholder}}"
		{{with .Value}}value="{{.}}"{{end}}>`))
)

var updateFlag bool

func init() {
	flag.BoolVar(&updateFlag, "update", false, "set the update flag to update the expected output of all golden file tests run")

}

func TestHTML(t *testing.T) {
	tests := map[string]struct {
		tpl     *template.Template
		strct   interface{}
		want    string
		wantErr error
	}{
		"A basic form with values": {
			tpl: tplTypeNameValue,
			strct: struct {
				Name  string
				Email string
			}{
				Name:  "Michael Scott",
				Email: "michael@dundermifflin.com",
			},
			want: "TestHTML_basic.golden",
		},
		"A form with custom struct tags": {
			tpl: tplAll,
			strct: struct {
				LabelTest       string `form:"label=This is custom"`
				NameTest        string `form:"name=full_name"`
				TypeTest        int    `form:"type=number"`
				PlaceholderTest string `form:"placeholder=your value goes here..."`
				Nested          struct {
					MultiTest string `form:"name=NestedMulti;label=This is nested;type=email;placeholder=user@example.com"`
				}
			}{
				PlaceholderTest: "value and placeholder",
			},
			want: "TestHTML_structTags.golden",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := form.HTML(tc.tpl, tc.strct)
			if err != tc.wantErr {
				t.Fatalf("HTML() err = %v; want %v", err, tc.wantErr)
			}
			if updateFlag {
				writeFile(t, tc.want, string(got))
				t.Logf("Updated golden file %s", tc.want)
			}
			want := template.HTML(readFile(t, tc.want))
			if got != want {
				t.Errorf("HTML() - do not match golden file")
				gotFileName := strings.Replace(tc.want, ".golden", ".got", 1)
				writeFile(t, gotFileName, string(got))
				t.Errorf("To compare run : diff %s %s", gotFileName, tc.want)
			}
		})
	}
}

func writeFile(t *testing.T, filename, contents string) {
	f, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Error opening file %s: %v", filename, err)
	}
	defer f.Close()
	fmt.Fprint(f, contents)
}

func readFile(t *testing.T, filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Error opening file %s: %v", filename, err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("Error reading file %s; %v", filename, err)
	}
	return b
}
