package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux" 

	bolt "github.com/boltdb/bolt"
)

//DBase type used for storing BoltDB
type DBase struct {
	DB       *bolt.DB
	Settings map[string]string
}

// Product struct used for storing product data
type Product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	IsBought bool   `json:"isBought"`
	DateAdded time.Time `json:"dateAdded"`
}

//User - name/pass - used for login/signup
type User struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

//Token used for tokens array. TODO: move Users and tokens to mongoDB
type Token struct {
	Name  string
	Token string
}

//Err used for error handling in http requests
type Err struct {
	Code int
	Text string
}

var products []Product
var dB DBase
var auth DBase

var tokens map[string]string

//var CurrBucket string

// Logger method for anything
func Logger(msg string, file string) {
	f, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(f)
	log.Println(msg + "\n")
	f.Close()
	return
}

// GetProductListEndpoint used for retriving all products in list
func GetProductListEndpoint(w http.ResponseWriter, req *http.Request) {
	auth := req.Header.Get("auth")
	if auth == "" {
		var err Err
		err.Code = 401
		err.Text = "Unauthorized"
		str, _ := json.Marshal(err)
		http.Error(w, string(str), 401)
		return
	}
	bucketName := tokens[auth]
	var prods []Product
	dB.DB.View(func(tx *bolt.Tx) error { //TODO: move to separate function
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var p Product
			json.Unmarshal(v, &p)
			prods = append(prods, p)
		}
		return nil
	})
	json.NewEncoder(w).Encode(prods)
}

// GenerateGUID generates UUID/GUID
func GenerateGUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// AddProductEndpoint used for creating new product in db
func AddProductEndpoint(w http.ResponseWriter, req *http.Request) {
	var pr Product 
	auth := req.Header.Get("auth")
	if auth == "" {
		var err Err
		err.Code = 401
		err.Text = "Unauthorized"
		str, _ := json.Marshal(err)
		http.Error(w, string(str), 401)
		return
	}
	json.NewDecoder(req.Body).Decode(&pr)
	bucketName := tokens[auth]
	if bucketName == "" {
		http.Error(w, "User not authorized", 401)
		return
	}
	pr.ID = GenerateGUID()
	pr.DateAdded = time.Now()
	dB.DB.Update(func(tx *bolt.Tx) error { //TODO: move to separate function
		prods, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
		temp, err := json.Marshal(pr)
		if err != nil {
			log.Println(err)
		}
		prods.Put([]byte(pr.ID), temp)
		return nil
	})
	json.NewEncoder(w).Encode(pr)
}

// DeleteProductEndpoint used for deleting old product by ID
func DeleteProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	auth := req.Header.Get("auth")
	if auth == "" {
		var err Err
		err.Code = 401
		err.Text = "Unauthorized"
		str, _ := json.Marshal(err)
		http.Error(w, string(str), 401)
		return
	}
	bucketName := tokens[auth]
	dB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.Delete([]byte(params["id"]))
		return nil
	})
}

// GetProductEndpoint get certain product by ID
func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	auth := req.Header.Get("auth")
	if auth == "" {
		var err Err
		err.Code = 401
		err.Text = "Unauthorized"
		str, _ := json.Marshal(err)
		http.Error(w, string(str), 401)
		return
	}
	bucketName := tokens[auth]
	dB.DB.View(func(tx *bolt.Tx) error { //TODO: move to separate function
		b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))
		resp := b.Get([]byte(params["id"]))
		json.NewEncoder(w).Encode(string(resp))
		return nil
	})
}

// ToggleProductEndpoint change product status by ID.
func ToggleProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req) 
	auth := req.Header.Get("auth")
	if auth == "" {
		var err Err
		err.Code = 401
		err.Text = "Unauthorized"
		str, _ := json.Marshal(err)
		http.Error(w, string(str), 401)
		return
	}
	var p Product
	bucketName := tokens[auth]
	dB.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		json.Unmarshal(b.Get([]byte(params["id"])), &p)
		p.IsBought = !p.IsBought
		tmp, _ := json.Marshal(p)
		b.Put([]byte(params["id"]), tmp)
		json.NewEncoder(w).Encode(p)
		return nil
	})
}

// SignUpEndpoint used for IDK, like killing bytes?
func SignUpEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	var err Err
	json.NewDecoder(req.Body).Decode(&user)
	pass := md5.New()
	io.WriteString(pass, user.Pass)
	passHash := pass.Sum(nil)
	//is there user with same name?
	auth.DB.View(func(tx *bolt.Tx) error {
		bb := tx.Bucket([]byte("Users"))
		resp := bb.Get([]byte(user.Name))
		if resp != nil {
			err.Code = 500
			err.Text = "User already registered"
			str, _ := json.Marshal(err)
			http.Error(w, string(str), 500)
		}
		return nil
	})
	if err.Text != "" {
		return
	}
	auth.DB.Update(func(tx *bolt.Tx) error {
		users, _ := tx.CreateBucketIfNotExists([]byte("Users"))
		users.Put([]byte(user.Name), passHash)
		return nil
	})
	json.NewEncoder(w).Encode(user)
}

// SignInEndpoint used for achieving auth token by user TODO: add DDOS guards.
func SignInEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	var err Err
	var tt Token

	json.NewDecoder(req.Body).Decode(&user)
	pass := md5.New()
	io.WriteString(pass, user.Pass)
	passHash := pass.Sum(nil)

	auth.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		resp := b.Get([]byte(user.Name))
		if resp != nil {
			if bytes.Equal(resp, passHash) {
				t := make([]byte, 16)
				rand.Read(t)
				tt.Name = user.Name
				tt.Token = fmt.Sprintf("%X", t[0:16])
				tokens[tt.Token] = user.Name
			} else {
				err.Code = 500
				err.Text = "Wrong password"
				str, _ := json.Marshal(err)
				http.Error(w, string(str), 500)
				return nil
			}
		} else {
			err.Code = 500
			err.Text = "User not found"
			str, _ := json.Marshal(err)
			http.Error(w, string(str), 500)
		}
		return nil
	})
	if err.Text == "" {
		json.NewEncoder(w).Encode(tt)
	}
}

// InitDb initialises shoppinglist db (boltDb)
func InitDb() DBase {
	db, err := bolt.Open("list.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	return DBase{DB: db}
}

//InitLoginBase Initialises users DB. TODO - move to mongo
func InitLoginBase() DBase {
	db, err := bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("Users"))
		return nil
	})
	return DBase{DB: db}
}

func main() {
	tokens = make(map[string]string)
	auth = InitLoginBase()
	dB = InitDb()

	router := mux.NewRouter()
	router.HandleFunc("/productList", GetProductListEndpoint).Methods("GET")
	router.HandleFunc("/product", AddProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", ToggleProductEndpoint).Methods("PUT")
	router.HandleFunc("/product/{id}", DeleteProductEndpoint).Methods("DELETE")
	router.HandleFunc("/product/{id}", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/signin", SignInEndpoint).Methods("POST")
	router.HandleFunc("/signup", SignUpEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":1881", router))
}
