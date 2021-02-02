package compliance

import (
	"jsontest/models/compliance"
)

// FindDashbord return true if the Dashboard exist by the user and id
//
//	dashboardID := "123XY"
//	userID := "xyz123"
//	if compliance.FindDashbord(&database, userID, dashboardID ){
//		fmt.Printf("Dashbord with ID: %s exists on User with ID: %s", dashboardID, userID)
//	}
func FindDashbord(compliDB *compliance.ComplianceDB, userID string, dashboardID string) bool {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Dashboards {
			dashboard := &user.Dashboards[j]
			if dashboard.ID == dashboardID {
				return true
			}
		}
		return false
	}
	return false
}

// UpdateDashboards Update a existand Dasboard with an given new type of Dashboard
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

func DeleteDashboard(compliDB *compliance.ComplianceDB, userID string, dashboardID string) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}

		for j := range user.Dashboards {
			dashboard := &user.Dashboards[j]
			if dashboard.ID == dashboardID {
				user.Dashboards = append(user.Dashboards[:j], user.Dashboards[j+1:]...)
				break
			}
		}
		break
	}
}
