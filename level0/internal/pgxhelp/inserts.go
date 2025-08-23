package pgxhelp

import (
	"context"
	"fmt"
	"strings"

	"github.com/ermyar/WbTechSchool/l0/internal/utils"
)

const (
	baseStr = "INSERT INTO %s VALUES "
)

func getArgList(len int) string {
	stb := strings.Builder{}
	stb.WriteByte('(')
	for i := 1; i < len; i++ {
		stb.WriteString(fmt.Sprintf("$%d ,", i))
	}
	stb.WriteString(fmt.Sprintf("$%d)", len))
	return stb.String()
}

func buildSqlStr(name string, len int) string {
	stb := strings.Builder{}
	stb.WriteString(fmt.Sprintf(baseStr, name))
	stb.WriteString(getArgList(len))
	return stb.String()
}

func (conn *PgConnection) Insert(ctx context.Context, name string, args ...interface{}) error {

	_, err := conn.conn.Exec(ctx, buildSqlStr(name, len(args)), args...)

	if err != nil {
		conn.log.Error("Insert: error occur", utils.SlogError(err))
		if err1 := conn.Ping(ctx); err1 == nil {
			return utils.ErrWrongData
		}
		return err
	}

	conn.log.Info("Inserted succesfully")

	return nil
}
