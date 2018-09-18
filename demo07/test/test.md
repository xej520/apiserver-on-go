# 创建用户  
>curl -XPOST -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"username":"kong","password":"kong123"}'

# 查询用户列表  
>curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user -d'{"offset": 0, "limit": 20}'

# 获取用户详细信息  
>curl -XGET -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/kong

# 更新用户  
>curl -XPUT -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2 -d'{"username":"kong","password":"kongmodify"}'


# 删除用户  
>curl -XDELETE -H "Content-Type: application/json" http://127.0.0.1:8080/v1/user/2
