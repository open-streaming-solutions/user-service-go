package util

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertUUIDtoString(ID pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", ID.Bytes[0:4], ID.Bytes[4:6], ID.Bytes[6:8], ID.Bytes[8:10], ID.Bytes[10:16])
}
