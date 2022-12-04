<!-- trunk-ignore(markdownlint/MD041) -->
<p align="center">
  <img alt="golangci-lint logo" src="assets/fnple-go-logo.png" height="150" />
  <h3 align="center">fnple</h3>
  <p align="center">The simple fintech platform for people</p>
</p>

---

**`fnple`** is a simple fintech platform.

---

## Project Features
- [ ] Ledger
  - [ ] Account Management
  - [ ] Ledger Entries
  - [ ] Settlement
- [ ] Transaction Management
  - [ ] Transaction Creation
  - [ ] Transaction Execution
  - [ ] Transaction Rejection
  - [ ] Transaction Types
    - [ ] Money In
    - [ ] Transfer
    - [ ] Pay
    - [ ] Sell
    - [ ] Buy
  - [ ] Trader Management
  - [ ] Customer Management
- [ ] **v1.0** - API
  - [ ] **v1.0** Rest'ish API
  - [ ] **v1.1** GRPC
- [ ] **v1.1** - Android App
- [ ] **v1.2** - IOS App
- [ ] **v1.3** - Desktop App

---
## Components

| Name          | Description          |    Language    |    Type     |
| ------------- | -------------------- | :------------: | :---------: |
| **fnple-cli** | CLI for everything   |       Go       |     cli     |
| **ledger**     | Management of Habits |       Go       |   service   |
| **fnple-app** | Fnple App for Users  | Flutter / Dart | ui / mobile |

## Tools

| Name    | Description                                              | Language | Type  |
| ------- | -------------------------------------------------------- | :------: | :---: |
| **bob** | Bob the Builder, a light weight mono-repo managment tool |    Go    |  cli  |

---

## Technology

Below is a list of the technology primarily used in this project.

### Backend

https://github.com/stellar/go

- Go (aka golang)
- GRPC, REST, GQL
- K8s & Docker

### FrontEnd

- Flutter & Dart

## Project Structure

Below is a brief description of the project folder structure.

### `/cmd`

Contains the main executeable applications, namely `cli`

### `/internal`

Contains all the private code for this application to work.

### `/package`

Contains all libs/modules which may be shared i.e. `cli`.

### `/scripts`

Scripts to perform various tasks such as `build` or `sca` (static code analysis), etc. This allows us to keep our makefile in the root folder small.

## Project Setup

Below is a brief description of how to setup this project.

### BootStrap

The tools listed below are required for the build and/or packaging process.

| Category | Description                        | URL                       |
| -------- | ---------------------------------- | ------------------------- |
| Linter   | `golangci` is our linter of choice | https://golangci-lint.run |
|          |                                    |                           |

# Giving Thanks

A Big Shout Out! To the people who worked on the following, the work you have done has aided in my learning of the go language and eco-system.

- project layout inspired by <https://github.com/golang-standards/project-layout> and <https://github.com/stellar/go>
- linting from <https://golangci-lint.run>
- go install github.com/bazelbuild/bazelisk@latest
- `pkg/cli` inspired by <https://github.com/spf13/cobra>
- project logo inspired by my horrific graphic design and created with <https://krita.org/en/>
- scany <https://github.com/georgysavva/scany>
- pgx <https://github.com/jackc/pgx>

# Other

## Emoji Legend

| meaning | emoji              | text                 |
| ------- | ------------------ | -------------------- |
| done    | :heavy_check_mark: | `:heavy_check_mark:` |
| wip     | :construction:     | `:construction:`     |
| note    | :memo:             | `:memo:`             |

## Experiments

- proto <https://github.com/brettmostert/trple-proto>, experimented with
  - mono repository for proto defintions
  - generate defintions for multiple languages using buf and the standard grpc libs provided
  - scripted commit of go proto definitons <https://github.com/brettmostert/trple-proto-product-go>
- grpc <https://github.com/brettmostert/trple>, experimented with
  - using generated proto definitions from mono repo for repo
