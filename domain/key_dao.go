package dom_kgs

import (
	"fmt"
	"strings"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/client/dbCassandra"
	"test3/hariprathap-hp/system_design/utils_repo/errors"
)

const (
	queryPopulateKeyDB     = "INSERT into keys_avail (token_id) values (?)"
	queryPopulateKeyUsedDB = "INSERT into keys_used (token_id) values (?)"
	queryGetKey            = "select token_id from keys_avail limit ?"
	queryDeleteKey         = "delete from keys_avail WHERE token_id IN ?"
)

func (k *Key) Get(cnt int, isCache bool) ([]string, *errors.RestErr) {
	fmt.Println("select stmt -- \n", dbCassandra.GetSession().Query(queryGetKey, cnt))
	scanner := dbCassandra.GetSession().Query(queryGetKey, cnt).Iter().Scanner()

	var results []string
	var iter string
	for scanner.Next() {
		if err := scanner.Scan(&iter); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, iter)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	fmt.Println("Printing results -- ", results)

	if err := delete(results); err != nil {
		return nil, errors.NewInternalServerError(err.Message)
	}

	//prepare values to be removed in a single statement
	for i := 0; i < cnt; i++ {
		if err := k.Populate(results[i]); err != nil {
			return nil, errors.NewInternalServerError(err.Message)
		}
	}

	return results, nil
}

func (k *Key) Populate(count string) *errors.RestErr {
	if strings.Compare(count, "populate") == 0 {
		for i := 0; i < 25000; i++ {
			id := getID()
			fmt.Println(id)
			if err := dbCassandra.GetSession().Query(queryPopulateKeyDB, id).Exec(); err != nil {
				return errors.NewInternalServerError(err.Error())
			}
		}
	} else {
		if err := dbCassandra.GetSession().Query(queryPopulateKeyUsedDB, count).Exec(); err != nil {
			return errors.NewInternalServerError("Error while loading keys to used database")
		}
	}

	return nil
}

func delete(key []string) *errors.RestErr {
	fmt.Println("printing key\n", key)
	fmt.Println("Query is -- ", dbCassandra.GetSession().Query(queryDeleteKey, key))
	if err := dbCassandra.GetSession().Query(queryDeleteKey, key).Exec(); err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("error while deleting key from unused database")
	}
	return nil
}
