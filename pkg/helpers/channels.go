package helpers

import tea "charm.land/bubbletea/v2"

type ChannelMsg[T any] struct {
	Result T
	Err    error
}

func WaitForChan[T any](c <-chan T) tea.Cmd {
	return func() tea.Msg {
		event := <-c

		return ChannelMsg[T]{
			Result: event,
			Err:    nil,
		}
	}
}

