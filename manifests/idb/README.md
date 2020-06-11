## 注意事项

#### 1. yaml文件中的单引号 `''` 需要 转义 `\'\'`

#### 2. sql语句需要修改 `created` `updated`, `version` `meta_data` 中的镜像id

#### 3. 集群中若残留以下argocd资源，则会导致安装失败

```
$ kubectl get crd,clusterrole,ClusterRoleBinding -l app.kubernetes.io/part-of=argocd
```

#### 4. 卸载appcenter

```
$ kubectl delete crds,clusterrole,ClusterRoleBinding  -l app.kubernetes.io/part-of=argocd
$ kubectl delete ns appcenter
```