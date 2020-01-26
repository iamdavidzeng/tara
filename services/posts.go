package services

import (
	"context"
	"log"
	"net/http"
	"tara/dependencies"
	"tara/models"
	"tara/schemas"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPost Use to get one Post in mongodb
func GetPost(ctx *gin.Context) {
	return
}

func GetPosts(ctx *gin.Context) {
	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	// Find many posts
	findOptions := options.Find()
	findOptions.SetLimit(10)

	var postSchemas []schemas.PostSchema
	cur, err := collection.Find(mongoCtx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(mongoCtx) {
		var element models.Posts
		if err = cur.Decode(&element); err != nil {
			log.Fatal(err)
		}

		postSchemas = append(postSchemas, schemas.PostSchema{Title: element.Title, Content: element.Content, UserID: element.UserID, UpdatedAt: element.UpdatedAt})
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(mongoCtx)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": postSchemas})
}

func CreatePost(ctx *gin.Context) {
	var post schemas.PostSchema
	ctx.BindJSON(&post)

	initialPost := models.Posts{Title: post.Title, Content: post.Content, UserID: post.UserID, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	// Insert post to mongodb
	if _, err := collection.InsertOne(mongoCtx, initialPost); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Create successfully!"})
}

func UpdatePost(ctx *gin.Context) {
	return
}

func DeletePost(ctx *gin.Context) {
	return
}
