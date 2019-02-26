package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Agency Address
//----------------------------------------------------
func CreateAgencyAddress(database *gorm.DB) {
	if !database.HasTable(&ent.TableAgencyAddress{}) {

		database.CreateTable(&ent.TableAgencyAddress{})
		database.Model(&ent.TableAgencyAddress{}).AddForeignKey("agencyid", "table_agency(id)", "RESTRICT", "RESTRICT")
		database.Model(&ent.TableAgencyAddress{}).AddForeignKey("stateid", "table_state(id)", "RESTRICT", "RESTRICT")
	}
}
