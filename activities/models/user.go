package models

import (
  "time"
)

type User struct {
  Id        int64
  Firstname string
  Lastname  string
  Username  string
  Created   time.Time
  Updated   time.Time
}
