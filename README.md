# 🧩 GoLang LLD Design Patterns

This repository contains implementations of various **Low-Level Design (LLD)** patterns in **Go (Golang)**.  
Each pattern is implemented with simple and clear examples to help understand their core concepts and usage.

---

## 📚 Overview

Design patterns are proven solutions to recurring problems in software design.  
They help make code more modular, reusable, and maintainable.

This project demonstrates some of the most commonly used **Creational Design Patterns** in Go, including:
- Factory Pattern
- Abstract Factory Pattern
- Builder Pattern
- Singleton Pattern

---

## 🧠 Implemented Design Patterns

### 🏭 Factory Pattern
The **Factory Pattern** provides a way to create objects without specifying the exact class of the object being created.  
It defines an interface for object creation but lets subclasses decide which class to instantiate.  
This pattern promotes loose coupling and makes it easier to introduce new object types without modifying existing code.

**Use case example:**  
Creating different types of shapes, vehicles, or loggers based on a given input type.

---

### 🏢 Abstract Factory Pattern
The **Abstract Factory Pattern** provides an interface to create families of related or dependent objects without specifying their concrete classes.  
It is an extension of the Factory Pattern and is especially useful when the system needs to be independent of how its objects are created or composed.

**Use case example:**  
Creating UI elements for different operating systems (e.g., WindowsButton, MacButton) without changing client code.

---

### 🧱 Builder Pattern
The **Builder Pattern** separates the construction of a complex object from its representation.  
It allows step-by-step construction of objects and can create different representations of an object using the same construction process.

**Use case example:**  
Creating different configurations of a computer (CPU, GPU, RAM) or assembling a meal in a restaurant.

---

### 🔒 Singleton Pattern
The **Singleton Pattern** ensures that a class has only one instance throughout the application lifecycle.  
It also provides a global point of access to that instance.  
This is often used for logging, configuration, caching, or connection pool management.

**Use case example:**  
A single database connection manager or a global logger.

---
