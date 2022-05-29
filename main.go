package main

import (
	"errors"
)

type warning interface {
	Warning() bool
	error
}

func IsWarning(err error) bool {
	var w warning
	return errors.As(err, &w) && w.Warning()
}

type warningError struct {
	err error
}

func (w warningError) Error() string {
	return w.err.Error()
}

func (w warningError) Warning() bool {
	return true
}

func newWarning(err error) error {
	return &warningError{err}
}

// PostMessage tries to post a message to a given Slack channel (or
// user). It will return a warning for known issues (missing scope, archiving
// channel, auth issues) or an error for anything else. It should be used in
// favour of PostMessageContext in almost all contexts.
func PostMessage() (messageTs *string, err error) {
	messageOpts = append(messageOpts,
		slack.MsgOptionText(fallbackText, false))

	//lint:ignore SA1019 we want to allow this to be called only inside this helper
	_, messageTS, err := sc.PostMessageContext(
		ctx, channelID, messageOpts...)

	if err != nil {
		if slackerrors.IsChannelNotFoundErr(err) ||
			slackerrors.IsIsArchivedErr(err) ||
			slackerrors.IsAuthenticationError(err) ||
			slackerrors.IsMissingScopeErr(err) {
			return nil, newWarning(err)
		}
		return nil, err
	}
	return &messageTS, nil
}

func main() {
	result, err := PostMessage()
	if err != nil; {
		if IsWarning(err) {
			"do nothing"
		} else {
			"do something"
		}
	}
}
