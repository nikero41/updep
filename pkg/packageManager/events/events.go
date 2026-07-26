package events

type PmOutputEvent struct {
	Output string
	Err    error
	Done   bool
}
