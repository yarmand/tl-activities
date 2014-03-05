package models

import (
  "time"
)

type Activity struct {
  Id        int64
  UserId    int64
  User      *User
  EmitterID int64
  Emitter   *Emitter
  Content   string
  Created   time.Time
  Updated   time.Time
}
