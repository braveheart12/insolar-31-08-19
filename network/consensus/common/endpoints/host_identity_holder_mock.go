package endpoints

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "Inbound" can be found in github.com/insolar/insolar/network/consensus/common
*/
import (
	"github.com/insolar/insolar/network/consensus/common/cryptkit"
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
)

//HostIdentityHolderMock implements github.com/insolar/insolar/network/consensus/common.Inbound
type HostIdentityHolderMock struct {
	t minimock.Tester

	GetHostAddressFunc       func() (r Name)
	GetHostAddressCounter    uint64
	GetHostAddressPreCounter uint64
	GetHostAddressMock       mHostIdentityHolderMockGetHostAddress

	GetTransportCertFunc       func() (r cryptkit.CertificateHolder)
	GetTransportCertCounter    uint64
	GetTransportCertPreCounter uint64
	GetTransportCertMock       mHostIdentityHolderMockGetTransportCert

	GetTransportKeyFunc       func() (r cryptkit.SignatureKeyHolder)
	GetTransportKeyCounter    uint64
	GetTransportKeyPreCounter uint64
	GetTransportKeyMock       mHostIdentityHolderMockGetTransportKey
}

//NewHostIdentityHolderMock returns a mock for github.com/insolar/insolar/network/consensus/common.Inbound
func NewHostIdentityHolderMock(t minimock.Tester) *HostIdentityHolderMock {
	m := &HostIdentityHolderMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetHostAddressMock = mHostIdentityHolderMockGetHostAddress{mock: m}
	m.GetTransportCertMock = mHostIdentityHolderMockGetTransportCert{mock: m}
	m.GetTransportKeyMock = mHostIdentityHolderMockGetTransportKey{mock: m}

	return m
}

type mHostIdentityHolderMockGetHostAddress struct {
	mock              *HostIdentityHolderMock
	mainExpectation   *HostIdentityHolderMockGetHostAddressExpectation
	expectationSeries []*HostIdentityHolderMockGetHostAddressExpectation
}

type HostIdentityHolderMockGetHostAddressExpectation struct {
	result *HostIdentityHolderMockGetHostAddressResult
}

type HostIdentityHolderMockGetHostAddressResult struct {
	r Name
}

//Expect specifies that invocation of Inbound.GetNameAddress is expected from 1 to Infinity times
func (m *mHostIdentityHolderMockGetHostAddress) Expect() *mHostIdentityHolderMockGetHostAddress {
	m.mock.GetHostAddressFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetHostAddressExpectation{}
	}

	return m
}

//Return specifies results of invocation of Inbound.GetNameAddress
func (m *mHostIdentityHolderMockGetHostAddress) Return(r Name) *HostIdentityHolderMock {
	m.mock.GetHostAddressFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetHostAddressExpectation{}
	}
	m.mainExpectation.result = &HostIdentityHolderMockGetHostAddressResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of Inbound.GetNameAddress is expected once
func (m *mHostIdentityHolderMockGetHostAddress) ExpectOnce() *HostIdentityHolderMockGetHostAddressExpectation {
	m.mock.GetHostAddressFunc = nil
	m.mainExpectation = nil

	expectation := &HostIdentityHolderMockGetHostAddressExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *HostIdentityHolderMockGetHostAddressExpectation) Return(r Name) {
	e.result = &HostIdentityHolderMockGetHostAddressResult{r}
}

//Set uses given function f as a mock of Inbound.GetNameAddress method
func (m *mHostIdentityHolderMockGetHostAddress) Set(f func() (r Name)) *HostIdentityHolderMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetHostAddressFunc = f
	return m.mock
}

//GetNameAddress implements github.com/insolar/insolar/network/consensus/common.Inbound interface
func (m *HostIdentityHolderMock) GetNameAddress() (r Name) {
	counter := atomic.AddUint64(&m.GetHostAddressPreCounter, 1)
	defer atomic.AddUint64(&m.GetHostAddressCounter, 1)

	if len(m.GetHostAddressMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetHostAddressMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetNameAddress.")
			return
		}

		result := m.GetHostAddressMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetNameAddress")
			return
		}

		r = result.r

		return
	}

	if m.GetHostAddressMock.mainExpectation != nil {

		result := m.GetHostAddressMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetNameAddress")
		}

		r = result.r

		return
	}

	if m.GetHostAddressFunc == nil {
		m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetNameAddress.")
		return
	}

	return m.GetHostAddressFunc()
}

//GetHostAddressMinimockCounter returns a count of HostIdentityHolderMock.GetHostAddressFunc invocations
func (m *HostIdentityHolderMock) GetHostAddressMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetHostAddressCounter)
}

//GetHostAddressMinimockPreCounter returns the value of HostIdentityHolderMock.GetNameAddress invocations
func (m *HostIdentityHolderMock) GetHostAddressMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetHostAddressPreCounter)
}

