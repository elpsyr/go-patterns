package flyweight

import "testing"

func ExampleImageFlyweight_Data() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
}

func TestImageViewer_Display(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fatal()

	}
}
