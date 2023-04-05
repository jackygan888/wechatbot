# wechatbot
目前还在研究，功能还没经过测试，大家发现代码bug可以帮忙修改

###目前实现了以下功能
 + 群聊@回复
 + 私聊回复
 + 自动通过回复
 
# 注册openai
chatGPT注册获取到API key

# 安装使用
````
# 获取项目
git clone https://github.com/jackygan888/wechatbot.git

# 进入项目目录
cd wechatbot

# 复制配置文件
copy config.dev.json config.json

# 启动项目
go run main.go

启动前需替换config中的api_key
