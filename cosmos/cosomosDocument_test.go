package cosmos

import "testing"

func TestCosmosDocument(t *testing.T) {
	client := getDummyClient()
	coll := client.Database("dbtest").Collection("colltest")
	doc := coll.Document("doctest")

	if doc.client.rType != "docs" {
		t.Errorf("%+v", doc.client)
	}

	if doc.client.rID != "dbs/dbtest/colls/colltest/docs/doctest" {
		t.Errorf("%+v", doc.client)
	}

	if doc.client.path != "dbs/dbtest/colls/colltest/docs/doctest" {
		t.Errorf("%+v", doc.client)
	}

	docs := coll.Documents()
	if docs.client.rType != "docs" {
		t.Errorf("Wrong rType %s", docs.client.rType)
	}

	if docs.client.rID != "dbs/dbtest/colls/colltest/docs" {
		t.Errorf("Wrong rID %s", docs.client.rID)
	}

	if docs.client.path != "dbs/dbtest/colls/colltest/docs" {
		t.Errorf("Wrong path %s", docs.client.path)
	}
}
