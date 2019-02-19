package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Address
//----------------------------------------------------
func CreateAddress(database *gorm.DB) {
	if !database.HasTable(&ent.TableAddress{}) {

		database.CreateTable(&ent.TableAddress{})
		database.Model(&ent.TableAddress{}).AddForeignKey("addresstypeid", "table_addresstype(id)", "RESTRICT", "RESTRICT")
		database.Model(&ent.TableAddress{}).AddForeignKey("stateid", "table_state(id)", "RESTRICT", "RESTRICT")
	}
}
