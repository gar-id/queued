/*
Copyright © 2024 Gar (me@gar.my.id)
*/
package main

import (
	"github.com/gar-id/queued/apps/server/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load dot env file if available
	godotenv.Load(".env")

	cmd.Execute()
}
