package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"os/exec"
	"time"

	"github.com/spwashi/golang-images/utils"
	"github.com/spwashi/golang-images/utils/image/output"
	"github.com/spwashi/golang-images/utils/image/processing/grid"
)

func main() {
	imageName, quantX, quantY := utils.ReadArguments()

	file, err := os.Open(imageName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()

	filePaths := []string{}
	for i := quantX; i > 0; i -= 1 {
		xGap := i
		yGap := quantY
		fileName := generateImage(img, bounds, xGap, yGap)
		filePaths = append(filePaths, fileName)
	}
	for i := quantY - 1; i > 0; i -= 1 {
		xGap := 1
		yGap := i
		fileName := generateImage(img, bounds, xGap, yGap)
		filePaths = append(filePaths, fileName)
	}

	runFfmpeg(26)
}

func generateImage(img image.Image, bounds image.Rectangle, xGap int, yGap int) string {
	pixelGrid := grid.MakePixelGrid(img, bounds, xGap, yGap)
	outputImage := output.MakeImage(pixelGrid, bounds)
	outputImageName := nameImage(xGap, yGap)
	filePath := output.WriteJpeg(outputImage, outputImageName)
	fmt.Println(filePath)
	return filePath
}

func runFfmpeg(framerate int) {

	cmd := exec.Command(
		"ffmpeg",
		"-y",
		"-framerate", fmt.Sprint(framerate),
		"-pattern_type", "glob",
		"-i", "output/images/*.jpg",
		"-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2",
		"-c:v", "libx264",
		"-pix_fmt", "yuv420p",
		"output/animation."+fmt.Sprint(framerate)+".mp4",
	)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func nameImage(quantX int, quantY int) string {
	outputImageName := time.Now().Format("20060102_1545") + "--" + fmt.Sprint(quantX) + "x" + fmt.Sprint(quantY)
	return outputImageName
}
