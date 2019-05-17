package helpers

import (
	"encoding/hex"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func ObjectIdHex(s string) (bson.ObjectId, error) {
	d, err := hex.DecodeString(s)
	if err != nil || len(d) != 12 {
		return "", fmt.Errorf("invalid input to ObjectIdHex: %q", s)
	}

	return bson.ObjectId(d), nil
}
