package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field",
			struct {
				Name string
			}{"Santosh"},
			[]string{"Santosh"},
		},
		{"struct with two string field",
			struct {
				Name string
				City string
			}{"Santosh", "Pune"},
			[]string{"Santosh", "Pune"},
		},
		{"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Santosh", 23},
			[]string{"Santosh"},
		},
		{"nested fields",
			Person{"Santosh",
				Profile{23, "Pune"},
			},
			[]string{"Santosh", "Pune"},
		},
		{"pointers to things",
			&Person{"Santosh",
				Profile{23, "Pune"},
			},
			[]string{"Santosh", "Pune"},
		},
		{"using slices",
			[]Profile{
				{23, "Pune"},
				{26, "Bengaluru"},
			},
			[]string{"Pune", "Bengaluru"},
		},
		{"using arrays",
			[2]Profile{
				{23, "Pune"},
				{26, "Bengaluru"},
			},
			[]string{"Pune", "Bengaluru"},
		},
		{"using maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
