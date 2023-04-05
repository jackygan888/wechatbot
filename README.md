# wechatbot
本项目使用了OpenAI最新（2023-03-02）发布的大型语言模型gpt-3.5-turbo以其API和基于[openwechat](https://github.com/eatmoreapple/openwechat)开发。
# 当前版本
当前版本使用的是gpt-3.5-turbo模型，
当前版本不支持配置代理

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
