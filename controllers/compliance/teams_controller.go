package compliance

import "jsontest/models/compliance"

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
