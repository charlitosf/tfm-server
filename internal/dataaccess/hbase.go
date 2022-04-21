package dataaccess

import (
	"context"
	"errors"
	"os"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

// HBase client variable
var HBaseClient gohbase.Client
var host string

// initialize the HBase client
func init() {
	host = os.Getenv("HBASE_HOST")
	if host == "" {
		host = "192.168.22.132"
	}
	HBaseClient = gohbase.NewClient(host)
}

// Create user in the data base
// Given username and password
func CreateUser(username, password string) error {
	// return errors.New(host)
	value := map[string]map[string][]byte{"data": {"password": []byte(password)}}
	putReq, err := hrpc.NewPutStr(context.Background(), "users", username, value)
	if err != nil {
		return err
	}
	rsp, err := HBaseClient.Put(putReq)
	return errors.New(rsp.String())
}
