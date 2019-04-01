package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// TableAward
//----------------------------------------------------
func CreateTableGender(database *gorm.DB) {
	if !database.HasTable(&ent.TableGender{}) {

		database.CreateTable(&ent.TableGender{})
	}
}
