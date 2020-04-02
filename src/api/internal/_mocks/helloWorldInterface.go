package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type mockContext struct {
	mock.Mock
}

func (m *mockContext) Ping(c *gin.Context) {
	m.Called(c)
}
