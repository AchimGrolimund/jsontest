package compliance

type ComplianceDB struct {
	Users []Users `json:"users"`
}

type Users struct {
	ID         string       `json:"id"`
	Email      string       `json:"email"`
	Dashboards []Dashboards `json:"dashboards"`
	Detectors  []Detectors  `json:"detectors"`
	Teams      []Teams      `json:"teams"`
}

type Dashboards struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Deletetime   string `json:"deletetime"`
	SendInfoMail bool   `json:"send_info_mail"`
	SendWarnMail bool   `json:"send_warn_mail"`
}
type Detectors struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Deletetime   string `json:"deletetime"`
	SendInfoMail bool   `json:"send_info_mail"`
	SendWarnMail bool   `json:"send_warn_mail"`
}
type Teams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Deletetime   string `json:"deletetime"`
	SendInfoMail bool   `json:"send_info_mail"`
	SendWarnMail bool   `json:"send_warn_mail"`
}

/* JSON Plain

{
  "users": [
    {
      "id": "0",
      "email": "email@0.ch",
      "dashboards": [
        {
          "id": "0",
          "name": "DB_Name0",
          "deletetime": "0",
          "send_info_mail": true,
          "send_warn_mail": false
        }
      ],
      "detectors": [
        {
          "id": "0",
          "name": "Det_Name_0",
          "deletetime": "0",
          "send_info_mail": true,
          "send_warn_mail": true
        }
      ],
      "teams": [
        {
          "id": "0",
          "name": "Team_Name_0",
          "deletetime": "0",
          "send_info_mail": false,
          "send_warn_mail": false
        }
      ]
    },
    {
      "id": "1",
      "email": "email@1.ch",
      "dashboards": [
        {
          "id": "1",
          "name": "DB_Name1",
          "deletetime": "1",
          "send_info_mail": true,
          "send_warn_mail": false
        }
      ],
      "detectors": [
        {
          "id": "1",
          "name": "Det_Name_1",
          "deletetime": "1",
          "send_info_mail": true,
          "send_warn_mail": true
        }
      ],
      "teams": [
        {
          "id": "1",
          "name": "Team_Name_1",
          "deletetime": "1",
          "send_info_mail": false,
          "send_warn_mail": false
        }
      ]
    },
    {
      "id": "2",
      "email": "email@2.ch",
      "dashboards": [
        {
          "id": "2",
          "name": "DB_Name2",
          "deletetime": "2",
          "send_info_mail": true,
          "send_warn_mail": false
        }
      ],
      "detectors": [
        {
          "id": "2",
          "name": "Det_Name_2",
          "deletetime": "2",
          "send_info_mail": true,
          "send_warn_mail": true
        }
      ],
      "teams": [
        {
          "id": "2",
          "name": "Team_Name_2",
          "deletetime": "2",
          "send_info_mail": false,
          "send_warn_mail": false
        }
      ]
    },
    {
      "id": "3",
      "email": "email@3.ch",
      "dashboards": [
        {
          "id": "3",
          "name": "DB_Name3",
          "deletetime": "3",
          "send_info_mail": true,
          "send_warn_mail": false
        }
      ],
      "detectors": [
        {
          "id": "3",
          "name": "Det_Name_3",
          "deletetime": "3",
          "send_info_mail": true,
          "send_warn_mail": true
        }
      ],
      "teams": [
        {
          "id": "3",
          "name": "Team_Name_3",
          "deletetime": "3",
          "send_info_mail": false,
          "send_warn_mail": false
        }
      ]
    },
    {
      "id": "4",
      "email": "email@4.ch",
      "dashboards": [
        {
          "id": "4",
          "name": "DB_Name4",
          "deletetime": "4",
          "send_info_mail": true,
          "send_warn_mail": false
        }
      ],
      "detectors": [
        {
          "id": "4",
          "name": "Det_Name_4",
          "deletetime": "4",
          "send_info_mail": true,
          "send_warn_mail": true
        }
      ],
      "teams": [
        {
          "id": "4",
          "name": "Team_Name_4",
          "deletetime": "4",
          "send_info_mail": false,
          "send_warn_mail": false
        }
      ]
    }
  ]
}
*/
