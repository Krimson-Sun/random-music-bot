package events

type Event struct {
	Type    Type
	file_id string
	Meta    interface{}
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event) error
}
