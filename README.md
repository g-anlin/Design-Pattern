# 学习设计模式

## 工厂模式
### 概念
工厂模式（Factory Pattern）是最常用的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。
在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。
### 自己的理解
工厂模式实际上就是抽象一个“工厂”用来生产同类型的对象，这些对象都可以抽象出一个共同的方法，从而可以实现一个相同的接口，这样，工厂的生产方法就可以返回这个相同的接口类型。使用者在使用生产方法的时候只需要传入不同的类型标志就可以获得不同的对象实例，相当于对使用者屏蔽了各实现类的内部实现。
### 缺点
每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。这并不是什么好事。

## 抽象工厂模式
抽象工厂模式（Abstract Factory Pattern）是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。
在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。
### 自己的理解
其实就是生产工厂的工厂，相当于给生产不同对象的工厂外面又包了一层
### 缺点
产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。
用自己的话说就是，如果要扩展抽象工厂支持的工厂，那么就需要扩展工厂接口支持的方法，这样就需要在已有的所有工厂类中新增一个新的工厂生产方法，所以新增工厂很麻烦。

## 责任链模式
责任链模式是一种行为设计模式， 允许你将请求沿着处理者链进行发送。收到请求后，每个处理者均可对请求进行处理，或将其传递给链上的下个处理者。
该模式允许多个对象来对请求进行处理，而无需让发送者类与具体接收者类相耦合。链可在运行时由遵循标准处理者接口的任意处理者动态生成。
### 自己的理解
其实就是链表，只不过每个节点不是储存数据了，而是对消息进行处理。比较形象的理解就是把责任链理解成一条流水线，每个节点就是流水线上的处理节点，消息就是流水线上加工的产品。业务中常用的各种流程感觉就是责任链模式。
### 缺点
1、不能保证请求一定被接收。 2、系统性能将受到一定影响，而且在进行代码调试时不太方便，可能会造成循环调用。 3、可能不容易观察运行时的特征，有碍于除错。