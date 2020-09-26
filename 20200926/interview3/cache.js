//使用下划线定义的变量是否有特殊的意义
// 没有本质变化。变量前加下划线表示私有变量，函数名前加下划线表示私有函数。
// js没有私有变量的概念，为了维护方便，程序员约定俗成的规范
var _cache = {};
//使用了立即执行的语法(function(){})(),返回了缓存对象
exports.cache = (function () {
    var o = new Object();
    /**
     * 添加一个新缓存
     * @param {*} name 缓存名 
     * @param {*} value 缓存值
     * @param {*} ms 缓存时间 毫秒
     */
    o.addCache = function (name, value,ms) {
        //js中常用的运算符有哪些？
        var time = ms?ms:60000;
        var tcache = name;
        //对象后面加上中括号表示属性，支持动态访问、创建、修改、删除
        _cache[tcache] = value;
        setTimeout(() => {
            //删除对象属性
            delete _cache[tcache];
        }, time)
    }

    /**
     * 获取缓存的值
     * @param {*} name   
     */
    o.getCache = function(name) {
        return _cache[name]
    }

    /**
     * 删除缓存
     * @param name 待删除的缓存名称
     */
    o.delCache = function(name){
        delete _cache[name];
    }
    return o;
}
)();

//js中变量类型太弱了，导致需要结合上下文判断出变量类型，再调用对应的方法。比如数组对象的length属性，一个var变量能不能使用这个属性
// 编译时无法判断出来。这也带来了不同类型的变量使用运算符后结果不同

// 疑问？ 一个对象在内存中如何保存的，其他模块使用该对象时又是如何知道有哪些方法。