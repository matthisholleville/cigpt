<div align="center">
<br />
<p align="center">
  <a href="https://gitlab.com/tpamp/cigpt">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">cigpt</h3>
</p>
</div>


`cigpt` is an open-source utility developed in Go that retrieves error logs from CI pipelines and uses ChatGPT to perform an initial analysis and provide troubleshooting hints. The goal of this project is to assist all users in comprehending the errors in CI pipelines, which can be extremely complicated to read and understand.

This project is a fork of [k8sgpt](https://github.com/k8sgpt-ai/k8sgpt) and relies on a portion of its code base.

## Installation

### Pre-compiled executables

Get them [here](https://github.com/cigpt-ai/cigpt/releases)

### Verify installation

* Run `cigpt version`

## Quick Start

* Currently the default AI provider is OpenAI, you will need to generate an API key from [OpenAI](https://openai.com)
* Run `cigpt auth` to set it in cigpt.
* Run `cigpt analyze --project-id $GITLAB_PROJECT_ID --pipeline-id $GITLAB_PIPELINE_ID` to run a scan.

https://user-images.githubusercontent.com/99146727/228486925-9bf16909-e4e1-44e8-952f-413aa52533a8.mov

### Tools Supported

- [x] gitlab

## Usage

```
CI jobs debugging powered by AI

Usage:
  cigpt [command]

Available Commands:
  analyze     This command will analyze the error logs of a GitlabCI pipeline.
  auth        Authenticate with your chosen backend
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Print the version number of cigpt

Flags:
      --config string   config file (default is $HOME/.cigpt.yaml)
  -h, --help            help for cigpt
  -t, --toggle          Help message for toggle

Use "cigpt [command] --help" for more information about a command.
```

_Run a scan with the gitlab analyzers_

```
cigpt auth
cigpt analyze --project-id $GITLAB_PROJECT_ID --pipeline-id $GITLAB_PIPELINE_ID
```

## Upcoming major milestones

- [ ] Multiple AI backend support
- [ ] Multiple CI support ( github, jenkins, etc.. )

## Configuration

`cigpt` stores config data in `~/.cigpt.yaml` the data is stored in plain text, including your OpenAI key and CI token.

## Contributing

Coming soon..
