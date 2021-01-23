# 簡単グルメレビューアプリ 〜merpochi〜



## 本アプリについて

本リポジトリはバックエンドです。フロントエンドのリポジトリは下記参照。

[簡単グルメレビューアプリ 〜merpochi〜 フロントエンド](https://github.com/fujisawatk/merpochi-client)


本アプリは、ネイティブアプリです。配信方法はExpo.ioのpublish配信なので、使用するにはExpo.ioのクライアントアプリ導入が必要になります。

iOS : [‎Expo Client on the App Store](https://apps.apple.com/app/apple-store/id982107779)

Android : [Expo - Google Play のアプリ](https://play.google.com/store/apps/details?id=host.exp.exponent&referrer=www)

アプリ導入環境が整いましたら、下記リンクからプロジェクトページに飛び、カメラ等でQRコードを読み込んで頂くと、アプリが起動出来ます。

アプリ : [merpochi_client on Expo](https://expo.io/@fujisawatk/projects/merpochi_client)

また、インフラ環境はGithub Actionsのスケジュール実行で管理されており、API通信がam10:00〜pm16:00の間だけ使用可能となっておりますので、予めご了承ください。

## 概要
日頃からお世話になっている飲食店に、気軽に感謝の気持ちを伝えることが出来る簡単グルメレビューアプリです。
現在地付近の店舗検索に特化しており、アプリを開けば一瞬で付近にあるお店を探す事が出来ます。お食事を終えたら、ボタン一つでレビュー、画像や感想でといった詳細なレビューができる機能とお好きな方法で、感謝を伝える事が出来ます。

## 作成背景
飲食店で働くスタッフに、より多くの感謝の気持ちを伝えたいからです。
既存大手グルメサービスは、サービス自体が情報過多になっており、レビューが十分に機能していないと考えます。また、限られた人の評価により、レビュー自体に偏りが出てきてしまいます。
本アプリは、様々なシチュエーションで気軽にレビューを投稿出来、より多くの人たちのレビューを可視化する事で、飲食店で働くスタッフに感謝を還元できると思い、制作致しました。

## アプリ機能一覧
- ユーザー機能
  - メールアドレス登録・ログイン
  - プロフィール画像設定
  - 編集
- 店舗情報機能
  - 一覧表示
  - 詳細表示(マップ表示、複数枚画像表示)
  - 現在地付近の店舗を自動検索（ジャンル検索含む）
  - 店名・キーワード検索
  - 駅名検索
  - ジャンル検索
- レビュー機能
  - 簡単レビュー投稿(お気に入り)
  - 詳細レビュー投稿(テキスト、5段階評価、複数枚画像) 
  - 詳細レビュー編集・削除
- マイリスト機能
  - 行きたいお店登録(ブックマーク)
  - 行ったお店登録(お気に入り)
- コメント機能（レビューに対して）
  - 登録


## 使用技術
- **フロントエンド** 
  - 開発環境は、Expo CLI(ネイティブアプリ開発支援サービス)を使用
  - フレームワークとして、Vue-Native（React-NativeをVue.jsの使用感で開発できるラッパー）を採用
  - Native Baseを用いた、効率的なUI開発
  - Vuexで状態管理、モジュールで機能別に分けて管理
  - ぐるなびAPIを使用しての、店舗情報取得
  - 画像はbase64エンコード文字列で管理

- **バックエンド**
  - レイヤードアーキテクチャ＋DDD(ドメイン駆動設計)を採用
  - JWTでクライアント(アプリ)〜API間の認証を制御
  - 画像はAmazon S3で管理
  - 駅名情報はcsvに管理、sql操作でDBに格納

- **インフラ**
  - ECRでのDockerイメージ管理、ECS-Fargateでのコンテナ運用
  - CloudFormationでインフラのテンプレート化
  - SSMパラメータストアで環境変数管理
  - SSMエージェントでコンテナ内シェル管理
  - Route53とALBでAPIドメイン固定
  - Github Actionsで自動テスト・自動デプロイ。また、時限式にインフラ環境を構築することによる、運用コスト管理
  - Githubでバージョン管理、操作はコマンド操作のみで開発

## 使用技術一覧
- アプリケーション
  - Vue-Native
  - Golang
  - JWT
  - データベース
    - MySQL

- インフラ
  - 開発
    - Expo CLI
    - Docker/docker-compose
    - Postman

  - 本番
    - Expo CLI
    - AWS
      - ECR
      - ECS-Fargate
      - CloudFormation
      - Systems Manager
      - Route 53
      - S3
      - RDS

  - CI/CD
    - Github Actions

## こだわりポイント
- ネイティブアプリで制作したことです。
想定したシチュエーションが「外出中リアルタイムに評価できる様に」というのを実現したかったので、どこでも気軽に起動できるネイティブアプリとして制作致しました。


- 簡単レビュー機能を実装したことです。
ユーザーが想いを込めてレビュー出来るように登録ボタンのデザインにも拘りましたし、複数ページに登録ボタンを設置することで、ユーザーが登録しやすい環境も整えました。

## 困ったこと
- インフラの運用コストについてです。
コンテナ環境でAPI開発を行ったので、コンテナ運用・管理に優れているFargateを利用しましたが、コストがかなりかかってしまうことが難点でした。Github Actionsのインフラ構築をスケジュール実行で行うことにより、維持コストをかなり削減することができました。

## 参考資料
- **アーキテクチャ**
>
[![Image from Gyazo](https://i.gyazo.com/705d00dd164c997bd487ca9cffbe646a.png)](https://gyazo.com/705d00dd164c997bd487ca9cffbe646a)
>
- **DB設計ER図**
>
[![Image from Gyazo](https://i.gyazo.com/c95dbef9d1d64da8301832eb3e9dd8dd.png)](https://gyazo.com/c95dbef9d1d64da8301832eb3e9dd8dd)

[![Image from Gyazo](https://i.gyazo.com/fc52a1ceeaa70c074e79a1c8956f514f.png)](https://gyazo.com/fc52a1ceeaa70c074e79a1c8956f514f)


## これから実装したい機能
- テスト（実装途中）
- フォロー機能
- SNSログイン機能（LINEのみ？）
- 通知機能