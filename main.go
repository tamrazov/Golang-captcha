package main

import (
    "github.com/joho/godotenv"
    "os"
    "image"
		"image/png"
		"image/color"
		"golang.org/x/image/font"
		"golang.org/x/image/font/basicfont"
		"golang.org/x/image/math/fixed"
		"math/rand"
		"time"
)

func main() {
  _ = godotenv.Load(".env"); // start watch .env
	DIR_CAPTCHA := os.Getenv("DIR_CAPTCHA"); // parsed .env and create new const DIR_CAPTCHA
  CreateCap(DIR_CAPTCHA);
}

func CreateCap(pathImg string)  {
	text := StringWithCharset(4,charset) // Generate random string by custom method StringWithCharset()
	path :=  pathImg + text + ".png"; // Create path
	img := image.NewRGBA(image.Rect(0, 0, 100, 50)) // Create rectangle 100px x 50px
	addLabel(img, 20, 30, text) // Add random text in image
	f, err := os.Create(path) // Create file .png
	if err != nil {
		panic(err);
	}
	defer f.Close(); // Add file in folder
	if err := png.Encode(f, img); err != nil {
		panic(err);
	}
}
	
// This is a custom function for add text in image
func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{200, 100, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
			Dst:  img,
			Src:  image.NewUniform(col),
			Face: basicfont.Face7x13,
			Dot:  point,
	}
	d.DrawString(label)
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// function for create random string
func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}