//GetHostAddressFinished returns true if mock invocations count is ok
func (m *HostIdentityHolderMock) GetHostAddressFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetHostAddressMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetHostAddressCounter) == uint64(len(m.GetHostAddressMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetHostAddressMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetHostAddressCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetHostAddressFunc != nil {
		return atomic.LoadUint64(&m.GetHostAddressCounter) > 0
	}

	return true
}

type mHostIdentityHolderMockGetTransportCert struct {
	mock              *HostIdentityHolderMock
	mainExpectation   *HostIdentityHolderMockGetTransportCertExpectation
	expectationSeries []*HostIdentityHolderMockGetTransportCertExpectation
}

type HostIdentityHolderMockGetTransportCertExpectation struct {
	result *HostIdentityHolderMockGetTransportCertResult
}

type HostIdentityHolderMockGetTransportCertResult struct {
	r cryptkit.CertificateHolder
}

//Expect specifies that invocation of Inbound.GetTransportCert is expected from 1 to Infinity times
func (m *mHostIdentityHolderMockGetTransportCert) Expect() *mHostIdentityHolderMockGetTransportCert {
	m.mock.GetTransportCertFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetTransportCertExpectation{}
	}

	return m
}

//Return specifies results of invocation of Inbound.GetTransportCert
func (m *mHostIdentityHolderMockGetTransportCert) Return(r cryptkit.CertificateHolder) *HostIdentityHolderMock {
	m.mock.GetTransportCertFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetTransportCertExpectation{}
	}
	m.mainExpectation.result = &HostIdentityHolderMockGetTransportCertResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of Inbound.GetTransportCert is expected once
func (m *mHostIdentityHolderMockGetTransportCert) ExpectOnce() *HostIdentityHolderMockGetTransportCertExpectation {
	m.mock.GetTransportCertFunc = nil
	m.mainExpectation = nil

	expectation := &HostIdentityHolderMockGetTransportCertExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *HostIdentityHolderMockGetTransportCertExpectation) Return(r cryptkit.CertificateHolder) {
	e.result = &HostIdentityHolderMockGetTransportCertResult{r}
}

//Set uses given function f as a mock of Inbound.GetTransportCert method
func (m *mHostIdentityHolderMockGetTransportCert) Set(f func() (r cryptkit.CertificateHolder)) *HostIdentityHolderMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetTransportCertFunc = f
	return m.mock
}

//GetTransportCert implements github.com/insolar/insolar/network/consensus/common.Inbound interface
func (m *HostIdentityHolderMock) GetTransportCert() (r cryptkit.CertificateHolder) {
	counter := atomic.AddUint64(&m.GetTransportCertPreCounter, 1)
	defer atomic.AddUint64(&m.GetTransportCertCounter, 1)

	if len(m.GetTransportCertMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetTransportCertMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetTransportCert.")
			return
		}

		result := m.GetTransportCertMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetTransportCert")
			return
		}

		r = result.r

		return
	}

	if m.GetTransportCertMock.mainExpectation != nil {

		result := m.GetTransportCertMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetTransportCert")
		}

		r = result.r

		return
	}

	if m.GetTransportCertFunc == nil {
		m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetTransportCert.")
		return
	}

	return m.GetTransportCertFunc()
}

//GetTransportCertMinimockCounter returns a count of HostIdentityHolderMock.GetTransportCertFunc invocations
func (m *HostIdentityHolderMock) GetTransportCertMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetTransportCertCounter)
}

//GetTransportCertMinimockPreCounter returns the value of HostIdentityHolderMock.GetTransportCert invocations
func (m *HostIdentityHolderMock) GetTransportCertMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetTransportCertPreCounter)
}

//GetTransportCertFinished returns true if mock invocations count is ok
func (m *HostIdentityHolderMock) GetTransportCertFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetTransportCertMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetTransportCertCounter) == uint64(len(m.GetTransportCertMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetTransportCertMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetTransportCertCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetTransportCertFunc != nil {
		return atomic.LoadUint64(&m.GetTransportCertCounter) > 0
	}

	return true
}

type mHostIdentityHolderMockGetTransportKey struct {
	mock              *HostIdentityHolderMock
	mainExpectation   *HostIdentityHolderMockGetTransportKeyExpectation
	expectationSeries []*HostIdentityHolderMockGetTransportKeyExpectation
}

type HostIdentityHolderMockGetTransportKeyExpectation struct {
	result *HostIdentityHolderMockGetTransportKeyResult
}

type HostIdentityHolderMockGetTransportKeyResult struct {
	r cryptkit.SignatureKeyHolder
}

//Expect specifies that invocation of Inbound.GetTransportKey is expected from 1 to Infinity times
func (m *mHostIdentityHolderMockGetTransportKey) Expect() *mHostIdentityHolderMockGetTransportKey {
	m.mock.GetTransportKeyFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetTransportKeyExpectation{}
	}

	return m
}

//Return specifies results of invocation of Inbound.GetTransportKey
func (m *mHostIdentityHolderMockGetTransportKey) Return(r cryptkit.SignatureKeyHolder) *HostIdentityHolderMock {
	m.mock.GetTransportKeyFunc = nil
	m.expectationSeries = nil

	if m.mainExpectation == nil {
		m.mainExpectation = &HostIdentityHolderMockGetTransportKeyExpectation{}
	}
	m.mainExpectation.result = &HostIdentityHolderMockGetTransportKeyResult{r}
	return m.mock
}

