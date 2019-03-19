## let
### 为什么提出let解决哪些问题
- 块作用域
- 变量提升
- 暂时性死区
- 重复声明
### 为什么提出块作用域
- 内层变量覆盖外层变量
- 一些计数泄露

### 块作用域和函数声明
- 在es5中函数只能在顶层作用域或者函数作用域中声明，不能在块作用域中声明
- ES6的浏览器中，允许在块作用域中声明，声明类似var，会提升到函数或者全局作用域中，但是定义不变。
```
function test()
{
    if(false)
    {
       function f() {}
    }
}

function test()
{
    var f = undefined;
    if(flase)
    {
        f = function () {}
    }
}

```
- 尽量不要再块作用域中声明函数，如果需要改成函数表达式。

## const
- 和let类似
- const 保证的是变量指向的地址不能变更，对应对象来说，对象属性内容是可以谁便更改的。

## 声明变量
- ES5中变量声明方式有var、function。
- ES6中新增了let、const、class、import。
- 对于var和function在顶层声明还是和global或者window挂钩。
- let、const、class 不再和global或者window挂钩。