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

	database, err := compliance2.ComplianceDB.Load()
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

	compliance.DeleteDashboard(database, "1", "0")

	if err := database.Save(); err != nil {
		log.Fatalln(err)
	}

}
