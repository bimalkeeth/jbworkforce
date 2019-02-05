package relations

import con "jbworkforce/masterdataservice/master_dataaccess/contracts"

type IRelation interface {
	BuildDatabase(connInfo *con.ClientInfo) (err error)
}
