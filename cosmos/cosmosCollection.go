package cosmos

// Collection performs operations on a given collection.
type Collection struct {
	client Client
	db     Database
	collID string
}

type Collections struct {
	client Client
	db     Database
}

func (c Collection) Document(docID string) *Document {
	return newDocument(c, docID)
}

func (c Collection) Documents() *Documents {
	return newDocuments(c)
}

func (c Collection) UDF(id string) *UDF {
	return newUDF(c, id)
}

func (c Collection) UDFs() *UDFs {
	return newUDFs(c)
}

func (c Collection) StoredProcedure(id string) *StoredProcedure {
	return newStoredProcedure(c, id)
}

func (c Collection) StoredProcedures() *StoredProcedures {
	return newStoredProcedures(c)
}

func (c Collection) Trigger(id string) *Trigger {
	return newTrigger(c, id)
}

func (c Collection) Triggers() *Triggers {
	return newTriggers(c)
}

func newCollection(db Database, collID string) *Collection {
	db.client.path += "/colls/" + collID
	db.client.rType = "colls"
	db.client.rLink = db.client.path
	coll := &Collection{
		client: db.client,
		db:     db,
		collID: collID,
	}

	return coll
}

func newCollections(db Database) *Collections {
	db.client.path += "/colls"
	db.client.rType = "colls"
	coll := &Collections{
		client: db.client,
		db:     db,
	}

	return coll
}

func (c *Collections) Create(newColl *CollectionDefinition) (*CollectionDefinition, error) {
	respColl := &CollectionDefinition{}
	_, err := c.client.create(newColl, respColl)

	if err != nil {
		return nil, err
	}

	return respColl, err
}

func (c *Collections) ReadAll() (*CollectionDefinitions, error) {
	data := struct {
		Collections CollectionDefinitions `json:"DocumentCollections,omitempty"`
		Count       int                   `json:"_count,omitempty"`
	}{}
	_, err := c.client.read(&data)
	return &data.Collections, err
}

func (c *Collection) Read() (*CollectionDefinition, error) {
	coll := &CollectionDefinition{}
	_, err := c.client.read(coll)
	return coll, err
}

func (c *Collection) Delete() (*Response, error) {
	return c.client.delete()
}

func (c *Collection) Replace(i *IndexingPolicy, ret interface{}, opts ...CallOption) (*Response, error) {
	body := struct {
		ID             string          `json:"id"`
		IndexingPolicy *IndexingPolicy `json:"indexingPolicy"`
	}{c.collID, i}
	return c.client.replace(&body, ret, opts...)
}
