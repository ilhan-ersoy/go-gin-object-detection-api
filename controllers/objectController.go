package controllers

import (
	"context"
	"fmt"
	"go-gin-object-detection-api/database"
	helper "go-gin-object-detection-api/helpers"
	"go-gin-object-detection-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var objectsCollection *mongo.Collection = database.OpenCollection(database.Client, "objects")

func CreateObject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var object models.Object

		err := c.BindJSON(&object)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		existingObject := models.Object{}
		filter := bson.M{"image": object.Image}
		err = objectsCollection.FindOne(ctx, filter).Decode(&existingObject)

		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("A data with the same image already exists: %s", existingObject.ID.Hex()),
			})
			return
		} else if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		newObject := models.Object{
			ID:         primitive.NewObjectID(),
			Name:       object.Name,
			User_id:    object.User_id,
			Image:      object.Image,
			Labels:     helper.RemoveDuplicateStr(object.Labels),
			Created_at: time.Time{},
			Updated_at: time.Time{},
		}

		result, err := objectsCollection.InsertOne(ctx, newObject)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": result.InsertedID,
		})
	}
}

func GetUserObjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var objects []models.Object
		defer cancel()

		userId := c.Param("user_id")

		filter := bson.D{{"user_id", userId}}

		cursor, err := objectsCollection.Find(ctx, filter)

		if err != nil {
			panic(err)
		}

		if err = cursor.All(context.TODO(), &objects); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": objects,
		})
	}
}

func GetObject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var object models.Object
		defer cancel()

		itemID := c.Param("item_id")

		objID, _ := primitive.ObjectIDFromHex(itemID)

		err := objectsCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&object)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"data": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, object)

	}
}

func DeleteObject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		itemID := c.Param("item_id")

		objID, _ := primitive.ObjectIDFromHex(itemID)

		err := objectsCollection.FindOneAndDelete(ctx, bson.M{"_id": objID})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": err.Err(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})

	}
}

func DeleteObjects() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := c.Param("user_id")

		filter := bson.D{{"user_id", userId}}

		_, err := objectsCollection.DeleteMany(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"data": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": "ok",
		})
	}
}
