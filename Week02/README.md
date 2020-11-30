学习笔记
> pkg/errors \
> go2.0 的errors 参考 pkg/errors 的wrap实现

1. 错误需要被记录
2. 应用程序处理错误，保证100%的完整性
3. 错误只被报告一次(处理一次)
4. Wrap将底层错误封装成业务需要的错误，withMessage携带附加信息，在程序顶部或者中间件中统一打日志
5. errors.Cause(err) 获取根因
6. kit库不该wrap error
7. 应用代码中使用errors.New或errors.Errorf 返回
8. 调用包内函数或项目中的函数 error 直接返回（避免多次wrap）
9. **与三方库协作、内部kit库、标准的基础库 需要wrap保存根因**
10. 已处理过的错误（打印、降级处理）不应该往上抛
11. fmt.Errorf 会抛弃掉error中除文本以外的信息
12. go 1.13中新增unwrap方法获取根因
13. go 1.13新增 Is As 两个函数断言错误（Is会自动调unwrap获取根因）