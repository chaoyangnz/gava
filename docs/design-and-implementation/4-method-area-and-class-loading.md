Class loading is a recursive process. In the progression of a class loading, it may trigger other classes loading. To better trace the classloading process, I need each class load invocation to provide a TriggerReason. The categories include:
```
TRIGGER_BY_JAVA_REFLECTION    : 
"java reflection from name"
}
TRIGGER_BY_CHECK_OBJECT_TYPE  : 
"check java object type"
}
TRIGGER_BY_AS_SUPERCLASS      : 
"as super class"
}
TRIGGER_BY_AS_SUPER_INTERFACE : 
"as super interface"
}
TRIGGER_BY_AS_ARRAY_COMPONENT : 
"as array component"
}
TRIGGER_BY_RESOLVE_CLASS_REF  : 
"revolve symbol_ref in constant pool"
}
TRIGGER_BY_NEW_INSTANCE       : 
"new instance"
}
TRIGGER_BY_ACCESS_MEMBER      : 
"access field or method"
}
```
All the class loadings are logged into a trace log: log/classload.log. It is pretty human-readable I suppose.

JVM Spec says, method area is used to hold the Class representations. In my design, method area includes these components:

DefinedClasses registry: store all the defined classes globally

InitiatedClasses retistry: store all the initiated classes globally

StringPool: store all the interned String objects

BootstrapClassLoader: this bootstrap class loader must be implemented by VM in the native language

A class is denoted by a pair (N, Ld): N is the VM normalised name of a class while Ld is its defining class loader. This pair can uniquely identify a class.

There are two categories of classes to be created: non-array class and array class. Non-array class is created by class loader (bootstrap class loader or user-defined class loader). Array class is directly created by VM, because it has no binary representation (typically a class file in the file system).

The tricky thing is to implement the user-defined class loader. The algorithm can be found in

5.3.2. Loading Using a User-defined Class Loader.

The basic flow to load a class by a user-defined class loader, like AppClassLoader, ExtClassLoader, is as following:

1) The first class that is loaded by AppClassLoader is the initial class which contain the public void main(String[] args) method

2) All other classes referenced by this class will be loaded by the defining loader of this initial class.

3) It will invoke the loadClass(name) method of AppClassLoader and this loader implemented the delegation loading model. It firstly asks its parent class loader: ExtClassLoader.

4) If ExtClassLoader can load the requested class, the loading ends. Otherwise, ExtClassLoader will delegate to bootstrap class loader within VM.

5) If all the ancestor class loader cannot load, AppClassLoader will load the class by itself: find the class file in file system, read its content bytes, then invoke defineClass0(bytes) which is a VM native method.

6) Once loading is finished, the defining/initiating class registry records AppClassLoader as the defining and initiating class loader of the requested classes.

Another thing I took a lot of efforts is to initialise a class. Because the class initialisation needs to run the code of a special method <clinit>, it will cause the endless recursion if the execution order is not carefully organised. In my design, I delay the class initialisation to when needed. Typically, when these instructions are executed: invoke_static, put_static, get_static and new. By doing this, it can ensure the class initialisation can be executed on demand.