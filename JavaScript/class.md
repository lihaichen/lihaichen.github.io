### class 是什么
class 本质是构造函数的语法糖。

```
class A {}
typedef A // function
B.prototype.constructor === B // true
```

class内部的方法都是不可枚举的。这点和ES5有点区别。

### getter和setter

### 属性表达式

### class表达式

### 静态方法
只能通过类名调用，静态方法使用this时，this指向类的this，可以通过this调用本类的别的静态函数。
子类通过super调用父类的静态函数。

### new.target
通过new.target判断是否使用new创建，子类继承父类，父类的new.target是子类。

