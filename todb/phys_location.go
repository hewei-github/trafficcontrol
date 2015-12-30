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
	"gopkg.in/guregu/null.v3"
	"time"
)

type PhysLocation struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	ShortName   string      `db:"short_name" json:"shortName"`
	Address     string      `db:"address" json:"address"`
	City        string      `db:"city" json:"city"`
	State       string      `db:"state" json:"state"`
	Zip         string      `db:"zip" json:"zip"`
	Poc         null.String `db:"poc" json:"poc"`
	Phone       null.String `db:"phone" json:"phone"`
	Email       null.String `db:"email" json:"email"`
	Comments    null.String `db:"comments" json:"comments"`
	Region      int64       `db:"region" json:"region"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handlePhysLocation(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getPhysLocation(id)
	} else if method == "POST" {
		return postPhysLocation(payload)
	} else if method == "PUT" {
		return putPhysLocation(id, payload)
	} else if method == "DELETE" {
		return delPhysLocation(id)
	}
	return nil, nil
}

func getPhysLocation(id int) (interface{}, error) {
	ret := []PhysLocation{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from phys_location where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from phys_location"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postPhysLocation(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO phys_location("
	sqlString += "name"
	sqlString += ",short_name"
	sqlString += ",address"
	sqlString += ",city"
	sqlString += ",state"
	sqlString += ",zip"
	sqlString += ",poc"
	sqlString += ",phone"
	sqlString += ",email"
	sqlString += ",comments"
	sqlString += ",region"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:short_name"
	sqlString += ",:address"
	sqlString += ",:city"
	sqlString += ",:state"
	sqlString += ",:zip"
	sqlString += ",:poc"
	sqlString += ",:phone"
	sqlString += ",:email"
	sqlString += ",:comments"
	sqlString += ",:region"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putPhysLocation(id int, payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwirte the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE phys_location SET "
	sqlString += "name = :name"
	sqlString += ",short_name = :short_name"
	sqlString += ",address = :address"
	sqlString += ",city = :city"
	sqlString += ",state = :state"
	sqlString += ",zip = :zip"
	sqlString += ",poc = :poc"
	sqlString += ",phone = :phone"
	sqlString += ",email = :email"
	sqlString += ",comments = :comments"
	sqlString += ",region = :region"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delPhysLocation(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM phys_location WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
