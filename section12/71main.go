// interface
// カスタムエラーハンドリング


func main() {
	// カスタムエラー型の定義
	type UserError struct {
		Message string
		Code    int
	}

	// Error()メソッドの実装でerrorインターフェースを満たす
	func (e *UserError) Error() string {
		return fmt.Sprintf("エラーコード: %d, メッセージ: %s", e.Code, e.Message)
	}

	// カスタムエラーを生成
	err := &UserError{Message: "不正なユーザー入力です", Code: 400}

	// エラー処理
	if err != nil {
		fmt.Println(err)
	}
}

