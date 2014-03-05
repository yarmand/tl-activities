package main

import (
  "./config/development"
  "os"
)

func Environment() string {
  env := os.Getenv("GO_ENV")
  if len(env) != 0 {
    return env
  } else {
    return "development"
  }
}

type Config struct {
}

func Init() {
  switch Environment() {
  case "development":
    development.Init()
  }
}
