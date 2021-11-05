package mocks

import "github.com/mxschmitt/playwright-go"

type JSHandle struct {
	Value    string
	Children map[string]JSHandle
}

func (j JSHandle) AsElement() playwright.ElementHandle {
	return nil
}

func (j JSHandle) Dispose() error {
	return nil
}

func (j JSHandle) Evaluate(expression string, options ...interface{}) (interface{}, error) {
	return nil, nil
}

func (j JSHandle) EvaluateHandle(expression string, options ...interface{}) (playwright.JSHandle, error) {
	return &JSHandle{}, nil
}

func (j JSHandle) GetProperties() (map[string]playwright.JSHandle, error) {
	return nil, nil
}

func (j JSHandle) GetProperty(name string) (playwright.JSHandle, error) {
	return j.Children[name], nil
}

func (j JSHandle) JSONValue() (interface{}, error) {
	return nil, nil
}

func (j JSHandle) String() string {
	return j.Value
}
