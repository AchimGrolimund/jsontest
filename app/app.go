package app

import (
	"jsontest/controllers/compliance"
	compliance2 "jsontest/models/compliance"
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

	updateDate := compliance2.Dashboards{
		ID:         "9",
		Name:       "Updated Dashboardname",
		Deletetime: "1234567890",
	}

	compliance.UpdateDashboards(result, "9", &updateDate)

	if err := compliance.WriteComplianceDB(filepath, result); err != nil {
		log.Fatalln(err)
	}

}
