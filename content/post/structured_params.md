+++
Categories = ["Ruby"]
Description = "Rails で型安全なパラメータバリデーションとフォームオブジェクトを実現する gem structured_params を紹介します。"
Tags = ["Ruby", "Rails", "gem", "structured_params"]
comments = true
date = "2026-05-03T08:07:07+00:00"
title = "Rails の structured_params 入門（型安全な Params / Form Object）"

+++

この記事は https://github.com/Syati/structured_params の **structured_params（Ruby gem）** の紹介です。

Rails のコントローラーやフォーム周りで、こんな悩みはよくあります。

- `params` は基本的に文字列で入り、型変換が散らばる
- ネストした構造（object/array）が増えると `permit` が複雑になる
- バリデーションと変換（cast）の責務がモデルやコントローラーに漏れる

structured_params は、これらを **ActiveModel ベースの Params クラス**として切り出し、型変換・バリデーション・Strong Parameters の permit までを一貫して扱えるようにする gem です。

<!--more-->

## structured_params が提供するもの

structured_params の中心は `StructuredParams::Params` です。

- `attribute` で「期待する型」を宣言できる
- ActiveModel の `validates` がそのまま使える
- ネストした object / array の型変換（cast）を扱える
- Strong Parameters の `permit` リストを自動生成できる
- RBS 型定義があり、型の補完/検査と相性がよい

## クイックスタート

まず `StructuredParams.register_types` を呼んで、型キャストに必要な型を登録します。

```ruby
# Gemfile
gem 'structured_params'

# どこか初期化（例: config/initializers/structured_params.rb）
StructuredParams.register_types
```

## API パラメータバリデーション例

API のリクエストパラメータを「型変換 + バリデーション」して、モデルに渡すところまでを 1 つの Params クラスにまとめられます。

```ruby
class AddressParams < StructuredParams::Params
  attribute :street, :string
  attribute :city, :string
end

class UserParams < StructuredParams::Params
  attribute :name, :string
  attribute :age, :integer
  attribute :score, :integer
  attribute :tags, :array, value_type: :string            # プリミティブ配列
  attribute :address, :object, value_class: AddressParams # ネストオブジェクト

  # 型変換前の生文字列をバリデーション
  validates_raw :score, format: { with: /\A\d+\z/, message: 'must be numeric string' }
  validates :name, presence: true
  validates :age, numericality: { greater_than: 0 }
  validates :score, numericality: { greater_than_or_equal_to: 0 }
end

def create
  permitted = UserParams.permit(params, require: false)
  user_params = UserParams.new(permitted)

  if user_params.valid?
    User.create!(user_params.attributes)
  else
    render json: { errors: user_params.errors }, status: :unprocessable_entity
  end
end
```

`permit` を書き下す代わりに `UserParams.permit(...)` を呼べるので、コントローラー側がすっきりします。

## フォームオブジェクト例

structured_params はフォーム入力の扱いにも向いています。ActiveModel 互換なので、`validates` など Rails の文法でそのまま書けます。

```ruby
class UserRegistrationForm < StructuredParams::Params
  attribute :name, :string
  attribute :email, :string
  attribute :terms_accepted, :boolean

  validates :name, :email, presence: true
  validates :terms_accepted, acceptance: true
end

def create
  form = UserRegistrationForm.new(UserRegistrationForm.permit(params))

  if form.valid?
    User.create!(form.attributes)
    redirect_to root_path
  else
    render :new
  end
end
```

## いつ使う？（使いどころの考え方）

特に効果が大きいのは次のようなケースです。

- API が大きくなり、パラメータの型変換・バリデーションが散らばっている
- ネスト（object/array）が増えて Strong Parameters の管理がつらい
- “フォームの入力値” をモデルに変換するロジックが複雑になってきた

逆に、入力が単純で変換がほぼ不要な場合は、まずは `params.require(...).permit(...)` とモデルのバリデーションだけでも十分です。

## まとめ

- structured_params は Rails で「型安全な Params / Form Object」を作るための gem
- `attribute` と ActiveModel validations で、型変換と検証を 1 つのクラスに集約できる
- Strong Parameters の permit リストも自動生成でき、ネスト構造でも扱いやすい

詳しくは upstream のドキュメント（README / docs）を参照してください。
