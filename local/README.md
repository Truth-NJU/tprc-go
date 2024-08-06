# 使用本地PB文件
参考：https://trpc.group/zh/docs/languages/go/basics_tutorial/
以下操作均在local目录下进行
1. `go mod init local`初始项目
2. 创建protobuf文件，定义服务所需要的接口
3. `trpc create -p local.proto` 生成桩文件
4. 在生成的go.mod文件中会自动使用如下语句，引入并使用本地pb文件
`replace github.com/Truth-NJU/tprc-go/trpcprotocol/local => ./stub/github.com/Truth-NJU/tprc-go/trpcprotocol/local`
5. main.go中会注册message服务
6. 可以在message.go中具体实现message服务，这里返回消息id
7. 修改trpc_go.yaml文件，可以参照现在项目里的来写 
8. 运行main.go，即可启动服务
9. 运行test.http可以发送请求测试接口，请求接口的格式是trpc.test.local.Message/SendMessage