package main

import (
	"fmt"
	"os"
  stdimage "image"
  "image/jpeg"
  "path/filepath"
  "time"
  "sync"
  "strings"
)

var counter int

type pixel struct {
  r, g, b, a uint32
}

type image struct {
  name string
  pixels []pixel
  width int
  height int
}

func main() {
  start := time.Now()
  images := getImages()

  for i, image := range images {
    for j, pixel := range image.pixels {
      fmt.Println("Image: ", i, "\t Pixel: ", j, "\t r g b a: ", pixel)
      if j==10 {
        break
      }
    }
  }
  fmt.Println("PIXELS EXAMANINED: ",counter)
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

//load images
func loadImage(filename string) stdimage.Image {
  f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
  if(err!=nil) {
    panic(err)
  }
  return img
}

//get images
func getImages() []image{

  paths, err := getPaths()
  if err != nil {
		panic(err)
	}

  var mu sync.Mutex
  var wg sync.WaitGroup
  wg.Add(len(paths))

  var images []image
  for _, path := range paths {
    go func(p string) {
      pixels := getPixels(p)

      mu.Lock()
      {
        images = append(images, pixels)
      }
      mu.Unlock()
      wg.Done()
    }(path)
	}
  wg.Wait()
  return images
}

//get pixels
func getPixels(path string) image {
  img := loadImage(path)
  bounds := img.Bounds()
  pixels := make([]pixel, bounds.Dx() * bounds.Dy())

  for i:=0; i<bounds.Dx() * bounds.Dy(); i++ {
    x := i%bounds.Dx()
    y := i/bounds.Dx()
    r, g, b, a := img.At(x, y).RGBA()
    pixels[i].r = r
    pixels[i].g = g
    pixels[i].b = b
    pixels[i].a = a
    counter++
  }

  xs := strings.Split(path, "/")
	name := xs[(len(xs) - 1):][0]
	image := image{
		name:   name,
		pixels: pixels,
		width:  bounds.Dx(),
		height: bounds.Dy(),
	}

  return image
}

//get paths
func getPaths() ([]string, error) {
  const dir = "images/"
  var paths []string

  wf := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
  }

  if err := filepath.Walk(dir, wf); err != nil {
		return nil, err
  }

  return paths, nil
}
