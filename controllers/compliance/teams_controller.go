package compliance

import "jsontest/models/compliance"

// FindTeam return true if the Detector exist by the user and id
//
//	teamID := "123XY"
//	userID := "xyz123"
//	if compliance.FindTeam(&database, userID, teamID ){
//		fmt.Printf("Team with ID: %s exists on User with ID: %s", teamID, userID)
//	}
func FindTeam(compliDB *compliance.ComplianceDB, userID string, teamID string) bool {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Teams {
			team := &user.Teams[j]
			if team.ID == teamID {
				return true
			}
		}
		return false
	}
	return false
}

func UpdateTeams(compliDB *compliance.ComplianceDB, userID string, newTem *compliance.Teams) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		for j := range user.Teams {
			team := &user.Teams[j]
			if team.ID == newTem.ID {
				team.Name = newTem.Name
				team.Deletetime = newTem.Deletetime
				break
			}
		}
		break
	}
}

func AddTeam(compliDB *compliance.ComplianceDB, userID string, newTeam *compliance.Teams) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}
		user.Teams = append(user.Teams, *newTeam)
		break
	}
}

func DeleteTeam(compliDB *compliance.ComplianceDB, userID string, teamID string) {
	for i := range compliDB.Users {
		user := &compliDB.Users[i]
		if user.ID != userID {
			continue
		}

		for j := range user.Teams {
			team := &user.Teams[j]
			if team.ID == teamID {
				user.Teams = append(user.Teams[:j], user.Teams[j+1:]...)
				break
			}
		}
		break
	}
}
