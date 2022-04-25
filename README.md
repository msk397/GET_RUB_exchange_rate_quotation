# Get the ruble price of Bank of China

* Get the ruble spot exchange rate from Bank of China every five minutes
* Push through [bark](https://github.com/Finb/Bark), only push when the price changes
* Save the bark log in the log
* Use cron for scheduled tasks
* Support changing time zone

## Principle

* Get the HTML code of Bank of China's foreign exchange website through http library
* Parse the HTML code through [goquery](https://github.com/PuerkitoBio/goquery) to get the spot exchange rate of rubles
* Timing tasks are implemented through [cron](https://github.com/robfig/cron/)