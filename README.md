
# k8s 后台


# API 

## /k8s-admin
```text
暂无描述
```
#### 公共Header参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Query参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 公共Body参数
参数名 | 示例值 | 参数描述
--- | --- | ---
暂无参数
#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
## /k8s-admin/1.clust list
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:10010/cluster/list

#### 请求方式
> GET

#### Content-Type
> form-data

#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"success": true,
	"data": [
		{
			"name": "default"
		}
	],
	"msg": "查询集群成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 200 | Number | 
success | true | - | 成功响应
data | - | Object | 返回数据
data.name | default | String | 集群名称
msg | 查询集群成功 | String | 返回文字描述
## /k8s-admin/2.version information in cluster
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:10010/cluster/version/default

#### 请求方式
> GET

#### Content-Type
> form-data

#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"success": true,
	"data": {
		"major": "1",
		"minor": "22",
		"gitVersion": "v1.22.5",
		"gitCommit": "5c99e2ac2ff9a3c549d9ca665e7bc05a3e18f07e",
		"gitTreeState": "clean",
		"buildDate": "2021-12-16T08:32:32Z",
		"goVersion": "go1.16.12",
		"compiler": "gc",
		"platform": "linux/arm64"
	},
	"msg": "查询成功"
}
```
## /k8s-admin/3.addidation information in cluster
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:10010/cluster/extra/info/default

#### 请求方式
> GET

#### Content-Type
> form-data

#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"success": true,
	"data": {
		"used_cpu": 0.85,
		"total_cpu": 4,
		"used_memory": 251658240,
		"total_memory": 7964762112,
		"readyNodeNum": 1,
		"totalNodeNum": 1
	},
	"msg": "查询成功"
}
```
## /k8s-admin/4.node informaion in cluster
```text
暂无描述
```
#### 接口状态
> 开发中

#### 接口URL
> http://127.0.0.1:10010/cluster/nodes/default

#### 请求方式
> GET

#### Content-Type
> form-data

#### 预执行脚本
```javascript
暂无预执行脚本
```
#### 后执行脚本
```javascript
暂无后执行脚本
```
#### 成功响应示例
```javascript
{
	"code": 200,
	"success": true,
	"data": [
		{
			"name": "docker-desktop",
			"status": "True",
			"taints": null,
			"labels": {
				"beta.kubernetes.io/arch": "arm64",
				"beta.kubernetes.io/os": "linux",
				"kubernetes.io/arch": "arm64",
				"kubernetes.io/hostname": "docker-desktop",
				"kubernetes.io/os": "linux",
				"node-role.kubernetes.io/control-plane": "",
				"node-role.kubernetes.io/master": "",
				"node.kubernetes.io/exclude-from-external-load-balancers": ""
			},
			"os_image": "Docker Desktop",
			"internal_ip": "192.168.65.4",
			"annotations": {
				"kubeadm.alpha.kubernetes.io/cri-socket": "/var/run/dockershim.sock",
				"node.alpha.kubernetes.io/ttl": "0",
				"volumes.kubernetes.io/controller-managed-attach-detach": "true"
			},
			"kernel_version": "5.10.76-linuxkit",
			"kubelet_version": "v1.22.5",
			"creation_timestamp": "2022-04-05T15:12:30+08:00",
			"container_runtime_version": "docker://20.10.12"
		}
	],
	"msg": "查询成功"
}
```
参数名 | 示例值 | 参数类型 | 参数描述
--- | --- | --- | ---
code | 200 | Number | 
success | true | - | 成功响应
data | - | Object | 返回数据
data.name | docker-desktop | String | node名称
data.status | True | String | node
data.taints | - | Object | 
data.labels | - | Object | 
data.labels.beta.kubernetes.io/arch | arm64 | String | 
data.labels.beta.kubernetes.io/os | linux | String | 
data.labels.kubernetes.io/arch | arm64 | String | 
data.labels.kubernetes.io/hostname | docker-desktop | String | 
data.labels.kubernetes.io/os | linux | String | 
data.labels.node-role.kubernetes.io/control-plane | - | Object | 
data.labels.node-role.kubernetes.io/master | - | Object | 
data.labels.node.kubernetes.io/exclude-from-external-load-balancers | - | Object | 
data.os_image | Docker Desktop | String | 
data.internal_ip | 192.168.65.4 | String | 
data.annotations | - | Object | 
data.annotations.kubeadm.alpha.kubernetes.io/cri-socket | /var/run/dockershim.sock | String | 
data.annotations.node.alpha.kubernetes.io/ttl | 0 | String | 
data.annotations.volumes.kubernetes.io/controller-managed-attach-detach | true | String | 
data.kernel_version | 5.10.76-linuxkit | String | 
data.kubelet_version | v1.22.5 | String | 
data.creation_timestamp | 2022-04-05T15:12:30+08:00 | String | 
data.container_runtime_version | docker://20.10.12 | String | 
msg | 查询成功 | String | 返回文字描述




## Reference material
https://github.com/kubernetes/apimachinery/blob/master/pkg/fields/selector.go

