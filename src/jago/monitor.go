package jago

import (
	"sync"
	"github.com/orcaman/concurrent-map"
	"time"
)

type Monitor struct {
	object     *Object
	owner      *Thread
	l          sync.Locker
	lock       sync.Locker
	entryCount int

	ch         chan int

	waits   cmap.ConcurrentMap
}

func NewMonitor(obj *Object) *Monitor {
	m := &Monitor{}
	m.object = obj
	m.l = &sync.Mutex{}
	m.lock = &sync.Mutex{}
	m.ch = make(chan int)
	m.waits = cmap.New()
	return m
}

func (self *Monitor) Enter() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.Info("[monitor] thread '%s' #%d enter monitor on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	self.l.Lock()
	if self.owner == thread {
		self.entryCount++
		self.l.Unlock()
		self.waits.SetIfAbsent(tid2str(thread.id), thread)
		return
	} else {
		self.l.Unlock()
	}

	self.lock.Lock()

	self.l.Lock()
	self.owner = thread
	self.entryCount = 1
	self.waits.Remove(tid2str(thread.id))
	self.l.Unlock()
}

func (self *Monitor) Exit() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.Info("[monitor] thread '%s' #%d exit  monitor on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	self.l.Lock()
	var _unlock bool
	if self.owner == thread {
		self.entryCount--
		if self.entryCount == 0 {
			self.owner = nil
			_unlock = true
		}
	}
	self.l.Unlock()

	if _unlock {
		self.lock.Unlock()
	}
}

func (self *Monitor) HasOwner(thread *Thread) bool {
	self.l.Lock()
	isOwner := self.owner == thread
	self.l.Unlock()

	return isOwner
}

const _notify  = 1
const _wait_timeout = 2

func (self *Monitor) Wait(millis int64) (interrupted bool) {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.Info("[monitor] thread '%s' #%d wait on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	if thread.interrupted {
		thread.interrupted = false
		return true
	}

	// here current thread must acquire lock
	thread.waiting = true

	self.l.Lock()
	self.Exit()// temporarily release owner
	self.l.Unlock()

	go func() {
		time.Sleep(time.Duration(millis) * time.Millisecond)

		self.ch <- _wait_timeout

		if thread.waiting { // not interrupted
			thread.waiting = false
		}
	}()

	for {
		select {
		case ch := <- self.ch:
			if ch ==_notify || ch == _wait_timeout {
				thread.waiting = false
				interrupted = false
				break
			}
		case ch := <-self.owner.ch:
			if ch == _interrupt {
				thread.waiting = false
				thread.interrupted = false
				interrupted = true
				return
			}
		}
	}


	VM.ExecutionEngine.Info("[monitor] thread '%s' #%d wait end on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	self.l.Lock()
	self.Enter()// again compete to acquire owner
	self.l.Unlock()

	return
}

func (self *Monitor) NotifyAll() {
	thread := VM.CurrentThread()
	VM.ExecutionEngine.Info("[monitor] thread '%s' #%d notifyAll on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	if !self.waits.IsEmpty() {
		self.ch <- _notify
	}

}
