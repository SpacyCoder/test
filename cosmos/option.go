package cosmos

import "encoding/json"

// Consistency type to define consistency levels
type Consistency string

const (
	// Strong consistency level
	Strong Consistency = "strong"

	// Bounded consistency level
	Bounded Consistency = "bounded"

	// Session consistency level
	Session Consistency = "session"

	// Eventual consistency level
	Eventual Consistency = "eventual"
)

type CallOption func(r *Request) error

func PartitionKey(partitionKey interface{}) CallOption {

	var (
		pk  []byte
		err error
	)
	switch v := partitionKey.(type) {
	case json.Marshaler:
		pk, err = Serialization.Marshal(v)
	default:
		pk, err = Serialization.Marshal([]interface{}{v})
	}

	header := []string{string(pk)}

	return func(r *Request) error {
		if err != nil {
			return err
		}
		r.Header[HeaderPartitionKey] = header
		return nil
	}
}

// Upsert if set to true, Cosmos DB creates the document with the ID (and partition key value if applicable) if it doesn’t exist, or update the document if it exists.
func Upsert() CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderUpsert, "true")
		return nil
	}
}

// CrossPartition allows query to run on all partitions
func CrossPartition() CallOption {
	return func(r *Request) error {
		r.Header.Set(HeaderCrossPartition, "true")
		return nil
	}
}
