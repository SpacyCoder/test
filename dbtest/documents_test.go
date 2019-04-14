package dbtest

import (
	"os"
	"testing"
	"time"

	"github.com/SpacyCoder/test/cosmos"
)

func getClient() (*cosmos.Client, error) {
	testDbURL := os.Getenv("TEST_COSMOS_URL")
	return cosmos.New(testDbURL)
}

func TestCreateDocument(t *testing.T) {
	//TODO
}

var dbID = "db-test"
var collID = "coll-test"

type TestDoc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDatabaseOperations(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatalf("Creating client caused error: %s", err.Error())
	}

	// @TODO: currently responds with a 404 resource not found. But database is created correctly. Should try to find out why.
	// Same issue with delete.
	// Creating db
	/*
		newDbRef, err := client.Databases().Create(dbID)
		if err != nil {
			t.Fatalf("Creating database caused error: %s", err.Error())
		}
		if newDbRef.ID != dbID {
			t.Fatalf("Wrong ID: %s should be: test", newDbRef.ID)
		}*/

	// Reading db
	db := client.Database(dbID)
	testDb, err := db.Read()
	if err != nil {
		t.Fatalf("Reading database caused error: %s", err.Error())
	}
	if testDb.ID != dbID {
		t.Fatalf("Wrong ID: %s should be: %s", testDb.ID, dbID)
	}

	// List databases
	dbs := client.Databases()
	_, err = dbs.ReadAll()
	if err != nil {
		t.Fatalf("Listing databases caused error: %s", err.Error())
	}
}

func TestCollections(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatalf("Creating client caused error: %s", err.Error())
	}

	db := client.Database(dbID)
	colls := db.Collections()

	// Create collection
	/* 	_, err = colls.Create(cosmos.CollectionDefinition{Resource: cosmos.Resource{ID: collID}, PartitionKey: cosmos.PartitionKeyDefinition{Kind: "hash", Paths: []string{"/name"}}})
	   	if err != nil {
	   		t.Fatalf("Creating collection caused error: %s", err.Error())
	   	} */
	coll := db.Collection(collID)
	// Read collection
	collDef, err := coll.Read()
	if err != nil {
		t.Fatalf("Reading collection caused error: %s", err.Error())
	}
	if collDef.ID != collID {
		t.Fatalf("Wrong ID: %s should be: %s", collDef.ID, collID)
	}

	// List collections
	collDefs, err := colls.ReadAll()
	if err != nil {
		t.Fatalf("Listing collections caused error: %s", err.Error())
	}

	if len(*collDefs) != 1 {
		t.Fatalf("Number of collections are wrong: %d", len(*collDefs))
	}

	// Delete collection
	/* _, err = coll.Delete()
	if err != nil {
		t.Fatalf("Deleting collection caused error: %s", err.Error())
	}
	*/
}

func TestDocuments(t *testing.T) {
	client, err := getClient()
	if err != nil {
		t.Fatalf("Creating client caused error: %s", err.Error())
	}

	db := client.Database(dbID)
	coll := db.Collection(collID)
	docs := coll.Documents()

	user1 := &TestDoc{
		ID:   "user1",
		Name: "Lars",
		Age:  150,
	}

	// Create document
	_, err = docs.Create(user1, cosmos.PartitionKey(user1.Name))
	if err != nil {
		t.Fatalf("Creating collection caused error: %s", err.Error())
	}

	retUser := &TestDoc{}
	_, err = coll.Document("user1").Read(retUser, cosmos.PartitionKey("Lars"))
	if err != nil {
		t.Fatalf("Read collection caused error: %s", err.Error())
	}
	if retUser.Name != user1.Name {
		t.Fatalf("Wrong name: %s", retUser.Name)
	}

	user2 := &TestDoc{
		ID:   "user2",
		Name: "Trygve",
		Age:  150,
	}

	_, err = docs.Create(user2, cosmos.PartitionKey(user2.Name))

	/* 	coll := db.Collection(collID)

	   	// Read collection
	   	collDef, err := coll.Read()
	   	if err != nil {
	   		t.Fatalf("Reading collection caused error: %s", err.Error())
	   	}
	   	if collDef.ID != collID {
	   		t.Fatalf("Wrong ID: %s should be: %s", collDef.ID, collID)
	   	}

	   	// List collections
	   	collDefs, err := colls.ReadAll()
	   	if err != nil {
	   		t.Fatalf("Listing collections caused error: %s", err.Error())
	   	}

	   	if len(*collDefs) != 1 {
	   		t.Fatalf("Number of collections are wrong: %d", len(*collDefs))
	   	} */

	// Delete collection
	/* _, err = coll.Delete()
	if err != nil {
		t.Fatalf("Deleting collection caused error: %s", err.Error())
	} */

	users := []TestDoc{}
	_, err = docs.ReadAll(&users)
	if err != nil {
		t.Fatalf("Listing docs caused error: %s", err.Error())
	}
	if len(users) != 2 {
		t.Fatalf("Should be 2 but is: %d", len(users))
	}
	time.Sleep(1)
	qUsers := []TestDoc{}
	query := cosmos.Q("SELECT * FROM root WHERE root.name = @NAME", cosmos.P{Name: "@NAME", Value: user1.Name})
	_, err = docs.Query(query, &qUsers, cosmos.CrossPartition())
	if err != nil {
		t.Fatalf("Querying docs caused error: %s", err.Error())
	}
	if len(qUsers) != 1 {
		t.Fatalf("Should be 1 but is: %d", len(qUsers))
	}

	_, err = coll.Document("user1").Delete(cosmos.PartitionKey(user1.Name))
	if err != nil {
		t.Fatalf("Deleting user1 caused error: %s", err.Error())
	}
	_, err = coll.Document("user2").Delete(cosmos.PartitionKey(user2.Name))
	if err != nil {
		t.Fatalf("Deleting user2 caused error: %s", err.Error())
	}
}
