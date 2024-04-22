package mapStructure

import (
	"fmt"
)

func ExampleDecode() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
	// Output:
	// mapstructure.Person{Name:"Mitchell", Age:91, Emails:[]string{"one", "two", "three"}, Extra:map[string]string{"twitter":"mitchellh"}}
}

func ExampleDecode_errors() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}

	var result Person
	err := Decode(input, &result)
	if err == nil {
		panic("should have an error")
	}

	fmt.Println(err.Error())
	// Output:
	// 5 error(s) decoding:
	//
	// * 'Age' expected type 'int', got unconvertible type 'string', value: 'bad value'
	// * 'Emails[0]' expected type 'string', got unconvertible type 'int', value: '1'
	// * 'Emails[1]' expected type 'string', got unconvertible type 'int', value: '2'
	// * 'Emails[2]' expected type 'string', got unconvertible type 'int', value: '3'
	// * 'Name' expected type 'string', got unconvertible type 'int', value: '123'
}

func ExampleDecode_metadata() {
	type Person struct {
		Name string
		Age  int
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":  "Mitchell",
		"age":   91,
		"email": "foo@bar.com",
	}

	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md Metadata
	var result Person
	config := &DecoderConfig{
		Metadata: &md,
		Result:   &result,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
		panic(err)
	}

	if err := decoder.Decode(input); err != nil {
		panic(err)
	}

	fmt.Printf("Unused keys: %#v", md.Unused)
	// Output:
	// Unused keys: []string{"email"}
}

func ExampleDecode_weaklyTypedInput() {
	type Person struct {
		Name   string
		Age    int
		Emails []string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON, generated by a weakly typed language
	// such as PHP.
	input := map[string]interface{}{
		"name":   123,                      // number => string
		"age":    "42",                     // string => number
		"emails": map[string]interface{}{}, // empty map => empty array
	}

	var result Person
	config := &DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &result,
	}

	decoder, err := NewDecoder(config)
	if err != nil {
		panic(err)
	}

	err = decoder.Decode(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
	// Output: mapstructure.Person{Name:"123", Age:42, Emails:[]string{}}
}

func ExampleDecode_tags() {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	type Person struct {
		Name string `mapstructure:"person_name"`
		Age  int    `mapstructure:"person_age"`
	}

	input := map[string]interface{}{
		"person_name": "Mitchell",
		"person_age":  91,
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
	// Output:
	// mapstructure.Person{Name:"Mitchell", Age:91}
}

func ExampleDecode_embeddedStruct() {
	// Squashing multiple embedded structs is allowed using the squash tag.
	// This is demonstrated by creating a composite struct of multiple types
	// and decoding into it. In this case, a person can carry with it both
	// a Family and a Location, as well as their own FirstName.
	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		Family    `mapstructure:",squash"`
		Location  `mapstructure:",squash"`
		FirstName string
	}

	input := map[string]interface{}{
		"FirstName": "Mitchell",
		"LastName":  "Hashimoto",
		"City":      "San Francisco",
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s %s, %s", result.FirstName, result.LastName, result.City)
	// Output:
	// Mitchell Hashimoto, San Francisco
}

func ExampleDecode_remainingData() {
	// Note that the mapstructure tags defined in the struct type
	// can indicate which fields the values are mapped to.
	type Person struct {
		Name  string
		Age   int
		Other map[string]interface{} `mapstructure:",remain"`
	}

	input := map[string]interface{}{
		"name":  "Mitchell",
		"age":   91,
		"email": "mitchell@example.com",
	}

	var result Person
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", result)
	// Output:
	// mapstructure.Person{Name:"Mitchell", Age:91, Other:map[string]facade {}{"email":"mitchell@example.com"}}
}

func ExampleDecode_omitempty() {
	// NewActor omitempty annotation to avoid map keys for empty values
	type Family struct {
		LastName string
	}
	type Location struct {
		City string
	}
	type Person struct {
		*Family   `mapstructure:",omitempty"`
		*Location `mapstructure:",omitempty"`
		Age       int
		FirstName string
	}

	result := &map[string]interface{}{}
	input := Person{FirstName: "Somebody"}
	err := Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", result)
	// Output:
	// &map[Age:0 FirstName:Somebody]
}
