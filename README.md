# plog-backend

## ブランチ命名規則

このプロジェクトでは、以下のブランチ命名規則を採用しています。ブランチ名は一貫性を保ち、開発内容を明確にするために重要です。

### ブランチ名の種類

- **feature/機能名**
  - 新しい機能を開発するためのブランチ
  - 例: `feature/user-authentication`, `feature/profile-page`

- **bugfix/バグ内容**
  - バグ修正のためのブランチ
  - 例: `bugfix/login-issue`, `bugfix/profile-picture-upload`

- **design/デザイン内容**
  - デザインに関する変更を行うブランチ
  - 例: `design/landing-page`, `design/dashboard-ui`

### VSCodeのsetting.jsonのGoの設定

```
"[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true,
    "editor.tabSize": 2,
    "editor.detectIndentation": false,
    "editor.renderWhitespace": "all",
    "editor.unicodeHighlight.includeComments": false
}
```