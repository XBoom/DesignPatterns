# 原型模式
原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。

原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。


# 使用场景

如果对象的创建成本比较大，而同一个类的不同对象之间差别不大（大部分字段都相同） 在这种情况下，我们可以利用对已有对象（原型）进行复制（或者叫拷贝）的方式来创建新
对象，以达到节省创建时间的目的。

为何对象的创建成功比较大？
答：创建对象包含的申请内存、给成员变量赋值这一过程，本身并不会花费太多时间，
或者说对于大部分业务系统来说，这点时间完全是可以忽略的。