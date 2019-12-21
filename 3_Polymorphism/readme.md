
interfaces are not part of producer libraries - it is part of client. 

In go libs exists and interfaces are defined by user of those libs.


So you can have Reader ([]byte) (int ,error)

 * File: Reader ([]byte) (int ,error)
 * Tcp:  Reader ([]byte)  (int ,error)
 * WebSockets:  Reader ([]byte) (int ,error)
 
 If you go through Go standard libraries you will not see interfaces there
 That main difference from C# or Java which is based on interfaces.
 
 Go is not object oriented language but provides some attributes which provides some functionality of OOP. 
 (interfaces, messages, embedding, polymorphism, ). So we are free to ue design patterns.
 
 
 