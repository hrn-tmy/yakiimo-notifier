package constant

// validation messages
const (
	Required     = "%sは必須です。"
	Min          = "%sは%d個以上を指定してください。"
	MinLength    = "%sは%d以上で指定してください。"
	MaxLength    = "%sは%d以下で指定してください。"
	Format       = "%sは正しい形式で指定してください。"
	PasswordKind = "パスワードには、大文字・小文字・数字・記号のうち3種類を組み合わせてください。"
)

// validation items
const (
	Name       = "お名前"
	Email      = "メールアドレス"
	Password   = "パスワード"
	MachineID  = "機械ID"
	FinishedAt = "焼き上がり時間"
	Quantity   = "数量"
)
