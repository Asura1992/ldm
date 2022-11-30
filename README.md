hello 大佬鼠们，welcome to gayhub(O(∩_∩)O)，这个微服务框架使用 go-micro v4 结合 grpc-gateway网关，雏形尚未成熟，仍需完善
由于有些编译protobuf工具有修改过，小白可以直接吧tool目录下所有执行文件放入/usr/local/bin/ 下面
第一步:进入项目根目录执行 make proto 生成pb文件
第二步：请看Makefile文件命令

链路追踪:
 http://ip:16686

自动文档:
    浏览器打开 http://ip:9091/swagger-ui/  然后搜索  http://ip:9091/swagger/hello.swagger.json
        
