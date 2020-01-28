package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"tara/dependencies"
	"tara/models"
	"tara/schemas"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetPost Use to get one Post in mongodb
func GetPost(ctx *gin.Context) {
	var post models.Posts
	docID := ctx.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(docID)

	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	if err := collection.FindOne(mongoCtx, bson.M{"_id": objectID}).Decode(&post); err != nil {
		log.Fatal(err)
	}

	if post.Published != false {
		postSchema := schemas.PostSchema{ID: post.ID, Title: post.Title, Content: post.Title, UpdatedAt: post.UpdatedAt, Published: post.Published}
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": postSchema})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "data": "Post Not Found."})
	}
}

func GetPosts(ctx *gin.Context) {
	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	// Find many posts
	findOptions := options.Find()
	findOptions.SetLimit(10)

	var postSchemas []schemas.PostSchema
	cur, err := collection.Find(mongoCtx, bson.M{"published": bson.M{"$eq": true}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(mongoCtx) {
		var element models.Posts
		if err = cur.Decode(&element); err != nil {
			log.Fatal(err)
		}
		fmt.Println(element)

		postSchemas = append(postSchemas, schemas.PostSchema{ID: element.ID, Title: element.Title, Content: element.Content, UserID: element.UserID, UpdatedAt: element.UpdatedAt, Published: element.Published})
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
	docID := ctx.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(docID)

	var post schemas.PostSchema
	ctx.BindJSON(&post)

	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	updateOptions := bson.M{"updated_at": time.Now()}
	if post.Title != "" {
		updateOptions["title"] = post.Title
	}
	if post.Content != "" {
		updateOptions["content"] = post.Content
	}
	if post.Published == true {
		updateOptions["published"] = true
	}
	update := bson.M{"$set": updateOptions}

	result, err := collection.UpdateOne(mongoCtx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

func DeletePost(ctx *gin.Context) {
	docID := ctx.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(docID)

	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := dependencies.Client.Database("demo").Collection("posts")

	result, err := collection.DeleteOne(mongoCtx, bson.M{"_id": objectID})
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}
