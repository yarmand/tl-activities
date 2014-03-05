package development

import (
  "github.com/coocood/qbs"
  _ "github.com/lib/pq"
  "log"
)

func RegisterDb() {
  qbs.Register("postgres", "postgres://yarmand@localhost/tl-activities?connect_timeout=10&sslmode=disable", "tl-activities", qbs.NewPostgres())
}

func Init() {
  log.Println("initializing development environment")
  RegisterDb()
}
