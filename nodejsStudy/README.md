# nodejsStudy

##nodejs基本数据类型有哪些
number string boolean array
前三个直接赋值，array引用赋值意味着复制一个新数组必须深拷贝

##对象如何实现和初始化

##var let const的区别

##nodejs中的class如何使用，constructor是什么？其中this是指什么？

## js、ts函数语法
js 没有限制类型
function myFunction(a, b) {
    return a * b;
}
ts 有类型限制
function add(x: number, y: number): number {
    return x + y;
}

## ts中常量定义
const 常量名：常量类型

## nodejs & mongodb
使用mongoose来插入数据库 scheme
https://www.cnblogs.com/chris-oil/p/9142795.html

## gc
https://blog.csdn.net/qq_40028324/article/details/92970588

## promise 使用方法
问题：async 方法的返回值如何顺序获取到？

## js判断变量类型
instanceof
