# bookmarks
bookmarks is a bookmark manager for developers.

# Setup
Start the devcontainer in for example vscode.
In the devcontainer you have go, some vscode plugins, golangci-lint, cobra-cli.

# Run
The current implementation stores bookmarks to a json file in your home folder: /home/user/.config/bookmarks/bookmarks.json

add bookmark:
```bash
go run . add -t "My favorite website" denniswethmar.nl
```
if no title is provided then the url will be queried for an title.

search bookmarks:
```bash
go run . -s .nl
```

list entries:
```bash
go run . ls
```