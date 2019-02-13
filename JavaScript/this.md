## 为什么有this

讲为啥有this,首先讲对象,一个对象一般会存在属性和方法,一般方法是用于修改或者读取属性,如果没有this,在执行方法时需要将对象传递给方法,方便方法修改属性或者访问属性。  
this就是函数运行时所在的对象（环境）。

## this指向谁

在java、c++这类的语言中this就是指向调用方法的对象。   
但是JavaScript 支持运行环境动态切换，也就是说，this的指向是动态的。   

- this是方法（函数）调用时产生的，和定义（声明）无关。  
- this永远指向调用方法时的上下文。


## this的绑定

## 参考  

http://javascript.ruanyifeng.com/oop/this.html