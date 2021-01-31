package app

import (
	"jsontest/controllers/compliance"
	compliance2 "jsontest/models/compliance"
	"log"
)

const (
	filepath = "complianceDB.json"
)

func Start() {

	var complDB = new(compliance2.ComplianceDB)

	complDB, err := compliance.ReadComplianceDB(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		utils.MakeData(result, 5)

		if err := compliance.WriteComplianceDB(filepath, result); err != nil {
			log.Fatalln(err)
		}

	*/
	/*
		updateDate := compliance2.Dashboards{
			ID: "xyz",
			Name:       "Updated Dashboardname",
			Deletetime: "dfghghgh",
		}

	*/

	compliance.DeleteDashboard(complDB, "1", "0")

	if err := compliance.WriteComplianceDB(filepath, complDB); err != nil {
		log.Fatalln(err)
	}

}
