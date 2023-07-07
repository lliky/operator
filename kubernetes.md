
## Group  Version  Resource
*   Group  
Group 即资源组：分为**有名资源组**和**无名资源组**，deployment 为有组名，pod 没有组名。
    ```shell
    # deployment 有 group(apps) 和 version(v1)
    # pod  只有 version(v1)
    POST /apis/apps/v1/namespaces/{namespace}/deployments
    POST /apis/v1/namespaces/{namespace}/pods
    ```
* Version  
kubernetes 的版本分是三种：
  * Alpha: 内部测试版本， 如 v1alpha1
  * Beta: 经历了官方和社区测试的相对稳定的，如 v1beta1
  * Stable: 正式发布版，如 v1, v2
* Resource  
常见的 pod, service, deployment 等都属于资源 
  1. 被实例化的资源即资源对象（ResourceObject）
  2. 资源被分为持久性和非持久性，持久性会保存在 etcd，如 deployment，非持久性如 pod
  3. 资源有 8 种操作：create, delete, deletecollection, get, list, patch, update, watch, 每一种资源都支持其中的一分部
  4. 资源对象描述文件由五部分组成：apiVersion, kind, metadata, spec, status.
  5. API resource link: [kubernetes API Reference Docs](https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/)
```shell
# 查看所有资源
kubectl api-resources -owide

# 查看 apps group 下的资源
kubectl api-resources --api-group apps -owide

# 查看指定资源的详情
kubectl explain pod

# 查看所有 Group 和 Version
kubectl api-versions
```

## Store 的类型
* cache
* UndeltaStore
* FIFO
  * 能精确处理每个对象一次
  * 处理对象的最新版本
  * 不想处理已经删除的对象，从队列中移除
  * 不想定期重新处理对象
* DeltaFIFO
* Heap
* ExpirationCache
### DeltaFIFO
DeltaFIFO 能处理的事情
* 最多处理一次每个对象更改（增量）
* 能够看到之前的所有变化 ？？？
* 想要处理某些对象的删除
* 想要定期处理对象


**事件的产生和消费**
* 生产
  * Reflector 的 List
  * Reflector 的 Watch
  * Reflector 的 Resync

* 消费
  * 事件派发到 work queue
  * 刷新本地缓存

### SharedInformer 
两大功能：
1. 缓存我们关注的资源对象的最新状态的数据
2. 根据资源对象的变化事件来通知我们注册的事件处理方法

有三类 Informer
* NewSharedInformer
* New{deployment/pod}Informer
* NewSharedInformerFactory