
## 集合创建和删除

#### 更新文档
  
 - 语法格式:
 db.collection.update(
    <query>,          //update的查询条件，类似sql update查询内where后面的。
    <update>,         //update的对象和一些更新的操作符（如$,$inc,$unset,$set...）等，也可以理解为sql update查询内set后面的.
    {
      upsert: <boolean>,      //可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
      multi: <boolean>,       //可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
      writeConcern: <document>        //可选，抛出异常的级别
    }
 ) 
- 示例：将app_key中所有对象（文档），isSelfOSS属性更新为false，没有该属性就新增isSelfOSS属性false
    db.app_key.update({},{$set:{isSelfOSS:false}},{multi:true} )

- 示例2：将某个属性删除
    将app_key中所有对象（文档），selfOSS属性删除掉
    db.app_key.update({},{$unset:{selfOSS:true}},{multi:true} )
