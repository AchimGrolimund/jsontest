package app

import (
	"jsontest/controllers/compliance"
	"jsontest/utils"
	"log"
)

const (
	filepath = "complianceDB.json"
)

func Start() {

	result, err := compliance.ReadComplianceDB(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	utils.MakeData(result, 5)

	if err := compliance.WriteComplianceDB(filepath, result); err != nil {
		log.Fatalln(err)
	}

}
