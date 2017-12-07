### GLIDE Менеджер пакетов

`The easiest way to install the latest release on Mac or Linux is with the following script:

`curl https://glide.sh/get | sh`

On Mac OS X you can also install the latest release via Homebrew:

`$ brew install glide`

On Ubuntu Precise (12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04) you can install from our PPA:

`sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update`
`sudo apt-get install glide`

## ORM Postgres
https://github.com/src-d/go-kallax

## Запуск
Просто  `make run`

**Для след команд нужен kallax ORM**

`cd ~/go && go get https://github.com/src-d/go-kallax/...`

При изменении моделей:
- Для создания миграций `make migrations` 
- Для генерации моделей `make modelsgen`
- Накатить миграции `make migrate action=up env=development`