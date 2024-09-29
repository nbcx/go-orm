package orm

import (
	"context"
	"time"

	"github.com/nbcx/go-orm/orm/internal/models"
)

// Invocation represents an "Orm" invocation
type Invocation struct {
	Method string
	// Md may be nil in some cases. It depends on method
	Md interface{}
	// the args are All arguments except context.Context
	Args []interface{}

	mi *models.ModelInfo
	// f is the Orm operation
	f func(ctx context.Context) []interface{}

	// insideTx indicates whether this is inside a transaction
	InsideTx    bool
	TxStartTime time.Time
	TxName      string
}

func (inv *Invocation) GetTableName() string {
	if inv.mi != nil {
		return inv.mi.Table
	}
	return ""
}

func (inv *Invocation) execute(ctx context.Context) []interface{} {
	return inv.f(ctx)
}

// GetPkFieldName return the primary key of this table
// if not found, "" is returned
func (inv *Invocation) GetPkFieldName() string {
	if inv.mi.Fields.Pk != nil {
		return inv.mi.Fields.Pk.Name
	}
	return ""
}
