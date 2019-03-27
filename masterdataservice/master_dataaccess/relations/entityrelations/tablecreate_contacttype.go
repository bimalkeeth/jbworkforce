package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateContactType(database *gorm.DB) {

	if !database.HasTable(&ent.TableContactType{}) {

		database.CreateTable(&ent.TableContactType{})
		database.Model(&ent.TableContactType{}).AddUniqueIndex("idx_table_contacttype_contacttype", "contacttype")
		database.Model(&ent.TableContactType{}).Association("Contacts")
	}
}
