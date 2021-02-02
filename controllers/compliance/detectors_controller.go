package compliance

import "jsontest/models/compliance"

// FindDetector return true if the Detector exist by the user and id
//
//	detectorID := "123XY"
//	userID := "xyz123"
//	if compliance.FindDetector(&database, userID, detectorID ){
//		fmt.Printf("Detector with ID: %s exists on User with ID: %s", detectorID, userID)
//	}
func FindDetector(compliDB *compliance.ComplianceDB, userID string, detectorID string) bool {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Detectors {
			detector := &user.Detectors[j]
			if detector.ID == detectorID {
				return true
			}
		}
		return false
	}
	return false
}

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

func DeleteDetector(compliDB *compliance.ComplianceDB, userID string, detectorID string) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}

		for j := range user.Detectors {
			detector := &user.Detectors[j]
			if detector.ID == detectorID {
				user.Detectors = append(user.Detectors[:j], user.Detectors[j+1:]...)
				break
			}
		}
		break
	}
}
