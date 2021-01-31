package compliance

import (
	"jsontest/models/compliance"
)

func UpdateDashboards(compliDB *compliance.ComplianceDB, userID string, newDashboard *compliance.Dashboards) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Dashboards {
			dashboard := &user.Dashboards[j]
			if dashboard.ID == newDashboard.ID {
				dashboard.Name = newDashboard.Name
				dashboard.Deletetime = newDashboard.Deletetime
				break
			}
		}
		break
	}
}

func AddDashboard(compliDB *compliance.ComplianceDB, userID string, newDashboard *compliance.Dashboards) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		user.Dashboards = append(user.Dashboards, *newDashboard)
		break
	}
}

func DeleteDashboard(compliDB *compliance.ComplianceDB, userID string, _id string) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}

		for j := range user.Dashboards {
			dashboard := &user.Dashboards[j]
			if dashboard.ID == _id {
				user.Dashboards = append(user.Dashboards[:j], user.Dashboards[j+1:]...)
				break
			}
		}
		break
	}
}
