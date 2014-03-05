package main

import (
  "./activities"
  "./activities/models"
)

func main() {
  Init()
  models.CreateModelsTables()
  activities.StartServer()
}
