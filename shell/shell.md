入门
输出echo printf，echo基础输出，但是各版本linux支持方式不同，难互相移植；进一步有printf，类似c语言，对比echo无自动换行
输入与输出可通过< >改变 >为覆盖写 >>为追加写 |为将前者输出作为后者输入
特殊文件/dev/null写入数据会被丢弃 /dev/tty会被自动重定向到一个终端例如 read content < /dev/tty
set -x在shell中打开命令执行追踪功能，以(+ )形式输出执行命令，set +x关闭功能
搜索
后向引用 \(子表达式\)\1，其中\diget表示第diget个子表达式匹配成功，例如\(['"]\).*\1匹配用单引号或者双引号括起来的内容而忽略到底是单引号还是双引号
正则表达式