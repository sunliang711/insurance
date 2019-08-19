package model

import (
	"fmt"
	"reflect"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/sunliang711/insurance/types"
)

// AddClaim TODO
// 2019/08/14 14:02:53
func AddClaim(c *types.Claim) error {
	rt := reflect.TypeOf(*c)
	count := rt.NumField()
	log.Debugf("num filed: %v", count)
	questionArr := make([]string, count)
	for i := 0; i < count; i++ {
		questionArr[i] = "?"
	}

	questionStr := strings.Join(questionArr, ",")
	sql := "insert into claim values (" + questionStr + ")"
	log.Debugf("sql: %v", sql)

	_, err := db.Exec(sql, c.Name, c.CardType, c.Number, c.Sex, c.Birthday, c.Phone, c.City, c.ClaimDate, c.ClaimReason, c.ClaimType, c.Description)
	if err != nil {
		msg := fmt.Sprintf("exec sql: %v error: %v", sql, err)
		log.Error(msg)
		return fmt.Errorf(msg)
	}
	log.Debugf("AddClaim OK")
	return nil
}
