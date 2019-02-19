package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Agency
//----------------------------------------------------
func CreateAgency(database *gorm.DB) {
	if !database.HasTable(&ent.TableAgency{}) {

		database.CreateTable(&ent.TableAgency{})
		database.Model(&ent.TableAgency{}).Association("agencyaddresses")
		database.Model(&ent.TableAgency{}).Association("agencycontacts")
	}
}
