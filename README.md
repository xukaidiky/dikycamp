# dikycamp
徐恺的专属github仓库。千里万里。
version 1.0.0709

======记录一：20230709=======

1、进行了vpn代理提交代码到github远程测试项。配置了config文件。

2、20.205.243.160  对应ssh.github.com 443 域名信息。

3、20230710 提交模块二作业，创建一个httpserver服务。

======记录二：20230720=======

1、构建本地镜像。

2、Dockerfile将模块二的httpserver容器化并推送dockerhuber官方镜像仓库。

3、一些问题和解决：

--Dockerfile

    包括拉取基础镜像，定义8090端口，以及把编译后文件复制到容器映射路径，最后定义应用启动入口。

--docker push 报denied: requested access to the resource is denied的问题。

    原因：docker hub新建的repository的基础结构为xkdiky/dikycamp，而构建镜像的repository却是xkdiky/dikycamp/httpserver。镜像不匹配，被禁止了。
    最后通过docker tag修改了镜像名，成功推送。

--docker run
    
    -d 不交互启动。

=============