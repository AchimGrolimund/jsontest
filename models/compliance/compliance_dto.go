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
    }
  ]
}
*/
