package main

import (
	"context"
	"cloud.google.com/go/bigtable"
)

func main() {
	ctx := context.Background()
	client, _ := bigtable.NewClient(ctx, "project", "instance")
	tbl := client.Open("user_profiles")
	
	mut := bigtable.NewMutation()
	mut.Set("segments", "tech", bigtable.Now(), []byte("true"))
	tbl.Apply(ctx, "user_123", mut)
}
