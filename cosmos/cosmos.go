package cosmos

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/google/uuid"
)

// Client model
type Client struct {
	key        string
	domain     string
	url        string
	path       string
	httpClient *http.Client
	rType      string // dbs,colls,docs,udfs,sprocs,triggers
	rLink      string
}

func (c *Client) getURL() string {
	return c.domain + c.path
}

var buffers = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// New create a new CosmosDB instance
func New(connString string) (*Client, error) {
	if connString == "" {
		return nil, errors.New("Invalid connection string")
	}
	array := strings.Split(connString, ";")
	path := strings.TrimPrefix(array[0], "AccountEndpoint=")
	if path == "" {
		return nil, errors.New("Invalid connection string")
	}
	key := strings.TrimPrefix(array[1], "AccountKey=")
	if key == "" {
		return nil, errors.New("Invalid connection string")
	}
	httpClient := &http.Client{}
	return &Client{key, path, path, "", httpClient, "", ""}, nil
}

// Offer defines all operation on a single offer
func (c Client) Offer(offerID string) *Offer {
	return newOffer(c, offerID)
}

// Offers defines all operation possible on multiple offers
func (c Client) Offers() *Offers {
	return newOffers(c)
}

// Database returns a new Database struct that contains the opertaions you can do on single database
func (c Client) Database(dbID string) *Database {
	return newDatabase(c, dbID)
}

// Databases returns a new Databases struct used to get data about various databases
func (c *Client) Databases() *Databases {
	return newDatabases(c)
}

func (c *Client) query(query *SqlQuerySpec, body interface{}, opts ...CallOption) (*Response, error) {
	buf := buffers.Get().(*bytes.Buffer)
	var err error
	buf.Reset()
	defer buffers.Put(buf)

	if err = Serialization.EncoderFactory(buf).Encode(query); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.getURL(), buf)
	if err != nil {
		return nil, err
	}

	r := ResourceRequest(c.rLink, c.rType, req)
	if err = c.apply(r, opts); err != nil {
		return nil, err
	}

	r.QueryHeaders(buf.Len())
	return c.do(r, expectStatusCode(http.StatusOK), body)
}

func (c *Client) read(ret interface{}, opts ...CallOption) (*Response, error) {
	buf := buffers.Get().(*bytes.Buffer)
	buf.Reset()
	res, err := c.method(http.MethodGet, expectStatusCode(http.StatusOK), ret, buf, opts...)

	buffers.Put(buf)
	return res, err
}

// Create resource
func (c *Client) create(body, ret interface{}, opts ...CallOption) (*Response, error) {
	data, err := stringify(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	return c.method(http.MethodPost, expectStatusCodeXX(http.StatusOK), ret, buf, opts...)
}

// Replace resource
func (c *Client) replace(body, ret interface{}, opts ...CallOption) (*Response, error) {
	data, err := stringify(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	return c.method(http.MethodPut, expectStatusCode(http.StatusOK), ret, buf, opts...)
}

// Delete resource
func (c *Client) delete(opts ...CallOption) (*Response, error) {
	return c.method(http.MethodDelete, expectStatusCode(http.StatusNoContent), nil, &bytes.Buffer{}, opts...)
}

func (c *Client) execute(body, ret interface{}, opts ...CallOption) (*Response, error) {
	data, err := stringify(body)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	return c.method(http.MethodPost, expectStatusCode(http.StatusOK), ret, buf, opts...)
}

func (c *Client) method(method string, validator statusCodeValidatorFunc, ret interface{}, data *bytes.Buffer, opts ...CallOption) (*Response, error) {
	req, err := http.NewRequest(method, c.getURL(), data)
	if err != nil {
		return nil, err
	}

	r := ResourceRequest(c.rLink, c.rType, req)
	// apply headers
	if err = c.apply(r, opts); err != nil {
		return nil, err
	}

	return c.do(r, validator, ret)
}

// do sends request to cosmos, validates the response and returns is.
func (c *Client) do(r *Request, validator statusCodeValidatorFunc, respBody interface{}) (*Response, error) {
	resp, err := c.httpClient.Do(r.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if response has expected status code.
	if !validator(resp.StatusCode) {
		var errorMessage CosmosErrorMessage
		readJSON(resp.Body, &errorMessage)
		return &Response{resp.Header}, NewCosmosError(&errorMessage, resp.StatusCode)
	}

	if respBody == nil {
		return nil, nil
	}
	return &Response{resp.Header}, readJSON(resp.Body, respBody)
}

// readJSON response to given interface(struct, map, ..)
func readJSON(reader io.Reader, data interface{}) error {
	return Serialization.DecoderFactory(reader).Decode(&data)
}

// apply sets default headers and all call options given.
func (c *Client) apply(r *Request, opts []CallOption) (err error) {
	if err = r.DefaultHeaders(c.key); err != nil {
		return err
	}

	for i := 0; i < len(opts); i++ {
		if err = opts[i](r); err != nil {
			return err
		}
	}
	return nil
}

// createIDIfNotSet create a uuid if the resource has not explicitly set an id.
func (c *Client) createIDIfNotSet(doc interface{}) {
	if reflect.TypeOf(doc).String() == "string" {
		return
	}
	id := reflect.ValueOf(doc).Elem().FieldByName("ID")
	if id.IsValid() && id.String() == "" {
		id.SetString(uuid.New().String())
	}
}

// stringify turns arbitrary body type to byte string.
func stringify(body interface{}) (bt []byte, err error) {
	switch t := body.(type) {
	case string:
		bt = []byte(t)
	case []byte:
		bt = t
	default:
		bt, err = Serialization.Marshal(t)
	}
	return
}
