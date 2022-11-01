运行jaeger
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest

jaeger查看地址  http://127.0.0.1:16686

swagger文档 http://ip:9091/swagger-ui/  然后搜索  http://ip:9091/swagger/hello.swagger.json
