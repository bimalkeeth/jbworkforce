package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Allowance level
//----------------------------------------------------
func CreateAllowanceLevel(database *gorm.DB) {
	if !database.HasTable(&ent.TableAllowanceLevel{}) {
		database.CreateTable(&ent.TableAllowanceLevel{})
	}
}
