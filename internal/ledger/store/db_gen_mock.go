package store

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock"
)

// DBMock implements DB
type DBMock struct {
	t minimock.Tester

	funcDelete          func(key Key) (err error)
	inspectFuncDelete   func(key Key)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mDBMockDelete

	funcGet          func(key Key) (value []byte, err error)
	inspectFuncGet   func(key Key)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mDBMockGet

	funcNewIterator          func(pivot Key, reverse bool) (i1 Iterator)
	inspectFuncNewIterator   func(pivot Key, reverse bool)
	afterNewIteratorCounter  uint64
	beforeNewIteratorCounter uint64
	NewIteratorMock          mDBMockNewIterator

	funcSet          func(key Key, value []byte) (err error)
	inspectFuncSet   func(key Key, value []byte)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mDBMockSet
}

// NewDBMock returns a mock for DB
func NewDBMock(t minimock.Tester) *DBMock {
	m := &DBMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.DeleteMock = mDBMockDelete{mock: m}
	m.DeleteMock.callArgs = []*DBMockDeleteParams{}

	m.GetMock = mDBMockGet{mock: m}
	m.GetMock.callArgs = []*DBMockGetParams{}

	m.NewIteratorMock = mDBMockNewIterator{mock: m}
	m.NewIteratorMock.callArgs = []*DBMockNewIteratorParams{}

	m.SetMock = mDBMockSet{mock: m}
	m.SetMock.callArgs = []*DBMockSetParams{}

	return m
}

type mDBMockDelete struct {
	mock               *DBMock
	defaultExpectation *DBMockDeleteExpectation
	expectations       []*DBMockDeleteExpectation

	callArgs []*DBMockDeleteParams
	mutex    sync.RWMutex
}

// DBMockDeleteExpectation specifies expectation struct of the DB.Delete
type DBMockDeleteExpectation struct {
	mock    *DBMock
	params  *DBMockDeleteParams
	results *DBMockDeleteResults
	Counter uint64
}

// DBMockDeleteParams contains parameters of the DB.Delete
type DBMockDeleteParams struct {
	key Key
}

// DBMockDeleteResults contains results of the DB.Delete
type DBMockDeleteResults struct {
	err error
}

// Expect sets up expected params for DB.Delete
func (mmDelete *mDBMockDelete) Expect(key Key) *mDBMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("DBMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &DBMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &DBMockDeleteParams{key}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the DB.Delete
func (mmDelete *mDBMockDelete) Inspect(f func(key Key)) *mDBMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for DBMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by DB.Delete
func (mmDelete *mDBMockDelete) Return(err error) *DBMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("DBMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &DBMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &DBMockDeleteResults{err}
	return mmDelete.mock
}

//Set uses given function f to mock the DB.Delete method
func (mmDelete *mDBMockDelete) Set(f func(key Key) (err error)) *DBMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the DB.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the DB.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the DB.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mDBMockDelete) When(key Key) *DBMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("DBMock.Delete mock is already set by Set")
	}

	expectation := &DBMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &DBMockDeleteParams{key},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up DB.Delete return parameters for the expectation previously defined by the When method
func (e *DBMockDeleteExpectation) Then(err error) *DBMock {
	e.results = &DBMockDeleteResults{err}
	return e.mock
}

// Delete implements DB
func (mmDelete *DBMock) Delete(key Key) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(key)
	}

	params := &DBMockDeleteParams{key}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		want := mmDelete.DeleteMock.defaultExpectation.params
		got := DBMockDeleteParams{key}
		if want != nil && !minimock.Equal(*want, got) {
			mmDelete.t.Errorf("DBMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmDelete.DeleteMock.defaultExpectation.results
		if results == nil {
			mmDelete.t.Fatal("No results are set for the DBMock.Delete")
		}
		return (*results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(key)
	}
	mmDelete.t.Fatalf("Unexpected call to DBMock.Delete. %v", key)
	return
}

// DeleteAfterCounter returns a count of finished DBMock.Delete invocations
func (mmDelete *DBMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of DBMock.Delete invocations
func (mmDelete *DBMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to DBMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mDBMockDelete) Calls() []*DBMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*DBMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *DBMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DBMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DBMock.Delete")
		} else {
			m.t.Errorf("Expected call to DBMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to DBMock.Delete")
	}
}

type mDBMockGet struct {
	mock               *DBMock
	defaultExpectation *DBMockGetExpectation
	expectations       []*DBMockGetExpectation

	callArgs []*DBMockGetParams
	mutex    sync.RWMutex
}

