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

	bolt "github.com/coreos/bbolt"
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
}

type User struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Token struct {
	Name  string
	Token string
}

var Products []Product
var DB DBase
var Auth DBase
var Tokens []Token

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
	json.NewEncoder(w).Encode(Products)

}

func GenerateGuid() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// AddProductEndpoint used for creating new equation in memory and getting result
func AddProductEndpoint(w http.ResponseWriter, req *http.Request) {
	var pr Product
	json.NewDecoder(req.Body).Decode(&pr)

	pr.ID = GenerateGuid()

	Logger("Input name: "+pr.Name, "log.txt")
	Products = append(Products, pr)
	DB.DB.Update(func(tx *bolt.Tx) error {
		prods, _ := tx.CreateBucketIfNotExists([]byte("NewList"))
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
	for index, item := range Products {
		if item.ID == params["id"] {
			Products = append(Products[:index], Products[index+1:]...)
			break
		}
	}
}

// GetProductEndpoint used for deleting old product by ID
func GetProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	DB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("NewList"))
		resp := b.Get([]byte(params["id"]))
		log.Println(string(resp))
		json.NewEncoder(w).Encode(string(resp))
		return nil
	})
}

// GetProductEndpoint used for deleting old product by ID
func EditProductEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	DB.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("NewList"))
		resp := b.Get([]byte(params["id"]))
		json.NewEncoder(w).Encode(string(resp))
		return nil
	})
}

// SignUpEndpoint used for IDK, like killing bytes?
func SignUpEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	pass := md5.New()
	io.WriteString(pass, user.Pass)
	passHash := pass.Sum(nil)
	Auth.DB.Update(func(tx *bolt.Tx) error {
		users, _ := tx.CreateBucketIfNotExists([]byte("Users"))
		users.Put([]byte(user.Name), passHash)
		return nil
	})
	//TODO: errors.
	json.NewEncoder(w).Encode(user)
}

// SignInEndpoint used for achieving auth token by user
func SignInEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User
	json.NewDecoder(req.Body).Decode(&user)
	pass := md5.New()
	io.WriteString(pass, user.Pass)
	passHash := pass.Sum(nil)
	var tt Token
	Auth.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Users"))
		resp := b.Get([]byte(user.Name))
		if resp != nil {
			if bytes.Equal(resp, passHash) {
				t := make([]byte, 16)
				rand.Read(t)
				tt.Name = user.Name
				tt.Token = fmt.Sprintf("%X", t[0:16])
				Tokens = append(Tokens, tt)
			}
		}
		return nil
	})
	//TODO: errors.
	json.NewEncoder(w).Encode(tt)
}

func InitDb() DBase {
	db, err := bolt.Open("list.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	return DBase{DB: db}
}

func InitLoginBase() DBase {
	db, err := bolt.Open("users.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	return DBase{DB: db}
}

func main() {
	file, _ := os.Create("log.txt")
	fmt.Fprint(file, "Log started at: "+time.Now().String()+"\n")
	defer file.Close()

	Auth = InitLoginBase()
	DB = InitDb()

	router := mux.NewRouter()
	router.HandleFunc("/productList", GetProductListEndpoint).Methods("GET")
	router.HandleFunc("/product", AddProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", EditProductEndpoint).Methods("PUT")
	router.HandleFunc("/product/{id}", DeleteProductEndpoint).Methods("DELETE")
	router.HandleFunc("/product/{id}", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/signin", SignInEndpoint).Methods("POST")
	router.HandleFunc("/signup", SignUpEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":1881", router))
}
