package relations

import (
	db "jbworkforce/masterdataservice/master_dataaccess/common"
	con "jbworkforce/masterdataservice/master_dataaccess/contracts"
	entities "jbworkforce/masterdataservice/master_dataaccess/relations/entityrelations"
	"log"
)

type RelationBuilder struct{}

func (*RelationBuilder) BuildDatabase(connInfo *con.ClientInfo) (err error) {
	var ss db.Idb = &db.Db{}
	database, err := ss.OpenDatabase(connInfo)
	if err != nil {
		log.Fatal("database connection was not successful.")
		return err
	}
	database.SingularTable(true)
	entities.CreateActivityTable(database)
	entities.CreateAddressType(database)
	entities.CreateState(database)
	entities.CreateAddressType(database)
	entities.CreateAgency(database)
	entities.CreateAgencyAddress(database)
	entities.CreateAddress(database)
	entities.CreateAgencyContacts(database)
	entities.CreateAllocationStatus(database)
	entities.CreateAllowanceGroup(database)
	entities.CreateAllowance(database)
	entities.CreateAllowanceLevel(database)
	entities.CreateAvailabilityStatus(database)
	entities.CreateTableAward(database)
	entities.CreateCampus(database)
	entities.CreateClientSkillGroup(database)
	entities.CreateClientSkillType(database)
	entities.CreateHealthGroup(database)
	entities.CreateClinicalStream(database)
	entities.CreateClinicalStreamAttributes(database)
	entities.CreateCompetencyType(database)
	entities.CreateCompetencyCompletionStatus(database)
	entities.CreateContactType(database)
	entities.CreateOrganisation(database)
	entities.CreateHospital(database)
	entities.CreateCostCentre(database)
	entities.CreateFinancialYear(database)
	entities.CreateDateWeek(database)
	entities.CreateDivision(database)
	entities.CreateDepartment(database)
	entities.CreateTableGender(database)
	return nil
}
