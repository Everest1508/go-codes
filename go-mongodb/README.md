# MongoDB with Golang

## 1. Introduction to MongoDB

### 1.1 What is MongoDB?
MongoDB is a NoSQL database that stores data in a flexible, JSON-like format, known as BSON (Binary JSON). In MongoDB, schemas don't have to be specified upfront; they can differ every time. The key differences between SQL and NoSQL databases include:
- Schema Flexibility
- Data Storage: SQL databases store data in tables, and MongoDB stores data as documents in collections.
- Scalability

### 1.2 Key Features of MongoDB
- Document-Oriented
- Fast
- Indexing
- Replication

## 2. MongoDB with Golang

### 2.1 Connecting MongoDB with Go
To connect MongoDB with Go, you'll need to use the official MongoDB Go driver. MongoDB also has brief documentation for connecting MongoDB with Golang. Here is a quick guide:

1. Create a Go module.
2. Add the driver: `go get go.mongodb.org/mongo-driver/mongo`.
3. To fetch .env variables, add this dependency: `go get github.com/joho/godotenv`.
4. Create a database in MongoDB.
5. Create a `.env` file in the Go module and add `MONGODB_URI` variable with the connection string. For local MongoDB, the string will be `mongodb://localhost:27017`.

Now, the setup for MongoDB is complete.

### 2.1 Context Package
The MongoDB Go Driver uses the context package from Go's standard library to allow applications to signal timeouts and cancellations for any blocking method call. A blocking method relies on an external event, such as network input or output, to proceed with its task. An example of a blocking method in the Go Driver is the `insert()` method. If you want to perform an insert operation on a collection within 10 seconds, you can use a Context with a timeout. If the operation doesn't complete within the timeout, the method returns an error.

### Example: MongoDB and Golang Connection

```go
package main

// imports
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// main functions
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
```
