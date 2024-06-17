package controllers

import (
	"context"
	"time"
	"log"
	"net/http"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)


var SECRET_KEY = []byte("gosecretkey")

type User struct{
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

var client *mongo.Client

func getHash(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}

func GenerateJWT()(string,error){
	token:= jwt.New(jwt.SigningMethodHS256)
	tokenString, err :=  token.SignedString(SECRET_KEY)
	if err !=nil{
		log.Println("Error in JWT token generation")
		return "",err
	}
	return tokenString, nil
}

func UserSignup(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	collection := client.Database("GODB").Collection("user")
	ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx,user)
	json.NewEncoder(response).Encode(result)
}


func UserLogin(response http.ResponseWriter, request *http.Request){
  response.Header().Set("Content-Type","application/json")
  var user User
  var dbUser User
  json.NewDecoder(request.Body).Decode(&user)
  collection:= client.Database("GODB").Collection("user")
  ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
  err:= collection.FindOne(ctx, bson.M{"email":user.Email}).Decode(&dbUser)

  if err!=nil{
	  response.WriteHeader(http.StatusInternalServerError)
	  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	  return
  }
  userPass:= []byte(user.Password)
  dbPass:= []byte(dbUser.Password)

  passErr:= bcrypt.CompareHashAndPassword(dbPass, userPass)

  if passErr != nil{
	  log.Println(passErr)
	  response.Write([]byte(`{"response":"Wrong Password!"}`))
	  return
  }
  jwtToken, err := GenerateJWT()
  if err != nil{
	response.WriteHeader(http.StatusInternalServerError)
	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	return
  }
  response.Write([]byte(`{"token":"`+jwtToken+`"}`))
  
}