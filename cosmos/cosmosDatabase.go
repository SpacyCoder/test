package cosmos

type DatabaseData struct {
	ID    string `json:"id"`
	Rid   string `json:"_rid"`
	Ts    int    `json:"_ts"`
	Self  string `json:"_self"`
	Etag  string `json:"_etag"`
	Colls string `json:"_colls"`
	Users string `json:"_users"`
}

// ListDatabaseData is the struct of a Database Repsonse
type ListDatabaseData struct {
	Rid       string         `json:"_rid"`
	Count     int            `json:"_count"`
	Databases []DatabaseData `json:"Databases"`
}

type Database struct {
	client Client
	dbID   string
}

type Databases struct {
	client *Client
}

func newDatabase(client Client, dbID string) *Database {
	client.fullPath = client.basePath + "dbs/" + dbID
	client.postFix = "dbs/" + dbID
	client.rID = "dbs/" + dbID
	client.rType = "dbs"
	db := &Database{
		client: client,
		dbID:   dbID,
	}

	return db
}

func newDatabases(c *Client) *Databases {
	c.fullPath = c.basePath + "dbs"
	c.postFix = "dbs"
	c.rType = "dbs"
	c.rID = ""
	dbs := &Databases{
		client: c,
	}

	return dbs
}

func (db Database) Collection(collID string) *Collection {
	return newCollection(db, collID)
}

func (db Database) Collections() *Collections {
	return newCollections(db)
}

// Create a new database
func (db Databases) Create(body interface{}, opts ...CallOption) (*DatabaseDefinition, error) {
	dbDef := &DatabaseDefinition{}
	_, err := db.client.create(body, dbDef, opts...)
	if err != nil {
		return nil, err
	}

	return dbDef, err
}

// ReadAll databases
func (db *Databases) ReadAll(opts ...CallOption) (*DatabaseDefinitions, error) {
	data := struct {
		Databases DatabaseDefinitions `json:"Databases,omitempty"`
		Count     int                 `json:"_count,omitempty"`
	}{}
	_, err := db.client.read(&data, opts...)

	if err != nil {
		return nil, err
	}

	return &data.Databases, err
}

func (db *Database) Read() (*DatabaseDefinition, error) {
	ret := &DatabaseDefinition{}
	_, err := db.client.read(ret)
	return ret, err
}

func (db *Database) Delete() (*Response, error) {
	return db.client.delete()
}

func (db *Databases) Query(query *SqlQuerySpec, opts ...CallOption) (*DatabaseDefinitions, error) {
	data := struct {
		Databases DatabaseDefinitions `json:"Databases,omitempty"`
		Count     int                 `json:"_count,omitempty"`
	}{}

	_, err := db.client.query(query, &data, opts...)

	if err != nil {
		return nil, err
	}

	return &data.Databases, err
}
