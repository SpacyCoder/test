package cosmos

// Resource every document have these
type Resource struct {
	ID   string `json:"id,omitempty"`
	Self string `json:"_self,omitempty"`
	Etag string `json:"_etag,omitempty"`
	Rid  string `json:"_rid,omitempty"`
	Ts   int    `json:"_ts,omitempty"`
}

// Indexing policy
// TODO: Ex/IncludePaths
type IndexingPolicy struct {
	IndexingMode string `json: "indexingMode,omitempty"`
	Automatic    bool   `json: "automatic,omitempty"`
}

// DatabaseDefinition defines the structure of a database data query
type DatabaseDefinition struct {
	Resource
	Colls string `json:"_colls,omitempty"`
	Users string `json:"_users,omitempty"`
}

// DatabaseDefinitions slice of Database elements
type DatabaseDefinitions []DatabaseDefinition

// First returns first database in slice
func (d DatabaseDefinitions) First() *DatabaseDefinition {
	if len(d) == 0 {
		return nil
	}
	return &d[0]
}

// CollectionDefinition defiens the structure of a Collection
type CollectionDefinition struct {
	Resource
	IndexingPolicy IndexingPolicy         `json:"indexingPolicy,omitempty"`
	PartitionKey   PartitionKeyDefinition `json:"partitionKey"`
	Docs           string                 `json:"_docs,omitempty"`
	Udf            string                 `json:"_udfs,omitempty"`
	Sporcs         string                 `json:"_sporcs,omitempty"`
	Triggers       string                 `json:"_triggers,omitempty"`
	Conflicts      string                 `json:"_conflicts,omitempty"`
}

type PartitionKeyDefinition struct {
	Paths []string `json:"paths"`
	Kind  string   `json:"kind"`
}

// Collections slice of Collection elements
type CollectionDefinitions []CollectionDefinition

type CollectionsResponse struct {
	Collections CollectionDefinition `json:"Collections,omitempty"`
	Count       int                  `json:"_count,omitempty"`
}

// First returns first database in slice
func (c CollectionDefinitions) First() *CollectionDefinition {
	if len(c) == 0 {
		return nil
	}
	return &c[0]
}

// DocumentDefinition is the struct of a document
type DocumentDefinition struct {
	Resource
	Attachments string `json:"attachments,omitempty"`
}

// StoredProcedure structure
type StoredProcedureDefinition struct {
	Resource
	Body string `json:"body,omitempty"`
}

// UDF (User Defined Function) definition
type UDFDefinition struct {
	Resource
	Body string `json:"body,omitempty"`
}

// PartitionKeyRange partition key range model
type PartitionKeyRange struct {
	Resource
	PartitionKeyRangeID string `json:"id,omitempty"`
	MinInclusive        string `json:"minInclusive,omitempty"`
	MaxInclusive        string `json:"maxExclusive,omitempty"`
}

type Trigger struct {
	Resource
	Body             string `json:"body"`
	TriggerOperation string `json:"triggerOperation"`
	TriggerType      string `json:"triggerType"`
}

type MsToken struct {
	Token string `json:"token"`
	Range struct {
		Min string `json:"min"`
		Max string `json:"max"`
	} `json:"range"`
}
