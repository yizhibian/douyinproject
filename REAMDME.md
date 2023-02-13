## 一些项目结构

- idl  存放idl及其kitex代码生成器生成的文件
  -  douyin_user  user.thrift生成的文件
  -  user.thrift 
- pkg 项目里有些会依赖到的 还有一些可能没用上
  - 。。。(省略)
  - errno 错误拦截的格式
  - config 配置文件
    - sql sql文件
      - init.sql 初始化用户信息的sql文件
- server 存放服务端
  - api 网关
    - handler 处理层 提供给网关服务端的所有处理器
      - user_handler 用户模块的handler们
        - 。。。(其他具体的handler)
        - handler.go 在这个文件定义了一个结构体 用于绑定请求参数
        - 。。。(其他具体的handler)
      - video_handler 视频模块的handler们（大概这个格式
      - comment_handler 评论模块的handler们（大概这个格式
      - 。。。(其他模块的handler)
    - mw 中间件 目前仅有登陆的中间件
    - pack 封装层 用来封装回复的请求
      - base_resp.go 用于封装一些错误返回的格式
      - user_resp.go 用于封装用户请求返回的数据
      - 。。。(其他模块的pack)
    - router 路由层 定义路由
      - user_router.go 用户模块的路由信息
      - 。。。(其他模块的路由)
      - register.go 通过注册器把所有模块的路由一起注册了
    - rpc 远程调用 调用rpc请求
      - user.go 提供向user模块远程调用的方法
      - 。。。(其他模块远程调用的方法)
      - init.go 初始化这些模块
    - main.go 启动入口 需要先初始化
  - user 用户模块
    - 。。。

# 注意事项
### 最好在idl目录下新建目录来生成代码 因为代码生成完会有点乱而且不能撤回不是很方便
    kitex -service douyinuser(服务名称) ../user/thrift(文件存放地)
    //具体看官网 https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/code_generation/
### 路由层需要加入以下代码才会调用 登陆 的中间件
    user1 := r.Group("/user")
	user1.POST("/login", mw.JwtMiddleware.LoginHandler)
	user1.POST("/register", user_handler.Register)

	//在这之前都是没有登陆拦截的
	//使用这个后即加入登陆中间件 需携带token访问
	user1.Use(mw.JwtMiddleware.MiddlewareFunc())
	user1.GET("/", user_handler.GetUserInfo)
### 在业务代码需要用到用户自身的id可以通过以下代码获取
    //获取token带的id的方法
    //这个 c 是 *app.RequestContext
	if value, exists := c.Get("identity"); exists {
		fmt.Println("the id is ", value)
	}
### 最后一点 关于获取通过rpc获取用户信息的方法 若用户不存在并不会报错 只有连接失败之类的才会报错（我猜的）需要通过判断用户id是否为0从而判断有没有这个用户
	userInfo, err := rpc.GetUserInfo(context.Background(), &req)
	//应该是抛出连接一类的错误
    if err != nil {
		pack.SendBaseResponse(c, errno.ConvertErr(err), nil)
		return
	}
	//这里抛出用户不存在的错误
	if userInfo.GetId() == 0 {
		pack.SendBaseResponse(c, errno.NewErrNo(errno.NilValueErrCode, "user doesnt exit"), nil)
		return
	}