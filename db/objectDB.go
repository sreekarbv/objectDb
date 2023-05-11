package db

import (
	"context"
	"encoding/json"
	"endor/model"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"strings"
)

type ObjectDB interface {
	// Store will store the object in the data store. The object will have a
	// name and kind, and the Store method should create a unique ID.
	Store(ctx context.Context, object model.Object) error
	// GetObjectByID will retrieve the object with the provided ID.
	GetObjectByID(ctx context.Context, id string) (model.Object, error)
	// GetObjectByName will retrieve the object with the given name.
	GetObjectByName(ctx context.Context, name string) (model.Object, error)
	// ListObjects will return a list of all objects of the given kind.
	ListObjects(ctx context.Context, kind string) ([]model.Object, error)
	// DeleteObject will delete the object.
	DeleteObject(ctx context.Context, id string) error
}

type RedisDB struct {
	client *redis.Client
}

func NewRedisDB() *RedisDB {
	redisDb := new(RedisDB)
	redisDb.client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // host:port of the redis server
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	return redisDb
}
func (rDb RedisDB) Store(ctx context.Context, object model.Object) error {
	id := uuid.NewString()
	object.SetID(id)
	name := object.GetName()
	kind := object.GetKind()

	fmt.Println("Storing Object ", object)

	rDb.client.Set(id, name, 0)
	names, err := rDb.client.Get(kind).Result()
	if err == redis.Nil {
		fmt.Println(rDb.client.Set(kind, name, 0))
	} else if err == nil {
		names = names + "," + name
		fmt.Println(rDb.client.Set(kind, names, 0))
	} else {
		fmt.Println("Error while Getting names for kind ", kind, err)
		return err
	}

	jsonObj, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(rDb.client.Set(object.GetName(), jsonObj, 0))
	return nil
}

func (rDb RedisDB) GetObjectByID(ctx context.Context, id string) (model.Object, error) {
	name, err := rDb.client.Get(id).Result()
	if err != nil {
		fmt.Printf("Error %v while getting name for id %s\n", err, id)
		return nil, err
	}
	fmt.Printf("Name %s found for Id %s\n", name, id)
	return rDb.GetObjectByName(ctx, name)
}

func (rDb RedisDB) GetObjectByName(ctx context.Context, name string) (retObj model.Object, err error) {
	if name == "" {
		err = errors.New("Name is empty")
		return nil, err
	}
	fmt.Println("Getting by Name ", name)
	retVal, err := rDb.client.Get(name).Bytes()
	if err != nil {
		fmt.Println("Error while getting ", err)
	}
	retObj, err = model.UnmarshallObject(retVal)
	return
}

func (rDb RedisDB) ListObjects(ctx context.Context, kind string) ([]model.Object, error) {
	var retObjs []model.Object
	namesStr, err := rDb.client.Get(kind).Result()
	if err != nil {
		fmt.Printf("Error %v while getting names for kind %s\n", err, kind)
		return nil, err
	}
	fmt.Println("Found ", namesStr)
	names := strings.Split(namesStr, ",")
	for _, name := range names {
		fmt.Println("Calling GetObjectByName with ", name)
		obj, err := rDb.GetObjectByName(ctx, strings.TrimSpace(name))
		if err != nil {
			fmt.Printf("Error while getting object of Kind %s with Name %s - %v\n", kind, name, err)
			continue
		}
		fmt.Println("Appending object", obj)
		retObjs = append(retObjs, obj)
	}

	return retObjs, nil
}

func (rDb RedisDB) DeleteObject(ctx context.Context, id string) error {
	obj, err := rDb.GetObjectByID(ctx, id)
	if err != nil {
		fmt.Printf("Error %v while fetching object with id %s for deleting\n", err, id)
		return err
	}

	kind := obj.GetKind()
	namesStr, err := rDb.client.Get(kind).Result()
	if err != nil {
		fmt.Printf("Error %v while getting names for kind %s\n", err, kind)
		return err
	}
	delName := obj.GetName()
	names := strings.Split(namesStr, ",")
	var newNames []string

	for _, name := range names {
		if name != delName {
			newNames = append(newNames, name)
		}
	}
	newNamesStr := newNames[0]
	newNamesAfterStart := newNames[1:]
	for _, name := range newNamesAfterStart {
		newNamesStr = newNamesStr + "," + name
	}
	rDb.client.Set(kind, newNamesStr, 0)

	rDb.client.Del(id)
	rDb.client.Del(obj.GetName())
	return err
}
