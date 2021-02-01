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
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deletetime string `json:"deletetime"`
}
type Detectors struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deletetime string `json:"deletetime"`
}
type Teams struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deletetime string `json:"deletetime"`
}

/*
https://discord.com/channels/118456055842734083/118456055842734083/805382974995693609
*/
/* JSON Plain

{
	"users": [
		{
			"id": "ff",
			"email": "mail",
			"dashboards": [
				{
					"id": "",
					"name": "",
					"deletetime": ""
				}
			],
			"detectors": [
				{
					"id": "",
					"name": "",
					"deletetime": ""
				}
			],
			"teams": [
				{
					"id": "",
					"name": "",
					"deletetime": ""
				}
			]
		}
	]
}

*/
