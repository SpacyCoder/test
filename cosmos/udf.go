package cosmos

type UDF struct {
	client Client
	coll   *Collection
	udfID  string
}

type UDFs struct {
	client Client
	coll   *Collection
}

func newUDF(coll *Collection, udfID string) *UDF {
	coll.client.fullPath = coll.client.fullPath + "/udfs/" + udfID
	coll.client.postFix = coll.client.postFix + "/udfs/" + udfID
	coll.client.rType = "udfs"
	coll.client.rID = coll.client.postFix
	udf := &UDF{
		client: coll.client,
		coll:   coll,
		udfID:  udfID,
	}

	return udf
}

func newUDFs(coll *Collection) *UDFs {
	coll.client.fullPath = coll.client.fullPath + "/udfs"
	coll.client.rType = "udfs"
	coll.client.rID = coll.client.postFix
	udfs := &UDFs{
		client: coll.client,
		coll:   coll,
	}

	return udfs
}

func (u *UDF) Create(newUDF *UDFDefinition, opts ...CallOption) (*UDFDefinition, error) {
	createdUDF := &UDFDefinition{}

	_, err := u.client.create(newUDF, &createdUDF, opts...)

	if err != nil {
		return nil, err
	}

	return createdUDF, err
}

func (u *UDF) Replace(newUDF *UDFDefinition, opts ...CallOption) (*UDFDefinition, error) {
	createdUDF := &UDFDefinition{}

	_, err := u.client.create(newUDF, &createdUDF, opts...)

	if err != nil {
		return nil, err
	}

	return createdUDF, err
}

func (u *UDFs) ReadAll(opts ...CallOption) ([]UDFDefinition, error) {
	data := struct {
		Udfs  []UDFDefinition `json:"UserDefinedFunctions,omitempty"`
		Count int             `json:"_count,omitempty"`
	}{}

	_, err := u.client.read(&data, opts...)

	if err != nil {
		return nil, err
	}
	return data.Udfs, err
}

func (u *UDF) Delete(opts ...CallOption) (*Response, error) {
	return u.client.delete(opts...)
}
