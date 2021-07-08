package configEvent
/**
监视中心
   外部注入监视器，监视etcd节点的变化
   外部注入观察者，监听到etcd节点变化的事件和结果并处理
 */
type obServer interface {
	notify()
}
type WatchCentore struct {

}