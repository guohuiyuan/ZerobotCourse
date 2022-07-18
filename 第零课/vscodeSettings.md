vscode go插件配置,仅供参考

```
{
        "go.useLanguageServer": true,
        "timeline.excludeSources": [],
    
        "[go]": {
            "editor.snippetSuggestions": "none",
            "editor.formatOnSave": true,
            "editor.codeActionsOnSave": {
                "source.organizeImports": true
            }
        },
        
        "gopls": {
            "completeUnimported": true,
            "usePlaceholders": true,
            "completionDocumentation": true,
            "deepCompletion": true, 
            "matcher": "fuzzy",
            "hoverKind": "SynopsisDocumentation" // No/Synopsis/Full, default Synopsis
        },
        
        "files.eol": "\n", // formatting only supports LF line endings
    
    
        "go.languageServerExperimentalFeatures": {
            "format": true,
            "autoComplete": true,
            "rename": true,
            "goToDefinition": true,
            "hover": true,
            "signatureHelp": true,
            "goToTypeDefinition": true,
            "goToImplementation": true,
            "documentSymbols": true,
            "workspaceSymbols": true,
            "findReferences": true,
            "diagnostics": false
        },
        "emmet.excludeLanguages": [
            "markdown"
        ],
        "go.addTags": {
            "tags":"gorm",
            "options": "",
            "promptForTags": false,
            "transform": "snakecase"
        },
        "go.lintTool": "golangci-lint",
        "go.lintFlags": [
            "--fast"
        ],
        "files.autoSave": "onFocusChange",
        "go.autocompleteUnimportedPackages": true,
        "go.docsTool": "gogetdoc",
        "leetcode.endpoint": "leetcode-cn",
        "leetcode.workspaceFolder": "D:\\golang\\leetcode",
        "leetcode.defaultLanguage": "golang",
        "leetcode.filePath": {

             "default": {
                 "folder": "",
                 "filename": "${id}.${kebab-case-name}.${ext}"
        }
    },
    "leetcode.editor.shortcuts": [
        "submit",
        "test",
        "solution",
        "description"
    ],
    "leetcode.hint.configWebviewMarkdown": false,
        "go.testFlags": [
        "-v"
    ],
    "terminal.integrated.defaultProfile.windows": "GitBash",
    "terminal.integrated.defaultProfile.osx": "zsh",
    "terminal.integrated.profiles.windows": {
        "PowerShell": {
            "source": "PowerShell",
            "icon": "terminal-powershell"
        },
        "Command Prompt": {
            "path": "${env:windir}\\System32\\cmd.exe",
            "args": [],
            "icon": "terminal-cmd"
        },
        "GitBash": {
            "path": "D:\\Program Files\\Git\\bin\\bash.exe",
            "args": []
        }
    },
 }
```