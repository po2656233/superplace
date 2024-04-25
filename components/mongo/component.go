package superMongo

import (
	"context"
	"fmt"
	exReflect "github.com/po2656233/superplace/extend/reflect"
	"time"

	cprofile "github.com/po2656233/superplace/config"
	face "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	Component struct {
		face.Component
		dbMap map[string]map[string]*mongo.Database
	}

	// HashDb hash by group id
	HashDb func(dbMaps map[string]*mongo.Database) string
)

func NewComponent() *Component {
	return &Component{
		dbMap: make(map[string]map[string]*mongo.Database),
	}
}

func (*Component) Name() string {
	return exReflect.GetPackName(Component{})
}

func (s *Component) Init() {
	// load only the database contained in the `db_id_list`
	mongoIdList := s.App().Settings().Get("mongo_id_list")
	if mongoIdList.LastError() != nil || mongoIdList.Size() < 1 {
		clog.Warnf("[nodeId = %s] `mongo_id_list` property not exists.", s.App().NodeId())
		return
	}

	mongoConfig := cprofile.GetConfig("mongo")
	if mongoConfig.LastError() != nil {
		panic("`mongo` property not exists in config file.")
	}

	for _, groupId := range mongoConfig.Keys() {
		s.dbMap[groupId] = make(map[string]*mongo.Database)

		dbGroup := mongoConfig.GetConfig(groupId)
		for i := 0; i < dbGroup.Size(); i++ {
			item := dbGroup.GetConfig(i)

			var (
				enable  = item.GetBool("enable", true)
				id      = item.GetString("db_id")
				dbName  = item.GetString("db_name")
				uri     = item.GetString("uri")
				timeout = time.Duration(item.GetInt64("timeout", 3)) * time.Second
			)

			for j := 0; j < mongoIdList.Size(); j++ {
				dbId := mongoIdList.Get(j).ToString()
				if id != dbId {
					continue
				}

				if !enable {
					panic(fmt.Sprintf("[dbName = %s] is disabled!", dbName))
				}

				db, err := CreateDatabase(uri, dbName, timeout)
				if err != nil {
					panic(fmt.Sprintf("[dbName = %s] create mongodb fail. error = %s", dbName, err))
				}

				s.dbMap[groupId][id] = db
				clog.Infof("[dbGroup =%s, dbName = %s] is connected.", groupId, id)
			}
		}
	}
}

func CreateDatabase(uri, dbName string, timeout ...time.Duration) (*mongo.Database, error) {
	tt := 3 * time.Second

	if len(timeout) > 0 && timeout[0].Seconds() > 3 {
		tt = timeout[0]
	}

	o := options.Client().ApplyURI(uri)
	if err := o.Validate(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), tt)
	defer cancel()

	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	clog.Infof("ping db [uri = %s] is ok", uri)

	return client.Database(dbName), nil
}

func (s *Component) GetDb(id string) *mongo.Database {
	for _, group := range s.dbMap {
		for k, v := range group {
			if k == id {
				return v
			}
		}
	}
	return nil
}

func (s *Component) GetHashDb(groupId string, hashFn HashDb) (*mongo.Database, bool) {
	dbGroup, found := s.GetDbMap(groupId)
	if !found {
		clog.Warnf("groupId = %s not found.", groupId)
		return nil, false
	}

	dbId := hashFn(dbGroup)
	db, found := dbGroup[dbId]
	return db, found
}

func (s *Component) GetDbMap(groupId string) (map[string]*mongo.Database, bool) {
	dbGroup, found := s.dbMap[groupId]
	return dbGroup, found
}
