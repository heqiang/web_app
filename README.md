# web_app
##### 基于go web的脚手架  
##### 实现的功能如下 
+ 日志记录zap 及日志读取viper
+ mvc分层 controll层  loggic层及dao层
+ 数据库接入redis及mysql
+ gorm
+ jwt验证，雪花算法做为id
+ swagger文档
+ 采用令牌桶实现限流

docker容器话部署


#### 一 使用
> git clone https://github.com/heqiang/web_app.git
#### 二 配置更改
> 1 修改docker-compose.yaml mysql连接信息更改为自己的  
> 2 对应的项目config.yaml mysql信息更改为自己的  
> 3 其余根据自己的需求进行更改
#### 三 构建
 > docker-compose up 
***注意***
若出现mysql连接失败 根据下面的步骤完成后重新构建即可  
原因：mysql访问机制 root 用户本地可以访问 外网访问不行 所以需要授权
#### 四 mysql 外网访问授权
> 1 grant all privileges on *.* to 'root'@'[host]' identified by '[password]' with grant option;   
> 2 flush privileges;
