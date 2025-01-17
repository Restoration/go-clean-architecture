// Code generated by MockGen. DO NOT EDIT.
// Source: application/usecase/user_usecase.go
//
// Generated by this command:
//
//	mockgen -source application/usecase/user_usecase.go -destination tests/units/mock/application/usecase/mock_user.go
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	domain "go-clean-app/domain"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
	isgomock struct{}
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockUserUseCase) FindAll(ctx *gin.Context) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockUserUseCaseMockRecorder) FindAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockUserUseCase)(nil).FindAll), ctx)
}