// DBMockGetExpectation specifies expectation struct of the DB.Get
type DBMockGetExpectation struct {
	mock    *DBMock
	params  *DBMockGetParams
	results *DBMockGetResults
	Counter uint64
}

// DBMockGetParams contains parameters of the DB.Get
type DBMockGetParams struct {
	key Key
}

// DBMockGetResults contains results of the DB.Get
type DBMockGetResults struct {
	value []byte
	err   error
}

// Expect sets up expected params for DB.Get
func (mmGet *mDBMockGet) Expect(key Key) *mDBMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DBMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &DBMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &DBMockGetParams{key}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the DB.Get
func (mmGet *mDBMockGet) Inspect(f func(key Key)) *mDBMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for DBMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by DB.Get
func (mmGet *mDBMockGet) Return(value []byte, err error) *DBMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DBMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &DBMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &DBMockGetResults{value, err}
	return mmGet.mock
}

//Set uses given function f to mock the DB.Get method
func (mmGet *mDBMockGet) Set(f func(key Key) (value []byte, err error)) *DBMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the DB.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the DB.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the DB.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mDBMockGet) When(key Key) *DBMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("DBMock.Get mock is already set by Set")
	}

	expectation := &DBMockGetExpectation{
		mock:   mmGet.mock,
		params: &DBMockGetParams{key},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up DB.Get return parameters for the expectation previously defined by the When method
func (e *DBMockGetExpectation) Then(value []byte, err error) *DBMock {
	e.results = &DBMockGetResults{value, err}
	return e.mock
}

// Get implements DB
func (mmGet *DBMock) Get(key Key) (value []byte, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(key)
	}

	params := &DBMockGetParams{key}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.value, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		want := mmGet.GetMock.defaultExpectation.params
		got := DBMockGetParams{key}
		if want != nil && !minimock.Equal(*want, got) {
			mmGet.t.Errorf("DBMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmGet.GetMock.defaultExpectation.results
		if results == nil {
			mmGet.t.Fatal("No results are set for the DBMock.Get")
		}
		return (*results).value, (*results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(key)
	}
	mmGet.t.Fatalf("Unexpected call to DBMock.Get. %v", key)
	return
}

// GetAfterCounter returns a count of finished DBMock.Get invocations
func (mmGet *DBMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of DBMock.Get invocations
func (mmGet *DBMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to DBMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mDBMockGet) Calls() []*DBMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*DBMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *DBMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DBMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DBMock.Get")
		} else {
			m.t.Errorf("Expected call to DBMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to DBMock.Get")
	}
}

type mDBMockNewIterator struct {
	mock               *DBMock
	defaultExpectation *DBMockNewIteratorExpectation
	expectations       []*DBMockNewIteratorExpectation

	callArgs []*DBMockNewIteratorParams
	mutex    sync.RWMutex
}

// DBMockNewIteratorExpectation specifies expectation struct of the DB.NewIterator
type DBMockNewIteratorExpectation struct {
	mock    *DBMock
	params  *DBMockNewIteratorParams
	results *DBMockNewIteratorResults
	Counter uint64
}

// DBMockNewIteratorParams contains parameters of the DB.NewIterator
type DBMockNewIteratorParams struct {
	pivot   Key
	reverse bool
}

// DBMockNewIteratorResults contains results of the DB.NewIterator
type DBMockNewIteratorResults struct {
	i1 Iterator
}

// Expect sets up expected params for DB.NewIterator
func (mmNewIterator *mDBMockNewIterator) Expect(pivot Key, reverse bool) *mDBMockNewIterator {
	if mmNewIterator.mock.funcNewIterator != nil {
		mmNewIterator.mock.t.Fatalf("DBMock.NewIterator mock is already set by Set")
	}

	if mmNewIterator.defaultExpectation == nil {
		mmNewIterator.defaultExpectation = &DBMockNewIteratorExpectation{}
	}

	mmNewIterator.defaultExpectation.params = &DBMockNewIteratorParams{pivot, reverse}
	for _, e := range mmNewIterator.expectations {
		if minimock.Equal(e.params, mmNewIterator.defaultExpectation.params) {
			mmNewIterator.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmNewIterator.defaultExpectation.params)
		}
	}

	return mmNewIterator
}

// Inspect accepts an inspector function that has same arguments as the DB.NewIterator
func (mmNewIterator *mDBMockNewIterator) Inspect(f func(pivot Key, reverse bool)) *mDBMockNewIterator {
	if mmNewIterator.mock.inspectFuncNewIterator != nil {
		mmNewIterator.mock.t.Fatalf("Inspect function is already set for DBMock.NewIterator")
	}

	mmNewIterator.mock.inspectFuncNewIterator = f

	return mmNewIterator
}

