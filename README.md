# mux

Version: v0.1.3 (unreleased)

`mux` is a command line tool to manage tmux sessions. Can be installed with

## Install

```sh
brew tap joshddunn/tap
brew install joshddunn/tap/mux
```

## Commands

| command           | description                                     |
| ----------------- | ----------------------------------------------- |
| config            | Open config file (`~/.mux.json`) with `$EDITOR` |
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

## Development

### Shasum

Run the following commands to get the sha256 value

```sh
curl -sL https://github.com/joshddunn/mux/archive/refs/tags/<VERSION>.tar.gz > <VERSION>.tar.gz
shasum --algorithm 256 ./<VERSION>.tar.gz
```
