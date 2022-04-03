package dom_keys

import (
	"fmt"
	"test3/hariprathap-hp/system_design/tinyURL/utils/errors"
	"test3/hariprathap-hp/system_design/tinyURL_KeyStore_API/client/dbCassandra"
)

const (
	queryPopulateKeyDB = "INSERT into keys (token_id) values (?)"
)

func (k *Key) Get() (*string, *errors.RestErr) {
	return nil, nil
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