// Return sets up results that will be returned by DB.NewIterator
func (mmNewIterator *mDBMockNewIterator) Return(i1 Iterator) *DBMock {
	if mmNewIterator.mock.funcNewIterator != nil {
		mmNewIterator.mock.t.Fatalf("DBMock.NewIterator mock is already set by Set")
	}

	if mmNewIterator.defaultExpectation == nil {
		mmNewIterator.defaultExpectation = &DBMockNewIteratorExpectation{mock: mmNewIterator.mock}
	}
	mmNewIterator.defaultExpectation.results = &DBMockNewIteratorResults{i1}
	return mmNewIterator.mock
}

//Set uses given function f to mock the DB.NewIterator method
func (mmNewIterator *mDBMockNewIterator) Set(f func(pivot Key, reverse bool) (i1 Iterator)) *DBMock {
	if mmNewIterator.defaultExpectation != nil {
		mmNewIterator.mock.t.Fatalf("Default expectation is already set for the DB.NewIterator method")
	}

	if len(mmNewIterator.expectations) > 0 {
		mmNewIterator.mock.t.Fatalf("Some expectations are already set for the DB.NewIterator method")
	}

	mmNewIterator.mock.funcNewIterator = f
	return mmNewIterator.mock
}

// When sets expectation for the DB.NewIterator which will trigger the result defined by the following
// Then helper
func (mmNewIterator *mDBMockNewIterator) When(pivot Key, reverse bool) *DBMockNewIteratorExpectation {
	if mmNewIterator.mock.funcNewIterator != nil {
		mmNewIterator.mock.t.Fatalf("DBMock.NewIterator mock is already set by Set")
	}

	expectation := &DBMockNewIteratorExpectation{
		mock:   mmNewIterator.mock,
		params: &DBMockNewIteratorParams{pivot, reverse},
	}
	mmNewIterator.expectations = append(mmNewIterator.expectations, expectation)
	return expectation
}

// Then sets up DB.NewIterator return parameters for the expectation previously defined by the When method
func (e *DBMockNewIteratorExpectation) Then(i1 Iterator) *DBMock {
	e.results = &DBMockNewIteratorResults{i1}
	return e.mock
}

// NewIterator implements DB
func (mmNewIterator *DBMock) NewIterator(pivot Key, reverse bool) (i1 Iterator) {
	mm_atomic.AddUint64(&mmNewIterator.beforeNewIteratorCounter, 1)
	defer mm_atomic.AddUint64(&mmNewIterator.afterNewIteratorCounter, 1)

	if mmNewIterator.inspectFuncNewIterator != nil {
		mmNewIterator.inspectFuncNewIterator(pivot, reverse)
	}

	params := &DBMockNewIteratorParams{pivot, reverse}

	// Record call args
	mmNewIterator.NewIteratorMock.mutex.Lock()
	mmNewIterator.NewIteratorMock.callArgs = append(mmNewIterator.NewIteratorMock.callArgs, params)
	mmNewIterator.NewIteratorMock.mutex.Unlock()

	for _, e := range mmNewIterator.NewIteratorMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1
		}
	}

	if mmNewIterator.NewIteratorMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmNewIterator.NewIteratorMock.defaultExpectation.Counter, 1)
		want := mmNewIterator.NewIteratorMock.defaultExpectation.params
		got := DBMockNewIteratorParams{pivot, reverse}
		if want != nil && !minimock.Equal(*want, got) {
			mmNewIterator.t.Errorf("DBMock.NewIterator got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmNewIterator.NewIteratorMock.defaultExpectation.results
		if results == nil {
			mmNewIterator.t.Fatal("No results are set for the DBMock.NewIterator")
		}
		return (*results).i1
	}
	if mmNewIterator.funcNewIterator != nil {
		return mmNewIterator.funcNewIterator(pivot, reverse)
	}
	mmNewIterator.t.Fatalf("Unexpected call to DBMock.NewIterator. %v %v", pivot, reverse)
	return
}

// NewIteratorAfterCounter returns a count of finished DBMock.NewIterator invocations
func (mmNewIterator *DBMock) NewIteratorAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNewIterator.afterNewIteratorCounter)
}

// NewIteratorBeforeCounter returns a count of DBMock.NewIterator invocations
func (mmNewIterator *DBMock) NewIteratorBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNewIterator.beforeNewIteratorCounter)
}

