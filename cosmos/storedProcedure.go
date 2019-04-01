package cosmos

type StoredProcedure struct {
	client            Client
	coll              Collection
	storedProcedureID string
}

type StoredProcedures struct {
	client Client
	coll   *Collection
}

func newStoredProcedure(coll Collection, storedProcedureID string) *StoredProcedure {
	coll.client.fullPath = coll.client.fullPath + "/sprocs/" + storedProcedureID
	coll.client.postFix = coll.client.postFix + "/sprocs/" + storedProcedureID
	coll.client.rType = "sprocs"
	coll.client.rID = coll.client.postFix
	udf := &StoredProcedure{
		client:            coll.client,
		coll:              coll,
		storedProcedureID: storedProcedureID,
	}

	return udf
}

func newStoredProcedures(coll *Collection) *StoredProcedures {
	coll.client.fullPath = coll.client.fullPath + "/sprocs"
	coll.client.rType = "sprocs"
	coll.client.rID = coll.client.postFix
	udfs := &StoredProcedures{
		client: coll.client,
		coll:   coll,
	}

	return udfs
}

func (s *StoredProcedures) Create(newStoredProcedure *StoredProcedureDefinition, opts ...CallOption) (*StoredProcedureDefinition, error) {
	storedProcedureResp := &StoredProcedureDefinition{}
	_, err := s.client.create(newStoredProcedure, &storedProcedureResp, opts...)
	if err != nil {
		return nil, err
	}

	return storedProcedureResp, err
}

func (s *StoredProcedure) Replace(newStoredProcedure *StoredProcedureDefinition, opts ...CallOption) (*StoredProcedureDefinition, error) {
	storedProcedureResp := &StoredProcedureDefinition{}

	_, err := s.client.create(newStoredProcedure, &storedProcedureResp, opts...)

	if err != nil {
		return nil, err
	}

	return storedProcedureResp, err
}

func (s *StoredProcedures) ReadAll(opts ...CallOption) ([]StoredProcedureDefinition, error) {
	data := struct {
		StoredProcedures []StoredProcedureDefinition `json:"StoredProcedures,omitempty"`
		Count            int                         `json:"_count,omitempty"`
	}{}

	_, err := s.client.read(&data, opts...)

	if err != nil {
		return nil, err
	}
	return data.StoredProcedures, err
}

func (s *StoredProcedure) Delete(opts ...CallOption) (*Response, error) {
	return s.client.delete(opts...)
}

// Execute stored procedure
func (s *StoredProcedure) Execute(params, body interface{}, opts ...CallOption) error {
	return s.client.execute(params, &body, opts...)
}
