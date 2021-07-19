package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type etcd struct {
  client *clientv3.Client
  kv clientv3.KV
}

func New() (*etcd,error)  {
	// 客户端配置
	config := clientv3.Config{
		Endpoints: []string{"172.27.43.50:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	clt, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	return &etcd{
		client: clt,
	},err
}

func (s *etcd) buildKV() clientv3.KV {
	if s.kv == nil {
		s.kv = clientv3.NewKV(s.client)
	}
	return s.kv
}

func (e *etcd) Get(key string,opts ...clientv3.OpOption)  {
	e.buildKV().Get(context.TODO(),key,opts...)
}

func (e *etcd) Put(key string,value string,opts ...clientv3.OpOption)  {
	e.buildKV().Put(context.TODO(),key,value,opts...)
}
