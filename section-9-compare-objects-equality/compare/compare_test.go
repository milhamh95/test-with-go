package compare

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {
	arg := 4
	want := 16
	got := Square(arg)
	if got != want {
		t.Errorf("Square(%d) = %d; want %d", arg, got, want)
	}
}

func TestDog(t *testing.T) {
	morty := Dog{
		Name: "Morty",
		Age:  3,
	}

	morty2 := Dog{
		Name: "Morty",
		Age:  3,
	}

	t.Logf("morty=%p, morty2=%p", &morty, &morty2)
	if morty != morty2 {
		t.Errorf("morty != morty2")
	}
	oddie := Dog{
		Name: "Odie",
		Age:  12,
	}
	if morty == oddie {
		t.Errorf("morty == odie")
	}
}

func TestPtr(t *testing.T) {
	morty := &DogWithFn{
		Name: "Morty",
		Age:  3,
	}
	morty2 := &DogWithFn{
		Name: "Morty",
		Age:  3,
	}
	t.Logf("morty=%p, morty2=%p", morty, morty2)
	if !reflect.DeepEqual(morty, morty2) {
		t.Errorf("morty != morty2")
	}
}
