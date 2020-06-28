package app

import (
	"encoding/json"

	"github.com/grvlle/constellation_wallet/backend/models"
)

// GetAddressBook will query all the contacts in the address book
func (a *WalletApplication) GetAddressBook() string {

	var contacts []models.Contact
	a.DB.Find(&contacts)

	addressBook, err := json.Marshal(contacts)
	if err != nil {
		return ""
	}
	n := len(addressBook)
	s := string(addressBook[:n])
	return s
}

// StoreContact will add a new contact to the address book
func (a *WalletApplication) StoreContact(address string, name string, tag string, description string) bool {
	contact := &models.Contact{
		Address:     address,
		Name:        name,
		Tag:         tag,
		Description: description,
	}

	if contact == nil {
		return false
	}
	if err := a.DB.Model(&a.wallet).Where("wallet_alias = ?", a.wallet.WalletAlias).Association("Contact").Append(contact).Error; err != nil {
		a.log.Errorln("Unable to update the DB record with the new contact. Reason: ", err)
		a.sendError("Unable to update the DB record with the new contact. Reason: ", err)
		return false
	}
	a.log.Infoln("Successfully stored contact in DB")
	return true

}
