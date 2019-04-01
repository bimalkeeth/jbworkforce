package entityrelations

import (
	"github.com/jinzhu/gorm"
	ent "jbworkforce/masterdataservice/master_dataaccess/entities"
)

//----------------------------------------------------
// Contact Type
//----------------------------------------------------
func CreateFinancialYear(database *gorm.DB) {
	if !database.HasTable(&ent.TableFinancialYear{}) {
		database.CreateTable(&ent.TableFinancialYear{})
		database.Model(&ent.TableFinancialYear{}).AddUniqueIndex("idx_table_financialyear_financialyearname", "organisationabbr")
	}
}
