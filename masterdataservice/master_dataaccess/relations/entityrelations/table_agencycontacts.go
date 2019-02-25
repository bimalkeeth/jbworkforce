package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Agency Address
//----------------------------------------------------
func CreateAgencyContacts(database *gorm.DB) {
	if !database.HasTable(&ent.TableAgencyContacts{}) {

		database.CreateTable(&ent.TableAgencyAddress{})
		database.Model(&ent.TableAgencyContacts{}).AddForeignKey("agencyid", "table_agency(id)", "RESTRICT", "RESTRICT")
		database.Model(&ent.TableAgencyContacts{}).AddForeignKey("contactid", "table_contacts(id)", "RESTRICT", "RESTRICT")
	}
}
