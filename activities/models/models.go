package models

import (
  //"errors"
  "github.com/coocood/qbs"
  _ "github.com/lib/pq"
  "log"
  "reflect"
)

func CreateModelTable(structPtr interface{}) error {
  migration, err := qbs.GetMigration()
  if err != nil {
    return err
  }
  defer migration.Close()
  log.Printf("creating Table %T\n", structPtr)
  return migration.CreateTableIfNotExists(structPtr)
}

func CreateModelsTables() {
  CreateModelTable(new(User))
  CreateModelTable(new(Activity))
  CreateModelTable(new(Emitter))
}

func SetField(object interface{}, name string, value interface{}) {
  s := reflect.ValueOf(object).Elem()
  typeOfObject := s.Type()
  for i := 0; i < s.NumField(); i++ {
    if typeOfObject.Field(i).Name == name {
      field := s.Field(i)
      v := reflect.ValueOf(value)
      field.Set(v)
    }
  }
}

func FindById(id int64, object interface{}) (interface{}, error) {
  q, err := qbs.GetQbs()
  if err != nil {
    return nil, err
  }
  SetField(object, "Id", id)
  return object, q.Find(object)
}

func Create(data map[string]interface{}, object interface{}) (interface{}, error) {
  for k := range data {
    SetField(object, k, data[k])
  }
  q, err := qbs.GetQbs()
  if err != nil {
    return object, err
  }
  _, err = q.Save(object)
  return object, err
}
