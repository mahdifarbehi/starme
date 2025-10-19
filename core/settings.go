package core

import "os"

var JWT_SECRET string = os.Getenv("DB_URL")
