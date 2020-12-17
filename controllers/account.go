package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	AccountController struct {
		session *mgo.Session
	}
)

func NewAccountController(s *mgo.Session) *AccountController {
	return &AccountController{s}
}

const (
    DB_NAME       = "management"
    DB_COLLECTION = "account"
)


func (uc AccountController) GetUsers(c *gin.Context) {
	
	var results []models.Account
    err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Find(nil).All(&results)
    if err != nil{
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "404",
		})
		return
	}
  
    c.JSON(200, results)
}


func (uc AccountController) GetOneUser(c *gin.Context) {

	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "No ID Found",
		})
		return
	}

	oid := bson.ObjectIdHex(id)
	u := models.Account{}
	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).FindId(oid).One(&u)

	if err != nil {
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "404",
		})
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json")
    c.JSON(200, u)
}


func (uc AccountController) CreateUser(c *gin.Context) {

    var json models.Account
    c.Bind(&json) 

	u := uc.create_user(json.Name, json.Score, json.Email,c)

    if u.Name == json.Name {
        content := gin.H{
            "result": "Success",
            "name": u.Name,
            "score": u.Score,
            "email": u.Email,
        }
    
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }
	
}

func (uc AccountController) create_user(Name string, Score float64,Email string, c *gin.Context) models.Account {

    user := models.Account{
		Name:	Name,
		Score:	Score,
    	Email:  Email,   
    }

	err := uc.session.DB(DB_NAME).C(DB_COLLECTION).Insert(&user)

    if err != nil {
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "Cannot connect into mongo",
		})
	}
    return user
}


func (uc AccountController) RemoveUser(c *gin.Context) {

	id := c.Params.ByName("id")

	if !bson.IsObjectIdHex(id){
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "No ID Found",
		})
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).RemoveId(oid); err != nil{
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "Failed to remove",
		})
		return
	}
	c.JSON(200, gin.H{
		"message" : "success",
	})
		
}

func (uc AccountController) UpdateUser(c *gin.Context) {

	id := c.Params.ByName("id")
    var json models.Account
    c.Bind(&json) 

	if !bson.IsObjectIdHex(id) {
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "ID is invalid",
		})
		return
	}

	u := uc.update_user(id,json.Name, json.Score,json.Email,c)

    if u.Name == json.Name {
        content := gin.H{
            "result": "Success",
			"name": u.Name,
			"score": u.Score,
            "email": u.Email,
        }
        c.Writer.Header().Set("Content-Type", "application/json")
        c.JSON(201, content)
    } else {
        c.JSON(500, gin.H{"result": "An error occured"})
    }
}


func (uc AccountController) update_user(Id string,Name string, Score float64,Email string, c *gin.Context) models.Account {

    user := models.Account{
        Name:      Name,
        Email:    Email,
        Score:	Score,
    }
    
	oid := bson.ObjectIdHex(Id)
	
    if err := uc.session.DB(DB_NAME).C(DB_COLLECTION).UpdateId(oid, &user); err != nil {
		c.JSON(422, gin.H{
			"error" : true,
			"message" : "Cannot connect into mongo",
		})
	}
    return user
}