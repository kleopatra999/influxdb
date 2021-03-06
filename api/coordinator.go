package api

import (
	"github.com/Wikia/influxdb/cluster"
	cmn "github.com/Wikia/influxdb/common"
	"github.com/Wikia/influxdb/engine"
	"github.com/Wikia/influxdb/protocol"
)

// The following are the api that is accessed by any api

type Coordinator interface {
	// This is only used in the force compaction http endpoint
	ForceCompaction(cmn.User) error

	// Data related api
	RunQuery(cmn.User, string, string, engine.Processor) error
	RunQueryWithContext(cmn.User, string, string, engine.Processor, *RequestContext) error
	WriteSeriesData(cmn.User, string, []*protocol.Series) error

	// Administration related api
	CreateDatabase(cmn.User, string) error
	ListDatabases(cmn.User) ([]*cluster.Database, error)
	DropDatabase(cmn.User, string) error
	GetShardSpacesForDatabase(database string) []*cluster.ShardSpace
}
