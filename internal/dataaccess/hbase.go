package dataaccess

import (
	"context"
	"errors"
	"os"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const USERS_TABLE string = "users"
const USERS_DATA_COLFAM string = "data"

const USERS_PASSWORD_COL string = "password"
const USERS_SALT_COL string = "salt"
const USERS_PUBKEY_COL string = "pubkey"
const USERS_PRIVKEY_COL string = "privkey"
const USERS_NAME_COL string = "name"
const USERS_EMAIL_COL string = "email"

// HBase client variable
var HBaseClient gohbase.Client
var host string

// initialize the HBase client
func init() {
	host = os.Getenv("HBASE_HOST")
	if host == "" {
		host = "localhost"
	}
	HBaseClient = gohbase.NewClient(host)
}

// Create user in the database
// Given username and password
func CreateUser(user User, hashedPassword, salt []byte) error {
	value := map[string]map[string][]byte{USERS_DATA_COLFAM: {
		USERS_PASSWORD_COL: hashedPassword,
		USERS_SALT_COL:     salt,
		USERS_PUBKEY_COL:   []byte(user.PubKey),
		USERS_PRIVKEY_COL:  []byte(user.PrivKey),
		USERS_NAME_COL:     []byte(user.Name),
		USERS_EMAIL_COL:    []byte(user.Email),
	}}
	putReq, err := hrpc.NewPutStr(context.Background(), USERS_TABLE, user.Username, value)
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Get a user from the database
// Given username
// Return a user struct
func GetUser(username string) (*User, error) {
	getReq, err := hrpc.NewGetStr(context.Background(), USERS_TABLE, username)
	if err != nil {
		return nil, err
	}
	result, err := HBaseClient.Get(getReq)
	if err != nil {
		return nil, err
	}
	if result.Cells == nil {
		return nil, errors.New("nil cells")
	}
	if len(result.Cells) == 0 {
		return nil, errors.New("user not found")
	}
	var user User
	user.Username = username
	for _, cell := range result.Cells {
		if string(cell.Family) == USERS_DATA_COLFAM {
			switch string(cell.Qualifier) {
			case USERS_PUBKEY_COL:
				user.PubKey = string(cell.Value)
			case USERS_PRIVKEY_COL:
				user.PrivKey = string(cell.Value)
			case USERS_NAME_COL:
				user.Name = string(cell.Value)
			case USERS_EMAIL_COL:
				user.Email = string(cell.Value)
			}
		}
	}
	return &user, nil
}

// Get user's password and salt
// Given username
// Return a password and salt byte slices
func GetUserPasswordAndSalt(username string) ([]byte, []byte, error) {
	getReq, err := hrpc.NewGetStr(context.Background(), USERS_TABLE, username)
	if err != nil {
		return nil, nil, err
	}
	result, err := HBaseClient.Get(getReq)
	if err != nil {
		return nil, nil, err
	}
	if result.Cells == nil || len(result.Cells) == 0 {
		return nil, nil, errors.New("wrong credentials")
	}
	var password, salt []byte
	for _, cell := range result.Cells {
		if string(cell.Family) == USERS_DATA_COLFAM {
			if string(cell.Qualifier) == USERS_PASSWORD_COL {
				password = cell.Value
			} else if string(cell.Qualifier) == USERS_SALT_COL {
				salt = cell.Value
			}
		}
	}
	return password, salt, nil
}

// Delete user from database
// Given username
// Return error
func DeleteUser(username string) error {
	delReq, err := hrpc.NewDelStr(context.Background(), USERS_TABLE, username, nil)
	if err != nil {
		return err
	}
	_, err = HBaseClient.Delete(delReq)
	return err
}

// Update user's password
// Given username, password and salt
// Return error
func UpdateUserPassword(username string, hashedPassword, salt []byte) error {
	value := map[string]map[string][]byte{USERS_DATA_COLFAM: {
		USERS_PASSWORD_COL: hashedPassword,
		USERS_SALT_COL:     salt,
	}}
	putReq, err := hrpc.NewPutStr(context.Background(), USERS_TABLE, username, value)
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}
