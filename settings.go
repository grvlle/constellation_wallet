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

	if img.Height >= 201 || img.Width >= 201 {
		a.log.Warnf("Image resolution is too big. Cannot be bigger than 200x200 ")

		return "None"
	}
	a.StoreImagePathInDB(filename)
	return filename
}

// SetImagePath is called from the Login.Vue. It'll query the DB for the user's profile picture
// and return it to the FE to be displayed.
func (a *WalletApplication) SetImagePath() string {
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Error; err != nil {
		a.log.Errorf("Unable to query the DB record for the Image path. Reason: ", err)
		a.sendError("Unable to query the DB record for the Image path. Reason: ", err)
		return ""
	}
	a.log.Infoln("Profile Picture selected: ", a.wallet.ProfilePicture)
	return a.wallet.ProfilePicture
}

func (a *WalletApplication) StoreImagePathInDB(path string) {
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("ProfilePicture", path).Error; err != nil {
		a.log.Errorf("Unable to update the DB record with the Image path. Reason: ", err)
		a.sendError("Unable to update the DB record with the Image path. Reason: ", err)
	}
}

// SetWalletTag is called from the Login.Vue
func (a *WalletApplication) SetWalletTag() string {
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Error; err != nil {
		a.log.Errorf("Unable to query the DB record for the Image path. Reason: ", err)
		a.sendError("Unable to query the DB record for the Image path. Reason: ", err)
	}
	a.log.Infoln("Wallet Tag selected: ", a.wallet.WalletTag)
	return a.wallet.WalletTag
}

func (a *WalletApplication) StoreWalletLabelInDB(walletTag string) {
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Update("WalletTag", walletTag).Error; err != nil {
		a.log.Errorf("Unable to update the DB record with the wallet tag. Reason: ", err)
		a.sendError("Unable to update the DB record with the wallet tag. Reason: ", err)
	}
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
