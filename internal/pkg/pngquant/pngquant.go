package pngquant

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"github.com/jayden1228/optim/internal/pkg/config"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gogf/gf/os/gfile"
)

const (
	notFoundPngquant = "not found Pngquant install in path"
)

func Compress(input image.Image, speed string) (output[]byte, err error) {
	if err = speedCheck(speed); err != nil {
		return
	}

	var w bytes.Buffer
	err = png.Encode(&w, input)
	if err != nil {
		return
	}

	b := w.Bytes()
	output, err = CompressBytes(b, speed)
	if err != nil {
		return
	}
	return
}

func CompressBytes(input []byte, speed string) (output []byte, err error) {
	cmdPath := config.PngquantPath
	if !gfile.Exists(cmdPath) {
		return nil, errors.New(notFoundPngquant)
	}
	cmd := exec.Command(cmdPath, "-", "--speed", speed)
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	err = cmd.Run()

	if err != nil {
		return
	}

	output = o.Bytes()
	return
}

func speedCheck(speed string) (err error) {
	// conversion, as an aside, also forces the speed argument to be a number.
	speedInt, err := strconv.Atoi(speed)
	if err != nil {
		return
	}

	if speedInt > 10 {
		return fmt.Errorf("speed cannot exceed value of 10")
	}

	return
}
