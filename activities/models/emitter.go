package models

import (
  "time"
)

type Emitter struct {
  Id      int64
  Name    string
  Created time.Time
  Updated time.Time
}
