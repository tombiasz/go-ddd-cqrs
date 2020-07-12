package domain

import (
	"bytes"
	"fmt"
)

type DomainError struct {
	Property string
	Message  string
}

type DomainErrors []*DomainError

func (e DomainError) Error() string {
	return fmt.Sprintf("DomainError - %s - %s", e.Property, e.Message)
}

func NewDomainError(property, message string) *DomainError {
	return &DomainError{property, message}
}

func CombineDomainErrors(errs ...*DomainError) DomainErrors {
	var des DomainErrors = nil

	for _, err := range errs {
		if err != nil {
			des = append(des, err)
		}
	}

	return des
}

func (errs DomainErrors) Error() string {
	var buffer bytes.Buffer

	for _, e := range errs {
		buffer.WriteString(fmt.Sprintf("%s\n", e.Error()))
	}

	return buffer.String()
}
