package utils

import (
	"fmt"
	"jsontest/models/compliance"
)

/*
	This creates only a example ComplianceDB
*/

func MakeData(db *compliance.ComplianceDB, n int) {
	currentCount := len(db.Users)
	if currentCount <= 0 {
		currentCount = 0
	}

	for i := currentCount; i < n+currentCount; i++ {
		dbs := compliance.Dashboards{
			ID:           fmt.Sprintf("%d", i),
			Name:         fmt.Sprintf("DB_Name%v", i),
			Deletetime:   i,
			SendInfoMail: true,
			SendWarnMail: false,
		}

		det := compliance.Detectors{
			ID:           fmt.Sprintf("%d", i),
			Name:         fmt.Sprintf("Det_Name_%v", i),
			Deletetime:   i,
			SendInfoMail: true,
			SendWarnMail: true,
		}

		tea := compliance.Teams{
			ID:           fmt.Sprintf("%d", i),
			Name:         fmt.Sprintf("Team_Name_%v", i),
			Deletetime:   i,
			SendInfoMail: false,
			SendWarnMail: false,
		}

		_dbs := []compliance.Dashboards{}
		_dbs = append(_dbs, dbs)

		_det := []compliance.Detectors{}
		_det = append(_det, det)

		_tea := []compliance.Teams{}
		_tea = append(_tea, tea)

		usr := compliance.Users{
			ID:         fmt.Sprintf("%d", i),
			Email:      fmt.Sprintf("email@%v.ch", i),
			Dashboards: _dbs,
			Detectors:  _det,
			Teams:      _tea,
		}

		db.Users = append(db.Users, usr)
	}
}
