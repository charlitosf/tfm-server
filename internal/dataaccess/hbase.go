package dataaccess

import (
	"charlitosf/tfm-server/pkg/stringutilities"
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
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
const USERS_TOTP_COL string = "totp"

const PASSWORDS_TABLE string = "passwords"
const PASSWORDS_DATA_COLFAM string = "data"

const FILES_TABLE string = "files"
const FILES_DATA_COLFAM string = "data"

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
		USERS_TOTP_COL:     []byte(user.TOTPinfo),
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
			case USERS_TOTP_COL:
				user.TOTPinfo = string(cell.Value)
			}
		}
	}
	return &user, nil
}

// Get user's password and salt
// Given username
// Return a password and salt byte slices
func GetUserPasswordAndSalt(username string) ([]byte, []byte, error) {
	// Filter by password and salt
	family := map[string][]string{USERS_DATA_COLFAM: {USERS_PASSWORD_COL, USERS_SALT_COL}}
	getReq, err := hrpc.NewGetStr(context.Background(), USERS_TABLE, username, hrpc.Families(family))
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
		if string(cell.Qualifier) == USERS_PASSWORD_COL {
			password = cell.Value
		} else if string(cell.Qualifier) == USERS_SALT_COL {
			salt = cell.Value
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

// ---- PASSWORDS ----

// Get a password related to a username from a website from the database
// Given proprietary user, website and username
// Return a password string
func GetPassword(user, website, username string) (string, error) {
	passwords, err := GetPasswords(user, website)
	if err != nil {
		return "", err
	}
	for uname, password := range passwords {
		if uname == username {
			return password, nil
		}
	}
	return "", errors.New("password not found")
}

// Get passwords from website from database
// Given propietary user and website
// Return a map of usernames and passwords
func GetPasswords(user, website string) (map[string]string, error) {
	// Reverse website
	website = stringutilities.ReverseSplitJoin(website)
	// Filter by website
	family := map[string][]string{PASSWORDS_DATA_COLFAM: {website}}
	getReq, err := hrpc.NewGetStr(context.Background(), PASSWORDS_TABLE, user, hrpc.Families(family))
	if err != nil {
		return nil, err
	}
	result, err := HBaseClient.Get(getReq)
	if err != nil {
		return nil, err
	}
	if result.Cells == nil || len(result.Cells) == 0 {
		return make(map[string]string), nil
	}
	var passwords map[string]string = make(map[string]string)
	for _, cell := range result.Cells {
		if string(cell.Qualifier) == website {
			err = json.Unmarshal(cell.Value, &passwords)
			if err != nil {
				return nil, err
			}
		}
	}
	return passwords, nil
}

// Get all passwords from a user from the database
// Given propietary user
// Return a map of websites and passwords
func GetAllPasswords(user string) (map[string]map[string]string, error) {
	// Get whole row
	getReq, err := hrpc.NewGetStr(context.Background(), PASSWORDS_TABLE, user)
	if err != nil {
		return nil, err
	}
	result, err := HBaseClient.Get(getReq)
	if err != nil {
		return nil, err
	}
	if result.Cells == nil || len(result.Cells) == 0 {
		return make(map[string]map[string]string), nil
	}
	var passwords map[string]map[string]string = make(map[string]map[string]string)
	for _, cell := range result.Cells {
		if string(cell.Family) == PASSWORDS_DATA_COLFAM {
			var websitePasswords map[string]string = make(map[string]string)
			err = json.Unmarshal(cell.Value, &websitePasswords)
			if err != nil {
				return nil, err
			}
			passwords[stringutilities.ReverseSplitJoin(string(cell.Qualifier))] = websitePasswords
		}
	}
	return passwords, nil
}

// Create password in the database
// Given propietary user, website, username and password
// Return error
func CreatePassword(propietaryUser, website, username, password string) error {
	passwords, err := GetPasswords(propietaryUser, website)
	if err != nil {
		return err
	}

	// Reverse website
	website = stringutilities.ReverseSplitJoin(website)

	passwords[username] = password
	value, err := json.Marshal(passwords)
	if err != nil {
		return err
	}
	putReq, err := hrpc.NewPutStr(context.Background(), PASSWORDS_TABLE, propietaryUser,
		map[string]map[string][]byte{PASSWORDS_DATA_COLFAM: {
			website: value,
		}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Delete a password from the database
// Given propietary user, website and username
// Return error
func DeletePassword(propietaryUser, website, username string) error {
	passwords, err := GetPasswords(propietaryUser, website)
	if err != nil {
		return err
	}

	// Reverse website
	website = stringutilities.ReverseSplitJoin(website)

	delete(passwords, username)
	value, err := json.Marshal(passwords)
	if err != nil {
		return err
	}
	putReq, err := hrpc.NewPutStr(context.Background(), PASSWORDS_TABLE, propietaryUser,
		map[string]map[string][]byte{PASSWORDS_DATA_COLFAM: {
			website: value,
		}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Update a password in the database
// Given propietary user, website, username and password
// Return error
func UpdatePassword(propietaryUser, website, username, password string) error {
	passwords, err := GetPasswords(propietaryUser, website)
	if err != nil {
		return err
	}

	// Reverse website
	website = stringutilities.ReverseSplitJoin(website)

	passwords[username] = password
	value, err := json.Marshal(passwords)
	if err != nil {
		return err
	}
	putReq, err := hrpc.NewPutStr(context.Background(), PASSWORDS_TABLE, propietaryUser,
		map[string]map[string][]byte{PASSWORDS_DATA_COLFAM: {
			website: value,
		}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Create encoded file in the database
// Given propietary user, filename and file data
// Return error
func CreateFile(propietaryUser, filename string, data string) error {
	putReq, err := hrpc.NewPutStr(context.Background(), FILES_TABLE, propietaryUser,
		map[string]map[string][]byte{FILES_DATA_COLFAM: {
			filename: []byte(data),
		}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Get a file from the database
// Given propietary user and filename
// Return encoded file data
func GetFile(propietaryUser, filename string) (string, error) {
	getReq, err := hrpc.NewGetStr(context.Background(), FILES_TABLE, propietaryUser, hrpc.Families(map[string][]string{FILES_DATA_COLFAM: {filename}}))
	if err != nil {
		return "", err
	}
	result, err := HBaseClient.Get(getReq)
	if err != nil {
		return "", err
	}
	if result.Cells == nil || len(result.Cells) != 1 {
		return "", errors.New("file not found")
	}
	return string(result.Cells[0].Value), nil
}

// Delete a file from the database
// Given propietary user and filename
// Return error
func DeleteFile(propietaryUser, filename string) error {
	delReq, err := hrpc.NewDelStr(context.Background(), FILES_TABLE, propietaryUser, map[string]map[string][]byte{FILES_DATA_COLFAM: {filename: {}}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Delete(delReq)
	return err
}

// Update a file in the database
// Given propietary user, filename and file data
// Return error
func UpdateFile(propietaryUser, filename string, data string) error {
	putReq, err := hrpc.NewPutStr(context.Background(), FILES_TABLE, propietaryUser,
		map[string]map[string][]byte{FILES_DATA_COLFAM: {
			filename: []byte(data),
		}})
	if err != nil {
		return err
	}
	_, err = HBaseClient.Put(putReq)
	return err
}

// Get all filenames from the database
// Given propietary user
// Return list of filenames
func GetAllFilenames(propietaryUser string) ([]map[string]string, error) {
	// Filter by row
	filter := filter.NewRowFilter(filter.NewCompareFilter(filter.Equal, filter.NewSubstringComparator(propietaryUser)))
	scanReq, err := hrpc.NewScanStr(context.Background(), FILES_TABLE, hrpc.Filters(filter))
	if err != nil {
		return nil, err
	}
	scanner := HBaseClient.Scan(scanReq)
	var filenames []map[string]string = make([]map[string]string, 0)
	for result, err := scanner.Next(); err == nil; result, err = scanner.Next() {
		for _, cell := range result.Cells {
			if string(cell.Family) == FILES_DATA_COLFAM {
				var filename map[string]string = make(map[string]string)
				filename["name"] = string(cell.Qualifier)
				filenames = append(filenames, filename)
			}
		}
	}
	return filenames, nil
}
