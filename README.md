# :D

1. media-uploader
2. golang clean-architecture

---

## 1. Media-Uploader

### Spec

- Database : Postgresql 14
- Storage : GCP cloud storage
- Framework : Echo, Gorm

---

## 2. Golang Clean-Arch

`reference`

- https://manakuro.medium.com/clean-architecture-with-go-bce409427d31
- https://github.com/bxcodec/go-clean-arch/tree/9e174b8b0bbdfbab69bc293bb2905b2bb622155c
- https://github.com/evrone/go-clean-template/tree/3a9a40d72d21bb2d620fa76c4c3214e956a2e2e5

### What is Clean Architecture?

<!-- Clean Architecture image -->

![image](https://user-images.githubusercontent.com/37768791/175006980-11eda9ba-c36f-4d6e-86db-00e1f4044ab6.png)

```
Frameworks & Drivers -> Interface Adapter -> Application Business Rules -> Enterprise Business Rules
```

#### `Frameworks & Drivers`

= DB, Devices(Mobile, Desktop), Web

- We call it `infrastructure`
- End side of architecture
- Handle raw data or APIs

#### `Interface Adapters`

= Controllers, Gateways, Presenters

- We call it `adapter`
- Controll the value between data source and business rule handler
- Take role of adapter that connect outer data and business logic

#### `Application Business Rules`

= Use Cases

- We call it `usecase`
- Business logic handler
- Show purpose of this server

#### `Enterprise Business Rules`

= Entities

- We call it `domain`
- 'Enterprise' == No one except our company shouldn't know

## Directory Structure

#### Config

: store data of configuration

#### Domain(Entities)

: stand for 'Entities'

- model

#### Infrastructure

: stand for 'Framework & Drivers'

- datastore
  : create db instances
- router
  : define routing requests

#### Interface(Interface Adapter)

: stand for 'Interface Adaper' which takes role of translator btw `domain` and `infrastructure`. data fit for `usecases & entities` -> data fit for `DB & Web`

<!-- - repository
  : store db handler as a gateway -->

- presenter

  - format strings & dates
  - add presentation data like flags
  - prepare the data to be displayed in the UI

- controller

  - handle API requests
  - receive the user input like DTO
  - validate user input
  - convert the user input into model that use in usecase
  - call the usercase and pass it

#### Usecase

: stand for 'Use cases' that define what this server for(objective functions)
business logic located in here.

- interactor
  : input port(입력값이 들어오는 구간)
- presenter
  : output port(반환값이 통과하는 구간)
- repository
  : store db handler as a gateway

#### Registry

: resolving dependencies
