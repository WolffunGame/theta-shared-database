package builder

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// appendIfHasVal append key and val to map if value is not empty.
func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !utils.IsNil(val) {
		m[key] = val
	}
}
