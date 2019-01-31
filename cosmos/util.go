package cosmos

type QueryParam struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type SqlQuerySpec struct {
	Query      string       `json:"query"`
	Parameters []QueryParam `json:"parameters"`
}

func NewQuery(query string, queryParams ...QueryParam) *SqlQuerySpec {
	return &SqlQuerySpec{Query: query, Parameters: queryParams}
}

type Coordinates [][2]float64

type LineString struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

func NewLineString() *LineString {
	line := &LineString{Type: "LineString", Coordinates: Coordinates{}}
	return line
}

func (l *LineString) AddPoint(coords [2]float64) {
	l.Coordinates = append(l.Coordinates, coords)
}
