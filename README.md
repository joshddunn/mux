# mux

`mux` is a command line tool to manage tmux sessions. Can be installed with

## Install

```sh
brew tap joshddunn/tap
brew install joshddunn/tap/mux
```

## Configuration

The configuration file for `mux` is `~/.mux.json`

| key      | required | default | values                  | description                     |
| -------- | -------- | ------- | ----------------------- | ------------------------------- |
| sessions | yes      | []      | <Session Configuration> | Array of Session configurations |

### Session Configuration

| key          | required | default | values                 | description                    |
| ------------ | -------- | ------- | ---------------------- | ------------------------------ |
| name         | yes      |         | <string>               |                                |
| dir          | yes      |         | <directory>            |                                |
| zeroIndex    | no       | false   | <boolean>              |                                |
| selectWindow | no       | 1       | <number>               |                                |
| windows      | no       | []      | <Window Configuration> | Array of Window configurations |

### Window Configuration

| key          | required | default             | values                 | description                      |
| ------------ | -------- | ------------------- | ---------------------- | -------------------------------- |
| name         | yes      |                     | <string>               |                                  |
| dir          | no       | <Session directory> | <string>               |                                  |
| layout       | no       | default             | default, columns, rows |                                  |
| splitPercent | no       | 35                  | <number>               | Only used for the default layout |
| panes        | no       | []                  | <Pane Configuration>   |                                  |

### Pane Configuration

| key     | required | default            | values    | description |
| ------- | -------- | ------------------ | --------- | ----------- |
| dir     | no       | <Window directory> |           |             |
| command | no       |                    | <string>  |             |
| execute | no       | true               | <boolean> |             |

### Example Configuration

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
