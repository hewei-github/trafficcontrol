// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package todb

import (
	"encoding/json"
	"fmt"
	"time"
)

type Regex struct {
	Id          int64     `db:"id" json:"id"`
	Pattern     string    `db:"pattern" json:"pattern"`
	Type        int64     `db:"type" json:"type"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleRegex(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getRegex(id)
	} else if method == "POST" {
		return postRegex(payload)
	} else if method == "PUT" {
		return putRegex(id, payload)
	} else if method == "DELETE" {
		return delRegex(id)
	}
	return nil, nil
}

func getRegex(id int) (interface{}, error) {
	ret := []Regex{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from regex where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from regex"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postRegex(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO regex("
	sqlString += "pattern"
	sqlString += ",type"
	sqlString += ") VALUES ("
	sqlString += ":pattern"
	sqlString += ",:type"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putRegex(id int, payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwirte the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE regex SET "
	sqlString += "pattern = :pattern"
	sqlString += ",type = :type"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delRegex(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM regex WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
