package naming

import "testing"

// Type_usecase
func TestDog_puppy(t *testing.T) {}

// Type_Method_usecase
func TestDog_Bark_muzzled(t *testing.T) {}

func TestDog_Bark_unmuzzled(t *testing.T) {}

// Function_usecase
func TestSpeak_spanish(t *testing.T) {}

// Naming variables: arg, got, want ...
func TestColor(t *testing.T) {
	arg := "blue"
	want := "#0000ff"
	got := Color(arg)
	if got != want {
		t.Errorf("Color(%s) = %s; want %s", arg, got, want)
	}
}
