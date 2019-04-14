package cosmos

import "testing"

func TestStoredProcedure(t *testing.T) {
	client := getDummyClient()
	coll := client.Database("dbtest").Collection("colltest")
	sp := coll.StoredProcedure("myStoredProcedure")

	if sp.client.rType != "sprocs" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.rID != "dbs/dbtest/colls/colltest/sprocs/myStoredProcedure" {
		t.Errorf("%+v", sp.client)
	}

	if sp.client.path != "dbs/dbtest/colls/colltest/sprocs/myStoredProcedure" {
		t.Errorf("%+v", sp.client)
	}

	sps := coll.StoredProcedures()
	if sps.client.rType != "sprocs" {
		t.Errorf("Wrong rType %s", sps.client.rType)
	}

	if sps.client.rID != "dbs/dbtest/colls/colltest/sprocs" {
		t.Errorf("Wrong rID %s", sps.client.rID)
	}

	if sps.client.path != "dbs/dbtest/colls/colltest/sprocs" {
		t.Errorf("Wrong path %s", sps.client.path)
	}
}
