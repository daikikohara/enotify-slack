enotify-slack
=============

[![Build Status](https://drone.io/github.com/daikikohara/enotify-slack/status.png)](https://drone.io/github.com/daikikohara/enotify-slack/latest)
[![Coverage Status](https://img.shields.io/coveralls/daikikohara/enotify-slack.svg)](https://coveralls.io/r/daikikohara/enotify-slack?branch=master)


日本語の紹介記事は[こちら](http://qiita.com/kiida/items/373446edd2fb09da82ca)。

## Summary

This is a tool to get event information and send the information to a channel in [Slack](https://slack.com/).
The event information is provided by event provider's site such as [Meetup](http://www.meetup.com/) and [Eventbrite](https://www.eventbrite.com/)(the others are mainly Japanese sites).
The event information will be sent to Slack if the title or description contains keyword specified in the configuration file.

![screenshot](https://raw.github.com/wiki/daikikohara/enotify-slack/images/capture01.png)

## How to use

* Get binary and conf.yml from [release](https://github.com/daikikohara/enotify-slack/releases) and place them in the same directory.
* Configure conf.yml. Comment in the file is descriptive enough.
* Run the binary with `nohup` like `nohup ./enotify-slack &` or run it in screen/tmux session.

## Thanks

enotify-slack uses api shown below.

* Event provider's site
 * [ATND](http://api.atnd.org/)
 * [Connpass](http://connpass.com/about/api/)
 * [Doorkeeper](http://www.doorkeeperhq.com/developer/api)
 * [Eventbrite](http://developer.eventbrite.com/docs/)
 * [Meetup](http://www.meetup.com/meetup_api/)
 * [Partake](https://github.com/partakein/partake/wiki/PARTAKE-Web-API)
 * [StreetAcademy](http://www.street-academy.com/api)
 * [Zusaar](http://www.zusaar.com/doc/api.html)
* Slack
 * [Slack](https://api.slack.com/)

## License

enotify-slack is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
