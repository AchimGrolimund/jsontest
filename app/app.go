package app

import (
	"jsontest/controllers/compliance"
	compliance2 "jsontest/models/compliance"
	"jsontest/utils"
	"log"
)

func Start() {
	var database = compliance2.ComplianceDB{}

	if err := database.Load(); err != nil {
		log.Fatalln(err)
	}

	utils.MakeData(&database, 5)

	updateDate := compliance2.Dashboards{
		ID:         "xyz",
		Name:       "Dashboardname",
		Deletetime: "ffffffff",
	}

	compliance.UpdateDashboards(&database, "1", &updateDate)

	if err := database.Save(); err != nil {
		log.Fatalln(err)
	}

}
