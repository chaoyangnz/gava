package javo

import (
	"sync"
	"time"
)

type Header struct {
	hashCode Int
	class    *Class
	monitor  *Monitor

	// extra data for internal use only
	_thread    *Thread             // for java.lang.Thread object use only
	_type      Type                // for java.lang.Class object use only
	_backTrace []StackTraceElement // for java.lang.Throwable use only
}

type Object struct {
	header Header
	slots  []Value
}

type Monitor struct {
	object     *Object
	owner      *Thread
	l          sync.Locker
	lock       sync.Locker
	entryCount int

	ch chan int

	waits sync.Map
}

func NewMonitor(obj *Object) *Monitor {
	m := &Monitor{}
	m.object = obj
	m.l = &sync.Mutex{}
	m.lock = &sync.Mutex{}
	m.ch = make(chan int)
	m.waits = sync.Map{}
	return m
}

func (this *Monitor) Enter() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d enter monitor of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)

	this.l.Lock()
	if this.owner == thread {
		this.entryCount++
		VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d enter already own monitor of object %p (%s)  -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)
		this.l.Unlock()
		return
	}
	this.waits.LoadOrStore(thread.id, thread)
	this.l.Unlock()

	this.lock.Lock()

	this.l.Lock()
	this.owner = thread
	this.entryCount = 1
	this.waits.Delete(thread.id)
	this.l.Unlock()

	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d acquire monitor of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)
}

func (this *Monitor) Exit() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d exit monitor of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)

	this.l.Lock()
	var _unlock bool
	if this.owner == thread {
		this.entryCount--
		if this.entryCount == 0 {
			this.owner = nil
			_unlock = true
		}
		VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d exit already own monitor of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)
	}
	this.l.Unlock()

	if _unlock {
		this.lock.Unlock()
		VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d release monitor of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)
	}
}

func (this *Monitor) HasOwner(thread *Thread) bool {
	this.l.Lock()
	isOwner := this.owner == thread
	this.l.Unlock()

	return isOwner
}

const _notify = 1
const _wait_timeout = 2

func (this *Monitor) Wait(millis int64) (interrupted bool) {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d wait of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)

	if thread.interrupted {
		thread.interrupted = false
		return true
	}

	// here current thread must acquire lock
	thread.waiting = true

	this.Exit() // temporarily release owner

	if millis != 0 {
		go func() {
			time.Sleep(time.Duration(millis) * time.Millisecond)

			this.ch <- _wait_timeout

			if thread.waiting { // not interrupted
				thread.waiting = false
			}
		}()
	}

	select {
	case ch := <-this.ch:
		if ch == _notify || ch == _wait_timeout {
			thread.waiting = false
			interrupted = false
			break
		}
	case ch := <-VM.CurrentThread().ch:
		if ch == _interrupt {
			thread.waiting = false
			thread.interrupted = false
			interrupted = true
			return
		}
	}

	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d wait end of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)

	this.Enter() // again compete to acquire owner

	return
}

func (this *Monitor) NotifyAll() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d notifyAll of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)

	select {
	case this.ch <- _notify:
		VM.ExecutionEngine.threadsLogger.Info("[monitor] thread '%s' #%d notified of object %p (%s) -> entry: %d \n", thread.name, thread.id, this.object, this.object.header.class.Name(), this.entryCount)
	default:
		// do nothing. no wait
	}
}
