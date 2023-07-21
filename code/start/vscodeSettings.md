vscode go插件配置,仅供参考

```
{
  "go.useLanguageServer": true,
  "[go]": {
    "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "gopls": {
    "completeUnimported": true,
    "usePlaceholders": true,
    "completionDocumentation": true,
    "deepCompletion": true,
    "matcher": "fuzzy",
    "hoverKind": "SynopsisDocumentation" // No/Synopsis/Full, default Synopsis
  },
  "files.eol": "\n", // formatting only supports LF line endings
  "go.languageServerExperimentalFeatures": {
    "format": true,
    "autoComplete": true,
    "rename": true,
    "goToDefinition": true,
    "hover": true,
    "signatureHelp": true,
    "goToTypeDefinition": true,
    "goToImplementation": true,
    "documentSymbols": true,
    "workspaceSymbols": true,
    "findReferences": true,
    "diagnostics": false
  },
  "emmet.excludeLanguages": ["markdown"],
  "go.addTags": {
    "options": "",
    "promptForTags": false
  },
  "go.lintTool": "golangci-lint",
  "go.lintFlags": ["--fast"],
  "files.autoSave": "onFocusChange",
  "go.autocompleteUnimportedPackages": true,
  "go.docsTool": "gogetdoc",
  "terminal.integrated.defaultProfile.windows": "GitBash",
  "terminal.integrated.defaultProfile.osx": "zsh",
  "terminal.integrated.profiles.windows": {
    "GitBash": {
      "path": "D:\\Program Files\\Git\\bin\\bash.exe",
      "args": []
    },
    "PowerShell": {
      "source": "PowerShell",

      "overrideName": true,

      "args": ["-NoExit", "/C", "chcp 65001"],

      "icon": "terminal-powershell",

      "env": {
        "TEST_VAR": "value"
      }
    },
    "Command Prompt": {
      "path": "${env:windir}\\System32\\cmd.exe",
      "args": ["-NoExit", "/K", "chcp 65001"],
      "icon": "terminal-cmd"
    }
  },
  "leetcode.endpoint": "leetcode-cn",
  "leetcode.workspaceFolder": "D:\\golang\\leetcode",
  "leetcode.defaultLanguage": "golang",
  "leetcode.filePath": {
    "default": {
      "folder": "",
      "filename": "${id}.${kebab-case-name}.${ext}"
    }
  },
  "leetcode.editor.shortcuts": ["submit", "test", "solution", "description"],
  "leetcode.hint.configWebviewMarkdown": false,
  "go.testFlags": ["-v"],
  "git.suggestSmartCommit": false,
  "go.toolsManagement.autoUpdate": true,
  "[vue]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[markdown]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode",
    "editor.quickSuggestions": {
      "comments": "on",
      "strings": "on",
      "other": "on"
    }
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[html]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "Codegeex.Privacy": false,
  "[scss]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "Codegeex.Survey": false,
  "editor.links": false,
  "[python]": {
    "editor.formatOnType": true
  },
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "google-translate.serverDomain": "https://translate.google.as/",
  "google-translate.secondLanguage": "zh-cn",
  "google-translate.firstLanguage": "en",
  "google-translate.maxSizeOfResult": 10000,
  "vscodeGoogleTranslate.preferredLanguage": "Chinese (Simplified)",
  // Latex workshop
  "latex-workshop.latex.tools": [
    {
      "name": "latexmk",
      "command": "latexmk",
      "args": [
        "-synctex=1",
        "-interaction=nonstopmode",
        "-file-line-error",
        "-pdf",
        "%DOC%"
      ]
    },
    {
      "name": "xelatex",
      "command": "xelatex",
      "args": [
        "-synctex=1",
        "-interaction=nonstopmode",
        "-file-line-error",
        "%DOC%"
      ]
    },
    {
      "name": "pdflatex",
      "command": "pdflatex",
      "args": [
        "-synctex=1",
        "-interaction=nonstopmode",
        "-file-line-error",
        "%DOC%"
      ]
    },
    {
      "name": "bibtex",
      "command": "bibtex",
      "args": ["%DOCFILE%"]
    }
  ],
  "latex-workshop.latex.recipes": [
    {
      "name": "xelatex - bibtex - xelatex*2 ",
      "tools": ["xelatex", "bibtex", "xelatex", "xelatex"]
    },
    {
      "name": "xelatex*2",
      "tools": ["xelatex", "xelatex"]
    },
    {
      "name": "xelatex",
      "tools": ["xelatex"]
    },
    {
      "name": "latexmk",
      "tools": ["latexmk"]
    },
    {
      "name": "pdflatex -> bibtex -> pdflatex*2",
      "tools": ["pdflatex", "bibtex", "pdflatex", "pdflatex"]
    },
    {
      "name": "pdflatex",
      "tools": ["pdflatex"]
    }
  ],
  "latex-workshop.view.pdf.viewer": "tab",
  "latex-workshop.latex.clean.fileTypes": [
    "*.aux",
    "*.bbl",
    "*.blg",
    "*.idx",
    "*.ind",
    "*.lof",
    "*.lot",
    "*.out",
    "*.toc",
    "*.acn",
    "*.acr",
    "*.alg",
    "*.glg",
    "*.glo",
    "*.gls",
    "*.ist",
    "*.fls",
    "*.log",
    "*.fdb_latexmk"
  ],
  "latex-workshop.latex.autoClean.run": "never",
  "latex-workshop.latex.autoBuild.run": "never"
}

```