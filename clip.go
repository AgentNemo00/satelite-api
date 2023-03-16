package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sat-api/geometry"
	image2 "sat-api/image"
	"sat-api/model"
	"strings"
	"sync"
)

func ClipDir(path string, points [][]model.Point) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	pathToConvert := make([]string, 0)
	if !info.IsDir() {
		return fmt.Errorf("no dir")
	}
	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, entry := range dir {
		if !entry.IsDir() {
			pathToConvert = append(pathToConvert, entry.Name())
		}
	}
	wg := &sync.WaitGroup{}
	for i := 0; i < len(pathToConvert); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			err := ClipFile(filepath.Join(path, pathToConvert[index]), points[index])
			if err != nil {
				log.Println(err)
			}
		}(i)
	}
	wg.Wait()
	return nil
}

func ClipFile(path string, points []model.Point) error {
	imgFile, err := os.Open(path)
	defer imgFile.Close()
	if err != nil {
		return err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}
	xMin, yMin, xMax, yMax := image2.CalculateOptimizedSize(points, 0)
	vectors := make([]*geometry.Vector, 0)
	for _, point := range points {
		vectors = append(vectors, geometry.NewVector(point.X, point.Y, 0))
	}
	absolutPoly := geometry.NewPolygon(vectors)
	imagePoly := image2.NewImagePolygon(absolutPoly, image2.Bound{
		X: image2.Point{
			X: xMin,
			Y: xMax,
		},
		Y: image2.Point{
			X: yMin,
			Y: yMax,
		},
	})
	imgClipped := imagePoly.Clip(img, true)
	name := fmt.Sprintf("%s", filepath.Join(filepath.Dir(path),
		fmt.Sprintf("%s%s", fmt.Sprintf("%s-clipped", strings.TrimSuffix(filepath.Base(path),
			filepath.Ext(path))), ".png")))
	f, err := os.Create(name)
	defer f.Close()
	if err != nil {
		return err
	}
	return png.Encode(f, imgClipped)
}
