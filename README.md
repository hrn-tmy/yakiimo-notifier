# yakiimo-notifier

焼き芋の焼き上がりを登録会員にメール通知するシステム

焼き芋機械から「焼き上がり」が報告されると、その機械をお気に入り登録している会員に自動でメール通知を送信します

## 技術スタック

- **言語**: Go 1.25.0
- **Webフレームワーク**: Echo v4
- **データベース**: PostgreSQL（GORM）
- **メール送信**: SMTP（Mailpit）/ AWS SES
- **API定義**: OpenAPI 3.0（oapi-codegen）
- **ローカル開発**: Docker Compose

## アーキテクチャ

クリーンアーキテクチャに基づく層構成です

```go
api/
├── cmd/main.go                  # エントリーポイント
└── internal/
    ├── domain/                  # ドメインモデル
    ├── email/                   # メール送信（SMTP/SES切り替え可能）
    ├── gen/                     # OpenAPIから生成されたコード
    ├── handler/                 # HTTPリクエスト/レスポンス
    ├── infra/                   # DB接続などのインフラ
    ├── router/                  # ルーティング設定
    ├── usecase/                 # ビジネスロジック
    └── repository/              # データ永続化
```

## API エンドポイント

| メソッド | パス            | 説明                 |
| -------- | --------------- | -------------------- |
| POST     | `/users`        | 会員登録             |
| GET      | `/users`        | 通知対象会員の取得   |
| POST     | `/notify/ready` | 焼き上がり通知の送信 |

## ローカル開発環境のセットアップ

### 1. 依存サービスの起動

```bash
cd docker
docker compose up -d
```

以下のサービスが起動します。

| サービス   | 用途                  | ポート             |
| ---------- | --------------------- | ------------------ |
| PostgreSQL | データベース          | 5432               |
| LocalStack | AWS SESエミュレーター | 4566               |
| Mailpit    | メール確認UI          | 8025（SMTP: 1025） |

### 2. 環境変数の設定

`api/.env` を作成します。

### 3. アプリケーションの起動

```bash
cd api
go run cmd/main.go
```

### 4. メールの確認

Mailpit の Web UI（`http://localhost:8025`）で送信されたメールを確認できます。

## メール送信の切り替え

`EMAIL_DRIVER` 環境変数で送信方式を切り替えられます。

| 値     | 送信方式        | 用途         |
| ------ | --------------- | ------------ |
| `smtp` | SMTP（Mailpit） | ローカル開発 |
| `ses`  | AWS SES         | 本番環境     |
