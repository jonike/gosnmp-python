package gosnmp_python

// #cgo pkg-config: python2
// #include <Python.h>
import "C"
import "sync"

// globals

var runningPyPy bool = false
var runningPyPyMutex sync.RWMutex

// private functions

func releaseGIL() *C.PyThreadState {
	var tState *C.PyThreadState

	if C.PyEval_ThreadsInitialized() == 0 {
		return nil
	}

	tState = C.PyEval_SaveThread()

	return tState
}

func reacquireGIL(tState *C.PyThreadState) {
	if tState == nil {
		return
	}

	C.PyEval_RestoreThread(tState)
}

// public functions

func SetPyPy() {
	runningPyPyMutex.Lock()
	runningPyPy = true
	runningPyPyMutex.Unlock()
}

func GetPyPy() bool {
	runningPyPyMutex.RLock()
	val := runningPyPy
	runningPyPyMutex.RUnlock()

	return val
}
