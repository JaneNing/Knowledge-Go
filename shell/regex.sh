#!/bin/zsh
content=abcABC123\*ppp
#检查以a开头
echo $content | grep -E '^a'
#检查以p结束
echo $content | grep -E 'p$'
#检查以包含a，任意一个字符，c
echo $content | grep -E 'a.c'
#检查以包含a，任意数量个字符，p
echo $content | grep -E 'a.*p'
#检查以包含a，bcd任意一个字符，c
echo $content | grep -E 'a[bcd]c'
#检查以p出现3次
echo $content | grep -E 'p{3,3}'
#检查以p出现4-5次之间
echo $content | grep -E 'p{4,5}'
#检查以数字或大小写英文出现3-4次之间
echo $content | grpe -E '[[:alnum:]]{3,4}'
#检查以特殊字符
echo $content | grep -E '\*'