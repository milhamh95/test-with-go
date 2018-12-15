package naming

import (
	"testing"
)

func TestDog(t *testing.T) {

}

func TestSpeak(t *testing.T) {

}

func TestDog_Bark(t *testing.T) {

}

func TestDog_Bark_muzzled(t *testing.T) {

}

func TestDog_Bark_unmuzzled(t *testing.T) {

}

func TestSpeak_spanish(t *testing.T) {

}

func TestColor(t *testing.T) {
	arg := "blue"
	want := "#0000FF"
	got := Color(arg)
	if got != want {
		t.Errorf("Color(%s) = %s; want %s", arg, got, want)
	}
}
