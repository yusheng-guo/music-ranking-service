# music-ranking
A crawler developed in Golang: crawls music charts and sends this information to specified emails address at regular intervals.

### 响应数据`JSON`格式
```json
{
	"code": 0,
	"data": null,
	"message": ""
}
```
| 参数     |   说明                         |
|----------|-------------------------------|
| code     |  `=7` 请求失败，`=0` 请求成功   |
| data     |   请求成功时获取的数据          |
| message  |   请求携带消息                 |