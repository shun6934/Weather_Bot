# Weather_Bot

## What
石川県のお天気（状態と温度）を知らせてくれるSlackBot。

## Purpose
1. 石川の天気を即座にわかるようにする
2. これからの天気もすぐにわかる。

ex. 石川以外の天気も表示できるようにする。

## Function
- スラッシュコマンド`/weather` を打つとslackで石川の天気を教えてくれる。

## Environment
- go 1.12.5
- dep 0.5.3
- [OpenWeathermap](https://openweathermap.org/)
- [SlackAPI](https://api.slack.com/)

## Set up
1. go install

### dep
1. `go get -u github.com/golang/dep/cmd/dep` でdepをインストール
2. プロジェクト作成
3. `dep init` で初期化 -> `Gopkg.toml`と`Gopkg.lock`、`vendor`が作成される。
4. `dep ensure` でパッケージをインストール。

### OpenWeatherAPI
1. サイトにいき、アカウント作成
2. APIKey取得 = API_KEY

### SlackAPI
1. SlackAPIのサイトで[Create New App](https://api.slack.com/apps?new_app=1)
2. `Features -> Bot Users`で名前を設定して`Add Bot User`
3. `Settings -> Install App`の`Install App to Workspace`でWorkSpaceにボットをインストール
4. `Settings -> Basic Information` にあるVerification Token取得 = VERIFICATION_TOKEN

## Reference
OpenWeatherMap：
- [無料天気予報APIのOpenWeatherMapを使ってみる](https://qiita.com/nownabe/items/aeac1ce0977be963a740)

SlackBot作成：
- [Python3系でSlack Botの作成〜基礎的な対話を実装する](https://qiita.com/croissant1028/items/8d6334b76576762df349)