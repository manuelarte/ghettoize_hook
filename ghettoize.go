package ghuettoize

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
)

var _ logrus.Hook = new(ghettoize)

func NewGhettoize(options ...GhettoizeOption) (logrus.Hook, error) {
	oc := ollamaclient.New(usermodel.GetTextGenerationModel())
	for _, option := range options {
		option(oc)
	}
	if err := oc.PullIfNeeded(); err != nil {
		return nil, err
	}
	return ghettoize{oc: oc}, nil
}

type GhettoizeOption func(oc *ollamaclient.Config)

type ghettoize struct {
	oc *ollamaclient.Config
}

func (g ghettoize) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel}
}

func (g ghettoize) Fire(entry *logrus.Entry) error {
	prompt := fmt.Sprintf(`how would you make this %s log entry more like a ghetto sentence? The output should be already valid to be the log entry, don't give me options and chose one. Feel free to use emojis.
%s`, entry.Level, entry.Message)
	output, err := g.oc.GetOutput(prompt)
	if err == nil {
		entry.Message = output
	}
	return nil
}
