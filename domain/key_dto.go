package dom_keys

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type Key struct {
	Token string `json:"token_id"`
}

func getID() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Generate a snowflake ID.
	id := node.Generate()
	result := id.Base36()[3:]
	return result
}
