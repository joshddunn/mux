# mux

`mux` is a command line tool to manage tmux sessions.

## Install

```sh
brew tap joshddunn/tap
brew install joshddunn/tap/mux
```

## Commands

| command           | description                                     |
| ----------------- | ----------------------------------------------- |
| config            | Open config file (`~/.mux.json`) with `$EDITOR` |
| help              | Link to GitHub homepage                         |
| start \<session\> | Start session in config file                    |
| stop \<session\>  | Stop session in config file                     |

## Config

The config file for `mux` is `~/.mux.json`

| key      | required | default | values             | description              |
| -------- | -------- | ------- | ------------------ | ------------------------ |
| sessions | yes      | []      | \<Session Config\> | Array of Session configs |

### Session Config

| key          | required | default | values            | description             |
| ------------ | -------- | ------- | ----------------- | ----------------------- |
| name         | yes      |         | \<string\>        |                         |
| dir          | yes      |         | \<directory\>     |                         |
| zeroIndex    | no       | false   | \<boolean\>       |                         |
| selectWindow | no       | 1       | \<number\>        |                         |
| windows      | no       | []      | \<Window Config\> | Array of Window configs |

### Window Config

| key          | required | default               | values                 | description                               |
| ------------ | -------- | --------------------- | ---------------------- | ----------------------------------------- |
| name         | yes      |                       | \<string\>             |                                           |
| dir          | no       | \<Session directory\> | \<directory\>          |                                           |
| layout       | no       | default               | default, columns, rows |                                           |
| splitPercent | no       | 35                    | \<number\>             | Only relevant if using the default layout |
| panes        | no       | []                    | \<Pane Config\>        | Array of Pane configs                     |

### Pane Config

| key     | required | default              | values        | description                         |
| ------- | -------- | -------------------- | ------------- | ----------------------------------- |
| dir     | no       | \<Window directory\> | \<directory\> |                                     |
| command | no       |                      | \<string\>    |                                     |
| execute | no       | true                 | \<boolean\>   | Only relevant if command is defined |

### Example Config

```json
{
  "sessions": [
    {
      "name": "dev",
      "dir": "~/code",
      "zeroIndex": true,
      "selectWindow": 1,
      "windows": [
        {
          "name": "nvim editor",
          "panes": [{ "command": "nvim", "execute": false }]
        },
        {
          "name": "code",
          "panes": [{}, {}, {}]
        }
      ]
    }
  ]
}
```

### Config Validation

If you use `neovim`, `coc`, and the `coc-json` plugin, add the following to the `coc-settings.json` file to validate the config file.

```json
"json.schemas": [
{
  "url": "https://raw.githubusercontent.com/joshddunn/mux/refs/heads/main/embed/config.schema.json",
  "fileMatch": [".mux.json"]
}
]
```
