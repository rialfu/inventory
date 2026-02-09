package utils

import "strings"

type MultiError struct {
	Errors []error
}

func (m *MultiError) Error() string {
	if len(m.Errors) == 0 {
		return ""
	}
	msgs := make([]string, 0, len(m.Errors))
	for _, err := range m.Errors {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}
func (m *MultiError) ErrorSlice() []string {
	msgs := make([]string, 0, len(m.Errors))
	for _, err := range m.Errors {
		msgs = append(msgs, err.Error())
	}
	return msgs
}

func (m *MultiError) Add(err error) {
	if err != nil {
		m.Errors = append(m.Errors, err)
	}
}

func (m *MultiError) HasError() bool {
	return len(m.Errors) > 0
}
