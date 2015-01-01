enotify-slack
=============

[![Build Status](https://drone.io/github.com/daikikohara/enotify-slack/status.png)](https://drone.io/github.com/daikikohara/enotify-slack/latest)
[![Coverage Status](https://img.shields.io/coveralls/daikikohara/enotify-slack.svg)](https://coveralls.io/r/daikikohara/enotify-slack?branch=master)

This is a tool to get event information and send the information to a channel in [Slack](https://slack.com/).
The event information is provided by event support sites such as [Connpass](http://connpass.com/).
Currently this tool only gets event information provided by Japanese event support sites.
Although the following part of this document is written in Japanese, Godoc and comments in the source code are written in English.


## 概要

イベント支援サイトから勉強会の情報を取得してSlackに通知するためのツールです。
イベントのタイトルか説明にキーワードが含まれている場合、または指定したユーザが開催(または参加)しているイベントの情報をSlackの指定のチャネルに通知します。

![キャプチャ](https://raw.github.com/wiki/daikikohara/enotify-slack/images/capture01.png)

## 使い方

[バイナリ版](https://drone.io/github.com/daikikohara/enotify-slack/files)はLinuxのみ動作確認しています。<br>
バイナリ版を使う場合は以下の「取得」の項は不要です。設定ファイル(conf.yml)はバイナリと同じディレクトリに配置してください。

* 取得
```
go get github.com/daikikohara/enotify-slack
cd /path/to/enotify-slack
godep restore
go build
```
* 設定<br>
conf.ymlを必要に応じて編集して下さい。<br>
設定方法はconf.ymlのコメントを参照して下さい。
* 起動<br>
nohupを付けるかscreen/tmuxのセッションの中で起動して下さい。
```
nohup ./enotify-slack &
```

## Thanks

以下のAPIを利用させて頂いております。
ありがとうございます。

* イベント情報サイト
 * [ATND](http://api.atnd.org/)
 * [Connpass](http://connpass.com/about/api/)
 * [Doorkeeper](http://www.doorkeeperhq.com/developer/api)
 * [Partake](https://github.com/partakein/partake/wiki/PARTAKE-Web-API)
 * [StreetAcademy](http://www.street-academy.com/api)
 * [Zusaar](http://www.zusaar.com/doc/api.html)
* Slack通知
 * [Slack](https://api.slack.com/)

## License

enotify-slack is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
