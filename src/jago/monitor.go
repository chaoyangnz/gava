package jago

import "sync"

type Monitor struct {
	object     *Object
	owner      *Thread // *rtda.Thread
	l          sync.Locker
	lock       sync.Locker
	entryCount int
	cond       *sync.Cond
}

func NewMonitor(obj *Object) *Monitor {
	m := &Monitor{}
	m.object = obj
	m.l = &sync.Mutex{}
	m.lock = &sync.Mutex{}
	m.cond = sync.NewCond(m.lock)
	return m
}

func (self *Monitor) Enter(thread *Thread) {
	LOG.Info("[monitor] thread '%s' #%d enter monitor on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())
	self.l.Lock()
	if self.owner == thread {
		self.entryCount++
		self.l.Unlock()
		return
	} else {
		self.l.Unlock()
	}

	self.lock.Lock()

	self.l.Lock()
	self.owner = thread
	self.entryCount = 1
	self.l.Unlock()
}

func (self *Monitor) Exit(thread *Thread) {
	LOG.Info("[monitor] thread '%s' #%d exit monitor on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())
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

func (self *Monitor) Wait() {
	thread := VM_currentThread()
	LOG.Info("[monitor] thread '%s' #%d wait on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())
	self.l.Lock()
	oldEntryCount := self.entryCount
	oldOwner := self.owner
	self.entryCount = 0
	self.owner = nil
	self.l.Unlock()

	self.cond.Wait()
	LOG.Info("[monitor] thread '%s' #%d wait end on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())

	self.l.Lock()
	self.entryCount = oldEntryCount
	self.owner = oldOwner
	self.l.Unlock()
}

func (self *Monitor) NotifyAll() {
	thread := VM_currentThread()
	LOG.Info("[monitor] thread '%s' #%d notifyAll on object %p %s \n", thread.name, thread.id, self.object, self.object.header.class.Name())
	self.cond.Broadcast()
}
