package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Address Type
//----------------------------------------------------
func CreateAddressType(database *gorm.DB) {

	if !database.HasTable(&ent.TableAddressType{}) {

		database.CreateTable(&ent.TableAddressType{})
		database.Model(&ent.TableAddressType{}).AddUniqueIndex("idx_table_addresstype_addresstype", "addresstype")
		database.Model(&ent.TableAddressType{}).Association("addresses")
	}
}
