### 规范
* 不要基于一个slice，创建一个新的slice，因为在没扩容前，两个slice指向同一个底层数组
* 如果变量或是struct里的属性不需要包以外的访问，劲量小写，或者提供get方法。如果需要序列化，记得字段首字母大写
* 传递数组的时候最好使用指针传递，防止复制数组，占用内存过高
### 项目目录结构及说明

### 项目设计

### 模块扩展规范

### 测试规范

### 发布流程

### 资料链接
[官方文档](https://doczhcn.gitbook.io/etcd/index/index/api_grpc_gateway)
[etcd使用方式](https://www.bookstack.cn/read/For-learning-Go-Tutorial/src-chapter14-01.0.md#9rasul)
[etcd使用方式http](https://blog.csdn.net/meifannao789456/article/details/103480842
[go文档](https://cyent.github.io/golang/datatype/struct_anon/)
