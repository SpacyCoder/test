package cosmos

import (
	"encoding/json"
	"io"
)

func (q SqlQuerySpec) Read(p []byte) (n int, err error) {
	b, err := json.Marshal(q)
	copy(p, b)
	return len(b), io.EOF
}

type Document struct {
	client Client
	coll   Collection
	docID  string
}

type Documents struct {
	client Client
	coll   Collection
}

type DocumentData struct {
	ID          string `json:"id"`
	Rid         string `json:"_rid"`
	Self        string `json:"_self"`
	Etag        string `json:"_etag"`
	Ts          int    `json:"_ts"`
	Attachments string `json:"_attachments"`
}

type ListCosmosDocument struct {
	Rid       string      `json:"_rid"`
	Documents interface{} `json:"Documents"`
	Count     int         `json:"_count"`
}

func newDocument(coll Collection, docID string) *Document {
	coll.client.path += "/docs/" + docID
	coll.client.rType = "docs"
	coll.client.rLink = coll.client.path
	doc := &Document{
		client: coll.client,
		coll:   coll,
		docID:  docID,
	}

	return doc
}

func newDocuments(coll Collection) *Documents {
	coll.client.path += "/docs"
	coll.client.rType = "docs"
	docs := &Documents{
		client: coll.client,
		coll:   coll,
	}

	return docs
}

func (d *Documents) Create(doc interface{}, opts ...CallOption) (*Response, error) {
	d.client.createIDIfNotSet(doc)
	return d.client.create(doc, &doc, opts...)
}

func (d *Documents) ReadAll(docs interface{}, opts ...CallOption) (*Response, error) {
	data := struct {
		Documents interface{} `json:"Documents,omitempty"`
		Count     int         `json:"_count,omitempty"`
	}{Documents: docs}
	res, err := d.client.read(&data, opts...)
	return res, err
}

func (doc Document) Read(ret interface{}, opts ...CallOption) (*Response, error) {
	return doc.client.read(ret, opts...)
}

func (d Documents) Query(query *SqlQuerySpec, docs interface{}, opts ...CallOption) (*Response, error) {
	data := struct {
		Documents interface{} `json:"Documents,omitempty"`
		Count     int         `json:"_count,omitempty"`
	}{Documents: docs}
	res, err := d.client.query(query, &data, opts...)
	return res, err
}

func (d *Document) Replace(doc interface{}, opts ...CallOption) (*Response, error) {
	d.client.createIDIfNotSet(doc)
	return d.client.replace(doc, &doc, opts...)
}

func (d Document) Delete(opts ...CallOption) (*Response, error) {
	return d.client.delete(opts...)
}
