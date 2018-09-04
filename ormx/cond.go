package ormx

import "fmt"

// Condition defines condition elements.
type Condition struct {
	Key   string      `json:"key"`
	Op    string      `json:"op"`
	Value interface{} `json:"value"`
}

// Conditions contains a list of condition and it raw query string.
type Conditions struct {
	Conds []Condition `json:"conditions"`
	args  []interface{}
	query string
}

// NewConditions return a new initialized conditions.
func NewConditions() *Conditions {
	return &Conditions{
		Conds: make([]Condition, 0),
	}
}

// Parse parses conditions args and raw query.
func (c Conditions) Parse() (string, []interface{}) {
	c.args = make([]interface{}, 0)
	for _, cond := range c.Conds {
		c.query += fmt.Sprintf("%s %s ? and ", cond.Key, cond.Op)
		c.args = append(c.args, cond.Value)
	}
	if c.query != "" {
		c.query = c.query[:len(c.query)-5]
	}

	return c.query, c.args
}
