package errors

import "fmt"

//NotImplementedError
type NotImplementedError struct {
	Method    string `json:"method,omitempty"`
	Interface string `json:"interface,omitempty"`
}

func (e NotImplementedError) Error() string {
	return fmt.Sprintf("not implemented: interface: %s, method: %s", e.Interface, e.Method)
}

//ResourceNotFoundError
type ResourceNotFoundError struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (e ResourceNotFoundError) Error() string {
	return fmt.Sprintf("resource not found: name: %s, id: %s", e.Name, e.ID)
}

//DuplicatedEntryError
type DuplicatedEntryError struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (e DuplicatedEntryError) Error() string {
	return fmt.Sprintf("duplicated entry: name: %s, id: %s", e.Name, e.ID)
}

//ForbiddenError
type ForbiddenError struct {
	Message string `json:"message,omitempty"`
}

func (e ForbiddenError) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Message)
}

//InvalidArgumentError
type InvalidArgumentError struct {
	Message string `json:"message,omitempty"`
}

func (e InvalidArgumentError) Error() string {
	return fmt.Sprintf("invalid argument: %s", e.Message)
}
