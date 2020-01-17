package main

import (
	"image"
	"image/jpeg"
	"io"
	"os"
	"strings"
)

// UploadImage will forward the image path of the selected image.
func (a *WalletApplication) UploadImage() string {
	filePath := a.RT.Dialog.SelectFile()
	splitPath := strings.Split(filePath, "/")
	filename := splitPath[len(splitPath)-1]

	a.log.Info("Path to user uploaded image: " + filePath)
	err := CopyFile(filePath, a.paths.ImageDir+filename)
	if err != nil && filePath != "" {
		a.log.Errorf("Unable to copy image. ", err)
		a.sendError("Unable to change Image. ", err)
		return "None"
	}

	file, err := os.Open(filePath)
	if err != nil && filePath != "" {
		a.log.Errorf("Unable to open image. ", err)
		a.sendError("Unable to find Image on the path provided. ", err)
		return "None"
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		a.log.Info("Attempting to decode as JPEG")
		img, err = jpeg.DecodeConfig(file)
		if err != nil {
			a.log.Errorf("Unable to decode image configuration", err)
			a.sendError("Unable to change Image. ", err)
			return "None"
		}
	}

	a.log.Info("Uploaded image resolution is set to ", img.Height, "x", img.Width)

	if img.Height > 200 || img.Width > 200 {
		a.log.Warnf("Image resolution is too big. Needs to be lower than 200x200 ")

		return "None"
	}

	// TODO: Store filePath in persistent storage.

	return filename
}

// CopyFile the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
