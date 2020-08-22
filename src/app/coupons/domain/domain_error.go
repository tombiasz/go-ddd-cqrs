package domain

import (
	"bytes"
	"fmt"
)

type DomainError struct {
	Code    string
	Message string
}

type DomainErrors []*DomainError

func (e *DomainError) Error() string {
	return fmt.Sprintf("DomainError - %s - %s", e.Code, e.Message)
}

func (e *DomainError) AsDomainErrors() DomainErrors {
	return CombineDomainErrors(e)
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

func (errs DomainErrors) AsMap() map[string]string {
	m := make(map[string]string)

	for _, e := range errs {
		m[e.Code] = e.Message
	}

	return m
}
