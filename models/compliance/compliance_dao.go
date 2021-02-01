package compliance

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*
	Data access Model
*/
const (
	filepath = "complianceDB.json"
)

func checkIfComplianceDBExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func (db *ComplianceDB) Load() error {
	if check := checkIfComplianceDBExists(filepath); check == false {
		return nil
	}
	result, _ := ioutil.ReadFile(filepath)
	if err := json.Unmarshal(result, db); err != nil {
		return err
	}
	return nil
}

func (db *ComplianceDB) Save() error {
	jsonString, _ := json.Marshal(db)
	if err := ioutil.WriteFile(filepath, jsonString, os.ModePerm); err != nil {
		return err
	}
	return nil
}
