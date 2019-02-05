package common

import (
	"github.com/jinzhu/gorm"
	con "jbworkforce/masterdataservice/master_dataaccess/contracts"
)

type Idb interface {
	OpenDatabase(clientInfo *con.ClientInfo) (*gorm.DB, error)
}
