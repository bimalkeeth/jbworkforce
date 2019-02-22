package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Agency Address
//----------------------------------------------------
func CreateAllowanceGroup(database *gorm.DB) {
	if !database.HasTable(&ent.TableAllowanceGroup{}) {

		database.CreateTable(&ent.TableAllowanceGroup{})
		database.Model(&ent.TableAllowanceGroup{}).Association("alowances")
	}
}
