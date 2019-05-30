# 简介
秒拍希望能够将APK包解析的结果通过回调的方式通知秒拍，同时提供状态查询的接口
此 ufop 主要是解析秒拍需要的APK内容，内容的保存及再次查询可以通过 管道 保存到空间，然后查看保存的文件
ufop 解析的内容有

```json
{
	"app_name":	"应用名称",
	"package_name":	"包名",
	"version":	"应用版本",
	"version_code": "应用版本号",
	"size":	"应用大小",
	"md5":	"file_md5"
}
```

# 命令
该命令名称为`aparser`，对应的ufop实例名称为`ufop_prefix`+`aparser`。

```
?aparser
```

# 参数

无

# 常见错误

|错误信息|描述|
|-------|------|
|invalid parser command format|发送的ufop的指令格式不正确，请参考上面的命令格式设置正确的指令|
|retrieve resource apk failed|APK文件下载失败，可以直接重试|
|calcul apk's md5 failed|计算md5失败|
|get apk's size failed|获取APK大小失败|
|parse apk failed|解析APK文件失败|