//ExpectOnce specifies that invocation of Inbound.GetTransportKey is expected once
func (m *mHostIdentityHolderMockGetTransportKey) ExpectOnce() *HostIdentityHolderMockGetTransportKeyExpectation {
	m.mock.GetTransportKeyFunc = nil
	m.mainExpectation = nil

	expectation := &HostIdentityHolderMockGetTransportKeyExpectation{}

	m.expectationSeries = append(m.expectationSeries, expectation)
	return expectation
}

func (e *HostIdentityHolderMockGetTransportKeyExpectation) Return(r cryptkit.SignatureKeyHolder) {
	e.result = &HostIdentityHolderMockGetTransportKeyResult{r}
}

//Set uses given function f as a mock of Inbound.GetTransportKey method
func (m *mHostIdentityHolderMockGetTransportKey) Set(f func() (r cryptkit.SignatureKeyHolder)) *HostIdentityHolderMock {
	m.mainExpectation = nil
	m.expectationSeries = nil

	m.mock.GetTransportKeyFunc = f
	return m.mock
}

//GetTransportKey implements github.com/insolar/insolar/network/consensus/common.Inbound interface
func (m *HostIdentityHolderMock) GetTransportKey() (r cryptkit.SignatureKeyHolder) {
	counter := atomic.AddUint64(&m.GetTransportKeyPreCounter, 1)
	defer atomic.AddUint64(&m.GetTransportKeyCounter, 1)

	if len(m.GetTransportKeyMock.expectationSeries) > 0 {
		if counter > uint64(len(m.GetTransportKeyMock.expectationSeries)) {
			m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetTransportKey.")
			return
		}

		result := m.GetTransportKeyMock.expectationSeries[counter-1].result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetTransportKey")
			return
		}

		r = result.r

		return
	}

	if m.GetTransportKeyMock.mainExpectation != nil {

		result := m.GetTransportKeyMock.mainExpectation.result
		if result == nil {
			m.t.Fatal("No results are set for the HostIdentityHolderMock.GetTransportKey")
		}

		r = result.r

		return
	}

	if m.GetTransportKeyFunc == nil {
		m.t.Fatalf("Unexpected call to HostIdentityHolderMock.GetTransportKey.")
		return
	}

	return m.GetTransportKeyFunc()
}

//GetTransportKeyMinimockCounter returns a count of HostIdentityHolderMock.GetTransportKeyFunc invocations
func (m *HostIdentityHolderMock) GetTransportKeyMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.GetTransportKeyCounter)
}

//GetTransportKeyMinimockPreCounter returns the value of HostIdentityHolderMock.GetTransportKey invocations
func (m *HostIdentityHolderMock) GetTransportKeyMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.GetTransportKeyPreCounter)
}

//GetTransportKeyFinished returns true if mock invocations count is ok
func (m *HostIdentityHolderMock) GetTransportKeyFinished() bool {
	// if expectation series were set then invocations count should be equal to expectations count
	if len(m.GetTransportKeyMock.expectationSeries) > 0 {
		return atomic.LoadUint64(&m.GetTransportKeyCounter) == uint64(len(m.GetTransportKeyMock.expectationSeries))
	}

	// if main expectation was set then invocations count should be greater than zero
	if m.GetTransportKeyMock.mainExpectation != nil {
		return atomic.LoadUint64(&m.GetTransportKeyCounter) > 0
	}

	// if func was set then invocations count should be greater than zero
	if m.GetTransportKeyFunc != nil {
		return atomic.LoadUint64(&m.GetTransportKeyCounter) > 0
	}

	return true
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *HostIdentityHolderMock) ValidateCallCounters() {

	if !m.GetHostAddressFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetNameAddress")
	}

	if !m.GetTransportCertFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetTransportCert")
	}

	if !m.GetTransportKeyFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetTransportKey")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *HostIdentityHolderMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *HostIdentityHolderMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *HostIdentityHolderMock) MinimockFinish() {

	if !m.GetHostAddressFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetNameAddress")
	}

	if !m.GetTransportCertFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetTransportCert")
	}

	if !m.GetTransportKeyFinished() {
		m.t.Fatal("Expected call to HostIdentityHolderMock.GetTransportKey")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *HostIdentityHolderMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *HostIdentityHolderMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && m.GetHostAddressFinished()
		ok = ok && m.GetTransportCertFinished()
		ok = ok && m.GetTransportKeyFinished()

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if !m.GetHostAddressFinished() {
				m.t.Error("Expected call to HostIdentityHolderMock.GetNameAddress")
			}

			if !m.GetTransportCertFinished() {
				m.t.Error("Expected call to HostIdentityHolderMock.GetTransportCert")
			}

			if !m.GetTransportKeyFinished() {
				m.t.Error("Expected call to HostIdentityHolderMock.GetTransportKey")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *HostIdentityHolderMock) AllMocksCalled() bool {

	if !m.GetHostAddressFinished() {
		return false
	}

	if !m.GetTransportCertFinished() {
		return false
	}

	if !m.GetTransportKeyFinished() {
		return false
	}

	return true
}
