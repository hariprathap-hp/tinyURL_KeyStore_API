package dom_keys

import (
	"fmt"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/client/dbCassandra"
)

const (
	queryPopulateKeyDB = "INSERT into keys_avail (token_id) values (?)"
	queryGetKey        = "select token_id from keys_avail limit 1"
)

func (k *Key) Get() (*Key, *errors.RestErr) {
	iter := dbCassandra.GetSession().Query(queryGetKey).Iter()

	m := map[string]interface{}{}
	var results []Key
	for iter.MapScan(m) {
		results = append(results, Key{
			Token: m["token_id"].(string),
		})
		m = map[string]interface{}{}
	}
	if err := iter.Close(); err != nil {
		fmt.Println(err.Error())
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &results[0], nil
}

func (k *Key) Populate() *errors.RestErr {
	fmt.Println("Inside Populate")
	for i := 0; i < 25000; i++ {
		id := getID()
		fmt.Println(id)
		if err := dbCassandra.GetSession().Query(queryPopulateKeyDB, id).Exec(); err != nil {
			return errors.NewInternalServerError(err.Error())
		}
	}

	return nil
}
