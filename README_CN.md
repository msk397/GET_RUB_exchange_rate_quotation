# 获取中国银行卢布牌价

* 五分钟获取一次中国银行的卢布现汇牌价
* 通过[bark](https://github.com/Finb/Bark)推送，只有价格变动时才推送
* 将bark日志保存在log中
* 使用cron进行定时任务
* 支持更改时区

## 原理

* 通过http库获取中国银行的外汇网站的HTML代码
* 通过[goquery](https://github.com/PuerkitoBio/goquery)解析HTML代码，获取卢布现汇牌价
* 定时任务通过[cron](https://github.com/robfig/cron/)实现