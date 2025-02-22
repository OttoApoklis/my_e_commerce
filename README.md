# 秒杀电商平台设计
用户表 商品表 库存表 订单表 秒杀记录表 默认限额表 VIP用户限额表

## 项目特点：
### 高并发、高可靠、微服务架构
主要矛盾：库存扣减

### 业务逻辑
商品逻辑、秒杀逻辑、订单逻辑

秒杀客户：进入商品详情页，此时已经拉到了商品的信息
如果没到秒杀时间，页面上会显示倒计时
如果已经到了秒杀时间，每次有用户进入秒杀页面，后段都会返回秒杀是否结束，如果结束，就没必要进行了

秒杀开始到秒杀结束
检查名额是否已经用完,是的话暂时不能点击购买，考虑到可能会有放弃付款的，只是暂时不能下单
如果还有名额，可以点击购买
库存抢夺，主要用于做预扣库存
抢到了就创建订单

从订单到购买完成
1、在规定时间内付款
2、超时未支付，自动取消返还名额


### QPS设计
目前只有一台机器，设计上限2万 QPS，这也是测试得到的redis的上限

#### TODO 按道理应该要做redis高可用的，但是我们就一台机器，难免单点故障了

### 表设计
#### 表业务id使用UUID生成，该字段唯一约束，使用long型存储（非主键id,主键id不适合与业务相关）
#### 对于订单表、库存表、秒杀记录表的更新操作使用订单、库存状态、秒杀库存来实现乐观锁保证更新幂等性
#### 对于默认限额表、 VIP用户限额表使用限额、额度实现乐观锁来实现更新幂等性
#### 用户表、 商品表使用更新时间来实现乐观锁保证更新幂等性

#### 库存表
    主键id
    库存id
    秒杀库存

#### 秒杀记录表
    主键id
    秒杀id
    订单id
    用户id
    商品id
    抢购数量
    秒杀状态 待生成订单 等待支付 支付成功 秒杀主动取消 秒杀超时取消

#### 默认限额表

    主键id
    商品id
    限额数量

#### VIP用户限额表

    主键id
    商品id
    用户id
    剩余额度

#### 订单表

    主键id
    商品id
    买家用户id
    卖家用户id
    商品数量
    商品价格
    订单状态 已创建待支付 支付成功 主动取消 超时取消
    创建时间
    更新时间


#### 商品表
    主键id
    商品id
    商品名词
    商品价格
    商品图片地址
    卖家id
    创建时间
    更新时间


#### 用户表
    主键id
    用户id
    用户名
    密码md5加密
    名
    姓
    创建时间
    更新时间

### 接口设计
### 所有请求头携带参数：
    用户id

#### 秒杀接口
请求字段： 商品编号 商品数量

返回字段：

    状态码0 代表正常返回 返回信息 message携带正常的成功或失败的信息
    
    返回数据：秒杀id

接口逻辑：
    
    查询redis中对应商品的库存，由于redis是单线程，
    可以通过redis的库存实现乐观锁，
    先查库存，再使用这个库存条件更新去扣redis的库存，保证幂等性

#### 秒杀查询接口
请求字段： 秒杀id

返回字段： 

    状态码0 代表正常返回 返回信息 message携带正常的成功或失败的信息
    
    返回的数据： 
        秒杀订单状态 （待生成订单 等待支付 支付成功 秒杀主动取消 秒杀超时取消）
        订单id
        秒杀id
        商品id
        抢购数量

#### 商品查询接口
请求字段： 偏移量 查询数量

返回字段：
    
    状态吗0 代表正常返回 返回信息 message携带正常的成功或失败的信息
    
    返回的数据 商品数据列表(商品名词 商品价格 商品图片地址 卖家id)

#### 查询商品信息
请求字段： 商品id
返回字段：

    状态码0 代表正常返回 返回信息 message携带正常的成功或失败的信息

    返回的数据 商品表信息(商品名词 商品价格 商品图片地址 卖家id)

#### 查询订单列表
请求字段： 偏移量 查询数量

返回字段：

    状态码0 代表正常返回 返回信息 message携带正常的成功或失败的信息
    返回的数据 订单列表（订单id 订单状态 商品id 商品数量 商品价格）


## 工作进度
    用户模块和秒杀模块各自完成，trpc没有合适文档，暂时不整合服务间调度了


# kratos 容易上手的开源go微服务框架
## https://go-kratos.dev/docs/intro/design
## go install github.com/go-kratos/kratos/cmd/kratos/v2@latest


## go get github.com/google/wire/cmd/wire@latest

# TODO
# 整合nacos或etdc
