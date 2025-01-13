package mock

import (
	"github.com/stretchr/testify/mock"
)

const (
	indexZero int = iota
	indexFirst
	indexSecond
)

// ConfigMock - structure of configs mock
type ConfigMock struct {
	mock.Mock
}

// Run - mocked parse config files, env and flags - initializes config
func (cm *ConfigMock) Run() error {
	args := cm.Called()

	return args.Error(indexZero)
}

// GetBaseURL - mocked return base url
func (cm *ConfigMock) GetBaseURL() string {
	args := cm.Called()

	return args.String(indexZero)
}

// GetDbURL - mocked return database url
func (cm *ConfigMock) GetDbURL() string {
	args := cm.Called()

	return args.String(indexZero)
}

// GetFileSize - return file size limit
func (cm *ConfigMock) GetFileSize() int64 {
	args := cm.Called()

	return args.Get(indexZero).(int64)
}

// NewConfigMock - return new mocked config
func NewConfigMock() *ConfigMock {
	return new(ConfigMock)
}
