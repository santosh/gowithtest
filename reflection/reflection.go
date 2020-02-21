package reflection

import "reflect"

// Person has a field Name and a struct Profile.
type Person struct {
	Name    string
	Profile Profile
}

// Profile in local to Person and has fields Age and City.
type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}

		if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
