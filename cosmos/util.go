package cosmos

func Q(query string, queryParams ...QueryParam) *SqlQuerySpec {
	return &SqlQuerySpec{Query: query, Parameters: queryParams}
}

type P = QueryParam

// Coordinate = [lon, lat]
type Coordinate [2]float64

type Coordinates []Coordinate

// LineString struct defines a line string
type LineString struct {
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

// NewLineString creates a new LineString struct.
func NewLineString(coords ...Coordinate) *LineString {
	line := &LineString{Type: "LineString", Coordinates: coords}
	return line
}

// AddPoint is a helper method for adding point to a LineString
func (l *LineString) AddPoint(lon, lat float64) {
	l.Coordinates = append(l.Coordinates, Coordinate{lon, lat})
}
