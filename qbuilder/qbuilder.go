package qbuilder

import "github.com/SpacyCoder/test/cosmos"

type Condition struct {
	ConditionType string
	Value         string
}

type QueryBuilder struct {
	S          []string
	F          string
	Conditions []Condition
	Parameters []cosmos.QueryParam
}

// New creates a new QueryBuilder
func New() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(s ...string) *QueryBuilder {
	qb.S = s
	return qb
}

func (qb *QueryBuilder) From(f string) *QueryBuilder {
	qb.F = f
	return qb
}

func (qb *QueryBuilder) And(conditionsValues ...string) *QueryBuilder {
	var conditions []Condition
	for _, c := range conditionsValues {
		condition := Condition{ConditionType: "AND", Value: c}
		conditions = append(conditions, condition)
	}
	qb.Conditions = append(qb.Conditions, conditions...)
	return qb
}

func (qb *QueryBuilder) Or(conditionsValues ...string) *QueryBuilder {
	var conditions []Condition
	for _, c := range conditionsValues {
		condition := Condition{ConditionType: "OR", Value: c}
		conditions = append(conditions, condition)
	}
	qb.Conditions = append(qb.Conditions, conditions...)
	return qb
}

func (qb *QueryBuilder) Params(params ...cosmos.QueryParam) *QueryBuilder {
	qb.Parameters = append(qb.Parameters, params...)
	return qb
}

func (qb *QueryBuilder) Build() *cosmos.SqlQuerySpec {
	query := "SELECT "
	for i, s := range qb.S {
		if i == 0 {
			query += s
		} else {
			query += ", " + s
		}
	}

	query += " FROM " + qb.F
	if len(qb.Conditions) == 0 {
		return cosmos.Q(query)
	}

	query += " WHERE"
	for i, c := range qb.Conditions {
		if i == 0 {
			query += " " + c.Value
		} else {
			query += " " + c.ConditionType + " " + c.Value
		}
	}

	return cosmos.Q(query, qb.Parameters...)
}
