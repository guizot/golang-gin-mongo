package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guizot/go-gin-mongodb/config"
	model_user "github.com/guizot/go-gin-mongodb/src/models/user"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Static Collection
const UserCollection = "user"

// Get DB from Mongo Config
func MongoConfig() *mgo.Database {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

// Get All User Endpoint
func GetAllUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	users := model_user.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get All User",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": &users,
	})
}

// Get User Endpoint
func GetUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	user := model_user.User{}
	err := db.C(UserCollection).Find(bson.M{"id": &idParse}).One(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get User",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": &user,
	})
}

// Create User Endpoint
func CreateUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	user := model_user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = db.C(UserCollection).Insert(user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Insert User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Insert User",
		"user":    &user,
	})
}

// Update User Endpoint
func UpdateUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	user := model_user.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Get Body",
		})
		return
	}

	user.ID = idParse
	user.UpdatedAt = time.Now()

	err = db.C(UserCollection).Update(bson.M{"id": &idParse}, user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Update User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Update User",
		"user":    &user,
	})
}

// Delete User Endpoint
func DeleteUser(c *gin.Context) {
	db := *MongoConfig()
	fmt.Println("MONGO RUNNING: ", db)

	id := c.Param("id")                   // Get Param
	idParse, errParse := strconv.Atoi(id) // Convert String to Int
	if errParse != nil {
		c.JSON(200, gin.H{
			"message": "Error Parse Param",
		})
		return
	}

	err := db.C(UserCollection).Remove(bson.M{"id": &idParse})
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Error Delete User",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Succes Delete User",
	})
}
