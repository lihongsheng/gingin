#### 一致性hash函数
   使用crc32哈希函数获取节点的hash值或者请求key的hash值，使节点的值分布在2的32次方之间。

   请求key过来过来的时候，利用hash函数算出hash值，然后遍历已添加的节点的hash值，顺时针查找
如果，请求的hash值小于当前节点值，则这个请求落在当前节点的前边的节点上【也就是当前节点的下标-1】.

   结构体解释：
```go
type hashMap struct {
	Hash HashFun //hash函数
	HashValue []int //用来存储节点的hash值，且是有序的，后续这个slice遍历来获取请求的key落在的节点hash值，然后通过Mapping获取节点内容
	Mapping map[int]string //用来存储节点的hash值与节点的对应，
}
```
