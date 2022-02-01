package log

import "go.uber.org/zap"

// NewZap ...
func NewZap() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
