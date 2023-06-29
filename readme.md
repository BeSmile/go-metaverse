# readme

## 弹幕机器人MAC端

### 启动命令

```bin/bash
// 启动douyu
go run main.go live
//  启动bilibili
 go run main.go lh -p=bilibili
// 启动douyu
 go run main.go lh -p=douyu
```

## todo
- [ ] 支持多平台
  - [x] douYu
  - [ ] 哔哩哔哩
    - [x] 弹幕消息
    - [ ] 欢迎消息
    - [ ] 点赞消息
    - [ ] 礼物消息
- [ ] 支持语音播报
  - [x] 使用siri
  - [ ] 支持训练使用VITS-PaiMon
- [x] 图形化界面显示弹幕
  - [x] 加入直播间提示
  - [x] 弹幕颜色
  - [ ] 礼物提示

- [ ] ui部分使用react+ws重构
- [ ] 实现素材布局之类
- [ ] *弹幕游戏
