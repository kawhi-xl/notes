# MongoDB整理
### 一、基础背景
#### 1. nosql和关系型数据库
#### 2. bson数据格式
### 二、简单的CRUD操作
#### 1. 写入insert:   
    db.coll.insertOne()  
    db.coll.insertMany()  
    db.coll.insert() // 插入单条/多条
#### 2. 查询find:  
    db.coll.find()
#### 3. 更新update:  
    db.coll.updateOne()  
    db.coll.updateMany()  
    db.coll.replaceOne()  
    db.coll.update() // 默认更新单个doc,要更新匹配到的所有doc加上multi  
    additional:
    db.coll.findOneAndReplace()  
    db.coll.findOneAndUpdate()  
    db.coll.findAndModify()  
    db.coll.save()  
    db.coll.bulkWrite()  
#### 4. 删除delete:  
    db.coll.deleteOne()  
    db.coll.deleteMany()  
    db.coll.remove()  
    additional:  
    db.coll.findOneAndDelete() // 可以指定排序删除第一个doc  
    db.coll.findAndModify() // 可以指定排序删除第一个doc  
    db.coll.bulkWrite()  
### 三、aggregate聚合管道


### 四、index索引

### 五、transaction事务支持

### 六、replication和sharding(副本和分片)

### 七、change streams变更流

### 参考资料
mongo官方文档:https://docs.mongodb.com/manual/