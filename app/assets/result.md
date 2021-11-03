### ルビ振り結果
{{ range $idx, $line := .results }}
{{$idx}}. 処理結果
    - ふりがな：{{ $line.Furigana }}
    - テキスト：{{ $line.Surface }}
{{ end }}