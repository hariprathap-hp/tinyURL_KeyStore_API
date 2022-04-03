package dom_keys

import (
	"fmt"
	"strings"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/client/dbCassandra"
)

const (
	queryPopulateKeyDB     = "INSERT into keys_avail (token_id) values (?)"
	queryPopulateKeyUsedDB = "INSERT into keys_used (token_id) values (?)"
	queryGetKey            = "select token_id from keys_avail limit ?"
	queryDeleteKey         = "delete from keys_avail where token_id = ?"
)

func (k *Key) Get(cnt int) (*Key, *errors.RestErr) {
	iter := dbCassandra.GetSession().Query(queryGetKey, cnt).Iter()

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

	//delete it from un_used keys table and move it to used keys table
	if err := delete(results[0].Token); err != nil {
		return nil, errors.NewInternalServerError(err.Message)
	}

	//move it to used_keys table
	k.Populate(results[0].Token)
	return &results[0], nil
}

func (k *Key) Populate(count string) *errors.RestErr {
	fmt.Println("Inside Populate")
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

func delete(key string) *errors.RestErr {
	if err := dbCassandra.GetSession().Query(queryDeleteKey, key).Exec(); err != nil {
		return errors.NewInternalServerError("error while deleting key from usused database")
	}
	return nil
}
