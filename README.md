# Design Patterns in Go

![Gopher gopher](gopher.jpg)

``` go
import "github.com/TameshkCloud/Go-Design-Pattern"
```

Installation Guide

```
$ go get -u github.com/TameshkCloud/Go-Design-Pattern/...
```

 * [Design Patterns](#design-patterns)
 * [Go Versions Supported](#go-versions-supported)

## Design Patterns

Pattern    | Description
-----------|------------
Creational | As the name implies, it groups common practices for creating objects, so object creation is more encapsulated from the users that need those objects.
Structural | As the name implies, help us to shape our applications with commonly used structures and relationships.
Behavioral | Behavioral design patterns are design patterns that identify common communication patterns between objects and realize these patterns. By doing so, these patterns increase flexibility in carrying out this communication.

## Creational

Name       | Description                               
-----------|-------------------------------------------
[`Singleton`](./Creational/Singleton.go) | Ensure a class has only one instance, and provide a global point of access to it.
[`Builder`](./Creational/Builder.go) | Separate the construction of a complex object from its representation, allowing the same construction process to create various representations.
[`Factory Method`](./Creational/Factory_Method.go) | Define an interface for creating a single object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation to subclasses.
[`Prototype`](./Creational/Prototype.go) | Specify the kinds of objects to create using a prototypical instance, and create new objects from the 'skeleton' of an existing object, thus boosting performance and keeping memory footprints to a minimum.
[`Abstract Factory`](./Creational/Abstract_Factory.go) | Provide an interface for creating families of related or dependent objects without specifying their concrete classes.
[`Object Pool`](./Creational/Object_Pool.go) | Avoid expensive acquisition and release of resources by recycling objects that are no longer in use. Can be considered a generalisation of connection pool and thread pool patterns.

## Structural

Name       | Description                               
-----------|-------------------------------------------
[`Composite`](./Structural/Composite.go) | Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.
[`Adapter`](./Structural/Adapter.go) | Convert the interface of a class into another interface clients expect. An adapter lets classes work together that could not otherwise because of incompatible interfaces. The enterprise integration pattern equivalent is the translator.
[`Bridge`](./Structural/Bridge.go) | Decouple an abstraction from its implementation allowing the two to vary independently.
[`Proxy`](./Structural/Proxy.go) | Provide a surrogate or placeholder for another object to control access to it.
[`Facade`](./Structural/Facade.go) | Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.
[`Decorator`](./Structural/Decorator.go) | Attach additional responsibilities to an object dynamically keeping the same interface. Decorators provide a flexible alternative to subclassing for extending functionality.
[`Flyweight`](./Structural/Flyweight.go) | Use sharing to support large numbers of similar objects efficiently.

## Behavioral

Name       | Description                               
-----------|-------------------------------------------
[`Strategy`](./behavioral/strategy.go) | Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy lets the algorithm vary independently from clients that use it.

[`Chain of Responsibility`](./behavioral/chain_of_responsibility.go) | Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.
[`Command`](./behavioral/command.go) | Encapsulate a request as an object, thereby allowing for the parameterization of clients with different requests, and the queuing or logging of requests. It also allows for the support of undoable operations.
[`Template Method`](./behavioral/template_method.go) | Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
[`Memento`](./behavioral/memento.go) | Without violating encapsulation, capture and externalize an object's internal state allowing the object to be restored to this state later.
[`Interpreter`](./behavioral/interpreter.go) | Given a language, define a representation for its grammar 
along with an interpreter that uses the representation to interpret sentences in the language.
[`Visitor`](./behavioral/visitor.go) | Represent an operation to be performed on the elements of an object structure. Visitor lets a new operation be defined without changing the classes of the elements on which it operates.
[`State`](./behavioral/state.go) | Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.
[`Mediator`](./behavioral/mediator.go) | Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it allows their interaction to vary independently.
[`Observer`](./behavioral/observer.go) | Define a one-to-many dependency between objects where a state change in one object results in all its dependents being notified and updated automatically.
[`Iterator`](./behavioral/iterator.go) | Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.

## Lisence
Go_Design_Pattern is licensed under the MIT License.
see the [LICENSE.md](LICENSE.md) file for more details.

## Go Versions Supported

The most recent major version of Go is supported.

[creational-ref]: https://github.com/TameshkCloud/Go-Design-Pattern/Creational
[structural-ref]: https://github.com/TameshkCloud/Go-Design-Pattern/Structural
[behavioral-ref]: https://github.com/TameshkCloud/Go-Design-Pattern/Behavioral
