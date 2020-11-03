package api

import "gorm.io/gorm"

func (api *API) InitAuthors(database *gorm.DB) {
	db = database

}
