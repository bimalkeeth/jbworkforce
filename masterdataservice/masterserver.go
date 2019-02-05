package main

import (
	con "jbworkforce/masterdataservice/master_dataaccess/contracts"
	re "jbworkforce/masterdataservice/master_dataaccess/relations"
	"log"
)

func main() {
	var ff re.IRelation = &re.RelationBuilder{}

	err := ff.BuildDatabase(&con.ClientInfo{Client: "austin",
		UserName: "postgres",
		Password: "postgres",
		HostName: "localhost",
		Port:     5432})
	if err != nil {
		log.Fatal(err)
	}

}
