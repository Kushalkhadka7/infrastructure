package router

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"manager/config"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Router is a app level router.
type Router struct {
	*gin.Engine
}

type response struct {
	hostname string
}

type Post struct {
	_id  string `bson:"title,omitempty"`
	name string `bson:"body,omitempty"`
}

// New initializes new gin router.
func New(config *config.Config) *Router {
	gin.ForceConsoleColor()
	r := gin.New()

	r.Use(gin.Recovery())

	r.GET("/info", func(c *gin.Context) {
		password, err := ioutil.ReadFile("/data/PASSWORD")
		if err != nil {
			fmt.Println(err)
		}

		userName, err := ioutil.ReadFile("/data/USER_NAME")
		if err != nil {
			fmt.Println(err)
		}

		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":          config.Server.Name,
				"HOST_IP_ADDR":  string(value),
				"host":          c.Request.Host,
				"USER_NAME":     string(userName),
				"PASSWORD":      string(password),
				"url":           c.Request.RequestURI,
				"ADMIN_URL":     config.Server.AdminUrl,
				"MANAGER_URL":   config.Server.AuthUrl,
				"USER_NAME_ENV": os.Getenv("USER_NAME_ENV"),
				"PASSWORD_ENV":  os.Getenv("PASSWORD_ENV"),
			},
		})
	})

	r.GET("/", func(c *gin.Context) {
		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":         config.Server.Name,
				"host":         c.Request.Host,
				"url":          c.Request.RequestURI,
				"HOST_IP_ADDR": string(value),
			},
		})

	})

	r.GET("/auth", func(c *gin.Context) {
		cmd := exec.Command("hostname", "-i")
		value, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}

		resp, err := http.Get("http://dev-auth-service:4000/info")
		if err != nil {
			fmt.Println(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		var target interface{}

		err = json.Unmarshal(body, &target)
		if err != nil {
			panic(err.Error())
		}

		c.JSON(200, gin.H{
			"message": "Success",
			"data": map[string]interface{}{
				"name":         config.Server.Name,
				"host":         c.Request.Host,
				"url":          c.Request.RequestURI,
				"HOST_IP_ADDR": string(value),
			},
			"hostInfo": target,
		})
	})

	r.GET("/data", func(c *gin.Context) {
		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://auth:auth@example-mongodb-svc.dev-mongo:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false")

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			log.Fatal(err)
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected to MongoDB!")

		collection := client.Database("auth").Collection("sample")

		cur, err := collection.Find(context.TODO(), bson.D{{}})
		if err != nil {
			log.Fatal(err)
		}

		var results []*Post
		for cur.Next(context.TODO()) {

			// create a value into which the single document can be decoded
			var elem Post
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}

			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(results)
		// Close the cursor once finished
		cur.Close(context.TODO())

		fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

		c.JSON(200, gin.H{
			"message": "Success",
			"data":    &results,
		})
	})

	return &Router{r}
}
