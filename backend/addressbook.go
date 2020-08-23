package app

import (
	"encoding/json"

	"github.com/grvlle/constellation_wallet/backend/models"
)

// GetAddressBook will query all the contacts in the address book
func (a *WalletApplication) GetAddressBook() string {

	var contacts []models.Contact
	a.DB.Where("alias = ?", a.wallet.WalletAlias).Find(&contacts)

	addressBook, err := json.Marshal(contacts)
	if err != nil {
		a.log.Errorln("Unable to produce a JSON encoding from the retrieved DB records. Reason: ", err)
		a.sendError("Unable to produce a JSON encoding from the retrieved DB records. Reason: ", err)
		return ""
	}
	return string(addressBook)
}

// CreateContact will add a new contact to the address book
func (a *WalletApplication) CreateContact(address string, name string, tag string, description string) bool {
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

// UpdateContact will delete a contact from the address book
func (a *WalletApplication) UpdateContact(id uint, address string, name string, tag string, description string) bool {

	contact := &models.Contact{
		ID:          id,
		Address:     address,
		Name:        name,
		Tag:         tag,
		Description: description,
	}

	if contact == nil {
		return false
	}
	if err := a.DB.Save(&contact).Error; err != nil {
		a.log.Errorln("Unable to update the DB record with the existing contact. Reason: ", err)
		a.sendError("Unable to update the DB record with the existing contact. Reason: ", err)
		return false
	}
	return true
}

// DeleteContact will delete a contact from the address book
func (a *WalletApplication) DeleteContact(id uint) bool {
	var contact models.Contact
	if err := a.DB.First(&contact, id).Error; err != nil {
		a.log.Errorln("Unable to find the DB record for the contact to be deleted. Reason: ", err)
		a.sendError("Unable to find the DB record for the contact to be deleted. Reason: ", err)
		return false
	}

	a.DB.Delete(&contact)
	return true
}
