package javo

func register_java_util_concurrent_atomic_AtomicLong() {
	VM.RegisterNative("java/util/concurrent/atomic/AtomicLong.VMSupportsCS8()Z", JDK_java_util_concurrent_atomic_AtomicLong_VMSupportsCS8)
}

func JDK_java_util_concurrent_atomic_AtomicLong_VMSupportsCS8() Boolean {
	return TRUE
}