// Calls returns a list of arguments used in each call to DBMock.NewIterator.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmNewIterator *mDBMockNewIterator) Calls() []*DBMockNewIteratorParams {
	mmNewIterator.mutex.RLock()

	argCopy := make([]*DBMockNewIteratorParams, len(mmNewIterator.callArgs))
	copy(argCopy, mmNewIterator.callArgs)

	mmNewIterator.mutex.RUnlock()

	return argCopy
}

// MinimockNewIteratorDone returns true if the count of the NewIterator invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockNewIteratorDone() bool {
	for _, e := range m.NewIteratorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NewIteratorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNewIteratorCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNewIterator != nil && mm_atomic.LoadUint64(&m.afterNewIteratorCounter) < 1 {
		return false
	}
	return true
}

// MinimockNewIteratorInspect logs each unmet expectation
func (m *DBMock) MinimockNewIteratorInspect() {
	for _, e := range m.NewIteratorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DBMock.NewIterator with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NewIteratorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNewIteratorCounter) < 1 {
		if m.NewIteratorMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DBMock.NewIterator")
		} else {
			m.t.Errorf("Expected call to DBMock.NewIterator with params: %#v", *m.NewIteratorMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNewIterator != nil && mm_atomic.LoadUint64(&m.afterNewIteratorCounter) < 1 {
		m.t.Error("Expected call to DBMock.NewIterator")
	}
}

type mDBMockSet struct {
	mock               *DBMock
	defaultExpectation *DBMockSetExpectation
	expectations       []*DBMockSetExpectation

	callArgs []*DBMockSetParams
	mutex    sync.RWMutex
}

// DBMockSetExpectation specifies expectation struct of the DB.Set
type DBMockSetExpectation struct {
	mock    *DBMock
	params  *DBMockSetParams
	results *DBMockSetResults
	Counter uint64
}

// DBMockSetParams contains parameters of the DB.Set
type DBMockSetParams struct {
	key   Key
	value []byte
}

// DBMockSetResults contains results of the DB.Set
type DBMockSetResults struct {
	err error
}

// Expect sets up expected params for DB.Set
func (mmSet *mDBMockSet) Expect(key Key, value []byte) *mDBMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("DBMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &DBMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &DBMockSetParams{key, value}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the DB.Set
func (mmSet *mDBMockSet) Inspect(f func(key Key, value []byte)) *mDBMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for DBMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by DB.Set
func (mmSet *mDBMockSet) Return(err error) *DBMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("DBMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &DBMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &DBMockSetResults{err}
	return mmSet.mock
}

//Set uses given function f to mock the DB.Set method
func (mmSet *mDBMockSet) Set(f func(key Key, value []byte) (err error)) *DBMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the DB.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the DB.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the DB.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mDBMockSet) When(key Key, value []byte) *DBMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("DBMock.Set mock is already set by Set")
	}

	expectation := &DBMockSetExpectation{
		mock:   mmSet.mock,
		params: &DBMockSetParams{key, value},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up DB.Set return parameters for the expectation previously defined by the When method
func (e *DBMockSetExpectation) Then(err error) *DBMock {
	e.results = &DBMockSetResults{err}
	return e.mock
}

// Set implements DB
func (mmSet *DBMock) Set(key Key, value []byte) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(key, value)
	}

	params := &DBMockSetParams{key, value}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		want := mmSet.SetMock.defaultExpectation.params
		got := DBMockSetParams{key, value}
		if want != nil && !minimock.Equal(*want, got) {
			mmSet.t.Errorf("DBMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *want, got, minimock.Diff(*want, got))
		}

		results := mmSet.SetMock.defaultExpectation.results
		if results == nil {
			mmSet.t.Fatal("No results are set for the DBMock.Set")
		}
		return (*results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(key, value)
	}
	mmSet.t.Fatalf("Unexpected call to DBMock.Set. %v %v", key, value)
	return
}

// SetAfterCounter returns a count of finished DBMock.Set invocations
func (mmSet *DBMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of DBMock.Set invocations
func (mmSet *DBMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to DBMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mDBMockSet) Calls() []*DBMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*DBMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *DBMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DBMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DBMock.Set")
		} else {
			m.t.Errorf("Expected call to DBMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to DBMock.Set")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DBMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockDeleteInspect()

		m.MinimockGetInspect()

		m.MinimockNewIteratorInspect()

		m.MinimockSetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DBMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *DBMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone() &&
		m.MinimockNewIteratorDone() &&
		m.MinimockSetDone()
}
