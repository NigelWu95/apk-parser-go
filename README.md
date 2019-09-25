# 简介
此 ufop 主要是解析 APK 包的内容，解析返回的内容：  

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
该命令名称为`aparser`，对应的 ufop 实例名称为`ufop_prefix`+`aparser`  

## 同步使用方式

```
url?aparser
```

## 异步使用方式
可使用[七牛 pfop 接口](https://developer.qiniu.com/dora/api/3686/pfop-directions-for-use)提交 cmd：`aparser`，使用 pfop 可以通
过 "|"（管道符）使用 [saveas 命令](https://developer.qiniu.com/dora/api/1305/processing-results-save-saveas)将解析结果保存到空间，
然后查看保存的文件，并使用 persistentId 进行结果的查询  

# 参数

无

# 常见错误

|错误信息|描述   |
|-------|------|
|invalid parser command format|发送的ufop的指令格式不正确，请参考上面的命令格式设置正确的指令|
|retrieve resource apk failed |APK文件下载失败，可以直接重试|
|calcul apk's md5 failed      |计算md5失败|
|get apk's size failed        |获取APK大小失败|
|parse apk failed             |解析APK文件失败|