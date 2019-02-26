package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Agency Address
//----------------------------------------------------
func CreateAllocationStus(database *gorm.DB) {
	if !database.HasTable(&ent.TableAllocationStatus{}) {

		database.CreateTable(&ent.TableAllocationStatus{})
	}
}
