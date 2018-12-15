package img

import "io"

// Image is a fake image type
type Image struct{}

// Decode is a fake function to decode images
func Decode(r io.Reader) (*Image, error) {

	// use std lib image package to read image &
	// parse to local image

	return &Image{}, nil
}

// Crop is a fake function to crop images
func Crop(img *Image, x1, y1, x2, y2 int) error {
	return nil
}

// Encode is a fake function to encode images
func Encode(img *Image, w io.Writer) error {

	// use std lib image package to decode the image
	return nil
}
