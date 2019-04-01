package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateOrganisation(database *gorm.DB) {

	if !database.HasTable(&ent.TableOrganisation{}) {

		database.CreateTable(&ent.TableHospital{})
		database.Model(&ent.TableOrganisation{}).AddUniqueIndex("idx_table_organisation_organisationabbr", "organisationabbr")
		database.Model(&ent.TableOrganisation{}).AddUniqueIndex("idx_table_organisation_organisationname", "organisationname")
		database.Model(&ent.TableOrganisation{}).Association("Hospitals")
	}
}
