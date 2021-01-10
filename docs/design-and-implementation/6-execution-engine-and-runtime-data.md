
JVM Spec only says a thread should have a VM stack and each item is a Frame. A Frame is associated with an executing method. For non-active frame, it means its state should be snapshotted and can be reinstated later. The top frame is the active frame, and accordingly the method and class is active method and class respectively.

A frame can have an operand stack and all the computation will happen in this operand stack. The execution engine fetches or stores values from/to the operand stack. The maximum stack depth is pre-calculated in the compilation time.

For the PC, it is really a pointer to bytecode. I don’t think there is anything special. When a frame is stale, its PC will be snapshotted.

For native Stack, I just ignore this one. All the code in Go is run in Go’s environment and Go runtime take over the control for me.

The key is to carefully control the frame to run the program. Normal method execution won’t cause the frame switch, but return statement or exception throwing will cause frame switch. Apart from the explicit bytecode execution, there are some special cases to be handle, like <clinit> should be executed when a class is loading. This is actually a tricky case and need some dirty work.