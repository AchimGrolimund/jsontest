package compliance

import "jsontest/models/compliance"

func UpdateDetectors(compliDB *compliance.ComplianceDB, userID string, newDetector *compliance.Detectors) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Detectors {
			detector := &user.Detectors[j]
			if detector.ID == newDetector.ID {
				detector.Name = newDetector.Name
				detector.Deletetime = newDetector.Deletetime
				break
			}
		}
		break
	}
}

func AddDetector(compliDB *compliance.ComplianceDB, userID string, newDashboard *compliance.Detectors) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		user.Detectors = append(user.Detectors, *newDashboard)
		break
	}
}
