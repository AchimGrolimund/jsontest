package compliance

import (
	"encoding/json"
	"io/ioutil"
	"jsontest/models/compliance"
	"os"
)

func checkIfComplianceDBExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

func ReadComplianceDB(file string) (*compliance.ComplianceDB, error) {
	if check := checkIfComplianceDBExists(file); check == false {
		return nil, nil
	}
	var db = new(compliance.ComplianceDB)
	result, _ := ioutil.ReadFile(file)
	if err := json.Unmarshal(result, db); err != nil {
		return nil, err
	}
	return db, nil
}

func WriteComplianceDB(file string, data interface{}) error {
	jsonString, _ := json.Marshal(data)
	if err := ioutil.WriteFile(file, jsonString, os.ModePerm); err != nil {
		return err
	}
	return nil
}
