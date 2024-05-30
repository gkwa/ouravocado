# ouravocado

Purpose:
Find notes with fewer than 50 words.



## example usage

```bash

make && ./ouravocado index '/Users/mtm/Documents/Obsidian Vault' --ignore-path=.git --ignore-path=.obsidian --ignore-path=trash
ls -la '/Users/mtm/Library/Application Support/ouravocado/index.json'

jq -r 'sort_by(.size) | .[] | select(.wordCount == 0) | "\(.path)"' '/Users/mtm/Library/Application Support/ouravocado/index.json' | xargs -I{} -n1 -d '\n' -a - rm -f "{}"
jq -r 'sort_by(.size) | .[] | select(.wordCount < 50) | "\(.path)"' '/Users/mtm/Library/Application Support/ouravocado/index.json'


```

## install ouravocado


on macos/linux:
```bash

brew install gkwa/homebrew-tools/ouravocado

```


on windows:

```powershell

TBD

```
