# Changelog

## 0.3.0 (2026-05-01)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/railsinfra/rails-go/compare/v0.2.0...v0.3.0)

### Features

* **go:** add default http client with timeout ([614e9c7](https://github.com/railsinfra/rails-go/commit/614e9c7ec7d219459ce08653ae6f31cfb9238d91))
* support setting headers via env ([29aa10c](https://github.com/railsinfra/rails-go/commit/29aa10cc2d87a861afc1a888fe828ffc68fbe175))


### Chores

* avoid embedding reflect.Type for dead code elimination ([e749210](https://github.com/railsinfra/rails-go/commit/e7492106852c6a273e49036adaed3e398474e2aa))
* **internal:** more robust bootstrap script ([42e2ff9](https://github.com/railsinfra/rails-go/commit/42e2ff9e65bb275bd529a337869956f438d16392))

## 0.2.0 (2026-04-21)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/railsinfra/rails-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([8a42ab0](https://github.com/railsinfra/rails-go/commit/8a42ab0e20b800e002b09d1462008a030aa876a0))
* **api:** manual updates ([3ba3265](https://github.com/railsinfra/rails-go/commit/3ba3265ffb795a26ae1af4b1443d72ee89f6ac80))
* **api:** manual updates ([5d31c62](https://github.com/railsinfra/rails-go/commit/5d31c62f9d1c566d2bda1a22d965ae8768bf43e0))
* **api:** manual updates ([b8e499b](https://github.com/railsinfra/rails-go/commit/b8e499b4e0b4c066052eb0845f7a440c8d689e83))
* **internal:** support comma format in multipart form encoding ([9fdcf84](https://github.com/railsinfra/rails-go/commit/9fdcf84a5bfcc7f0294227e54cdfe9a78ab7e65e))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([438d749](https://github.com/railsinfra/rails-go/commit/438d749734c686157bf5b4a1867179711bffd804))
* **client:** use correct format specifier for header serialization ([695376a](https://github.com/railsinfra/rails-go/commit/695376a99bd5195ab4aae5c36c61dca9a916551e))
* **encoder:** correctly serialize NullStruct ([2b8abfe](https://github.com/railsinfra/rails-go/commit/2b8abfe582b288cc7894802e618fd2c63dbf1beb))
* prevent duplicate ? in query params ([ae8f823](https://github.com/railsinfra/rails-go/commit/ae8f823b0a9aa6caa0d33de7dddd3cf03e19f4fe))


### Chores

* **ci:** skip lint on metadata-only changes ([9794292](https://github.com/railsinfra/rails-go/commit/979429243cbb8b88f232edcc1e84c8c164de8e75))
* **ci:** skip uploading artifacts on stainless-internal branches ([b69f489](https://github.com/railsinfra/rails-go/commit/b69f489b0b0f86665791245a8ac3a47f89209c24))
* **ci:** support opting out of skipping builds on metadata-only commits ([8ba7542](https://github.com/railsinfra/rails-go/commit/8ba7542de41534e2ad684676d57ec18648a448a7))
* **client:** fix multipart serialisation of Default() fields ([6388cc9](https://github.com/railsinfra/rails-go/commit/6388cc926dd6f381b7d914231c89cae260327509))
* **internal:** codegen related update ([d2d37d5](https://github.com/railsinfra/rails-go/commit/d2d37d5395b3b2fa98f8cb87e2672824ae424a06))
* **internal:** codegen related update ([b6cd67e](https://github.com/railsinfra/rails-go/commit/b6cd67efe44b8ad4b939b296285faa0f70e86aeb))
* **internal:** codegen related update ([cb842fe](https://github.com/railsinfra/rails-go/commit/cb842fe12f7b92b82983435fd56a60b055326f23))
* **internal:** minor cleanup ([d7aa87b](https://github.com/railsinfra/rails-go/commit/d7aa87b182250a0e751ce1fc4d1e13d520e3291f))
* **internal:** move custom custom `json` tags to `api` ([3620d7f](https://github.com/railsinfra/rails-go/commit/3620d7f2e3595a4be4e58c1c51dd68ebc42e2c03))
* **internal:** remove mock server code ([94c6cf4](https://github.com/railsinfra/rails-go/commit/94c6cf44829253eb57cbafe52a2218b991797dde))
* **internal:** support default value struct tag ([f8602d1](https://github.com/railsinfra/rails-go/commit/f8602d13f6d8aff4d21c32ce3ffaaee99d22e416))
* **internal:** update gitignore ([3e95890](https://github.com/railsinfra/rails-go/commit/3e95890e39a891a6578af43ddf873a150d696c0b))
* **internal:** use explicit returns ([bfe406c](https://github.com/railsinfra/rails-go/commit/bfe406ce5bb45cbef052eb41077239b224e00dbe))
* **internal:** use explicit returns in more places ([f4e85c1](https://github.com/railsinfra/rails-go/commit/f4e85c166c70c0398fe00cbaef1f576fb4e68117))
* remove unnecessary error check for url parsing ([493b3f9](https://github.com/railsinfra/rails-go/commit/493b3f996cdc94cf73a2cab923986c5545a3528f))
* update docs for api:"required" ([e485efe](https://github.com/railsinfra/rails-go/commit/e485efe437ae16b574a2e83159801b05a025ef4c))
* update mock server docs ([266b08a](https://github.com/railsinfra/rails-go/commit/266b08a15a710e46bf4d670105d737df313600a9))
* update SDK settings ([ea26f74](https://github.com/railsinfra/rails-go/commit/ea26f744edfcf10ab3681fa44bff96e9c3bea6c5))
* update SDK settings ([2ddde87](https://github.com/railsinfra/rails-go/commit/2ddde87c8911b6786daf7b415443cddc4caa401b))

## 0.1.0 (2026-01-24)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/sibabale/rails-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** updated the sdk to be inline with rails features ([e617ec2](https://github.com/sibabale/rails-go/commit/e617ec202bdcea8830daad0954a1e2364c70839c))
* **client:** add a convenient param.SetJSON helper ([26ffe9b](https://github.com/sibabale/rails-go/commit/26ffe9bc2e9a9129882f8c7c7fe4d1d8ebe991a5))


### Chores

* update SDK settings ([5a4f728](https://github.com/sibabale/rails-go/commit/5a4f7287a2b96fa94326d273d6cb5f50672a15c2))
