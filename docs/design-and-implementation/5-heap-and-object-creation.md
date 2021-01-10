For the heap, currently it acts as an object factory. That means it can create all kinds of objects. Of course, instances of some well-known classes must be provided in this facility.

Typically ones are java.lang.Class, java.lang.Thread, java.lang.String, java.lang.Throwable.

Because the instances of these classes may be directly created by VM rather than in Java code.