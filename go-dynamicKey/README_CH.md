
当想更新动态密钥，需要外部设置redis：
set pub_key ```"{\"id\": \"9g1\", \"key\": \"hg92g02fewh7qk11\", \"iv\": \"546d4t546tw32e1f\"}"```


- id。密码盐，字符位数改变，客户端和应用后台的密码盐位数也需要改变。统一修改CodeSaltLen变量
- key。aes加密的key
- iv。aes加密的iv



