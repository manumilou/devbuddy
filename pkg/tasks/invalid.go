package tasks

import (
	"fmt"
)

type Invalid struct {
	definition interface{}
	err        error
}

func newInvalid(definition interface{}, err error) Task {
	return &Invalid{definition: definition, err: err}
}

func (u *Invalid) name() string {
	return "Invalid task"
}

func (u *Invalid) header() string {
	return ""
}

func (u *Invalid) perform(ctx *context) (err error) {
	ctx.ui.TaskWarning(fmt.Sprintf("%s: %+v", u.err, u.definition))
	return nil
}

func (u *Invalid) actions(ctx *context) (actions []taskAction) {
	return
}
