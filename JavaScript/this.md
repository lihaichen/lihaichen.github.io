## 为什么有this

讲为啥有this,首先讲对象,一个对象一般会存在属性和方法,一般方法是用于修改或者读取属性,如果没有this,在执行方法时需要将对象传递给方法,方便方法修改属性或者访问属性。
this就是函数运行时所在的对象（环境）。

## this指向谁

在java、c++这类的语言中this就是指向调用方法的对象。
但是JavaScript 支持运行环境动态切换，也就是说，this的指向是动态的。

- this是方法（函数）调用时产生的，和定义（声明）无关。
- this永远指向调用方法时的上下文。

### 例子

```
this === window // true
// 默认this就是指向window上下文。

function f() {return this;}
f() === window // true
// 调用函数f是在window上下文中调用。

obj = {}
obj.f = f
obj.f() === obj // true
// 将f函数赋值给obj对象，通过obj对象调用f，this指向obj。this和定义无关，只和调用有关。

var foo = obj.f
foo()=== obj // false
// 将obj.f 赋值给foo，foo只是引用obj.f的函数。调用foo函数，此时还是在window上下文中调用。this指向window。

function f1() {
    function f2(){
        return this;
    }
    return f2();
}
f1() === window // true
obj.f1 = f1
obj.f1() === obj // false
// 在f1函数内调用f2函数，无论f1本事在什么上下文中，f2的调用依然在默认上下文中。应避免多层this使用。

```

 一般使用回调函数容易改变this的指向，因为调用回调函数的位置一般是默认上下文。类似函数嵌套调用。但是js一般都是通过回调实现，那么就需要this绑定或者通过this传递（闭包或者参数）。


## this的绑定

## 参考

https://wangdoc.com/javascript/oop/this.html