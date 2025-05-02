```
Управление модулями дома

openapi: 3.0.0
info:
  title: Home API
  description: Сервис - управлением модулями групп
  version: 1.0.0
servers:
  - url: https://dev@magistrali.tech:3040
    description: dev стенд
tags:
  - name: Modules
    description: Модули
  - name: Rooms
    description: Комнаты
  - name: Devices
    description: Устройста

paths:
  /api/devices/v0/get:
    get:
      summary: Получение списка доступных устройств
      description: Получение списка устройств
      parameters:
        - name: type
          required: true
          description: Тип модуля
          in: query
          schema:
            type: string
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/DevicesV0GetMethodResponse'
      tags:
        - Devices
  /api/modules/v0/get:
    get:
      summary: Получение всех доступных модулей
      description: Получение информации по модулям.
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ModuleV0GetMethodResponse'
      tags:
        - Modules
  /api/modules/v0/update/{id}:
    post:
      summary: Обновление добавленного модуля
      description: Обновление модуля по id
      parameters:
        - name: id
          required: true
          in: path
          schema:
            type: string
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ModuleV0UpdateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ModuleDisableV0MethodResponse'
      tags:
        - Modules
  /api/modules/v0/delete/{id}:
    delete:
      summary: Отключение модуля
      description: Отключение модуля
      parameters:
        - name: id
          required: true
          in: path
          schema:
            type: string
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ModuleDisableV0MethodResponse'
      tags:
        - Modules
  /api/rooms/v0/get:
    get:
      summary: Получение комнат
      description: Получение комнат дома.
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/RoomsV0GetMethodResponse'
      tags:
        - Rooms
  /api/rooms/v0/get/{id}:
    get:
      summary: Получение информации по комнате
      description: Получение комнат дома.
      parameters:
        - name: id
          required: true
          in: path
          schema:
            type: string
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/RoomsV0GetMethodResponse'
      tags:
        - Rooms
  /api/rooms/v0/create:
    post:
      summary: Добавление комнаты
      description: Получение комнат дома.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RoomsV0CreateMethodResponse'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/RoomsV0CreateMethodResponse'
      tags:
        - Rooms
  /api/rooms/v0/{id}:
    post:
      summary: Обновление комнаты
      description: Получение комнат дома.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/RoomsV0CreateMethodResponse'
      parameters:
        - name: id
          required: true
          in: path
          schema:
            type: string
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/RoomsV0CreateMethodResponse'
      tags:
        - Rooms

components:
  schemas:
    DevicesV0GetMethodResponse:
      type: object
      properties:
        devices:
          type: array
          items:
            properties:
              name:
                type: string
                description: Название устройство
              type_id:
                type: string
                description: Тип устройства
              serial_number:
                type: string
                description: Номер девайса
              status:
                type: string
                description: Статус
              params:
                type: string
                description: Параметры
    ModuleV0GetMethodResponse:
      type: object
      properties:
        modules:
            type: array
            items:
              type: object
              properties:
                id:
                  type: string
                  description: ID модуля
                name:
                  type: string
                  description: Название модуля
                type:
                  type: string
                  description: Тип модуля
                status:
                  type: string
                  description: Тип модуля
                locations:
                  type: array
                  items:
                    properties:
                      location:
                        type: string
                        description: ID устройств
                devices:
                  type: array
                  items:
                   properties:
                      name:
                        type: string
                        description: Название устройство
                      type_id:
                        type: string
                        description: Тип устройства
                      serial_number:
                        type: string
                        description: Номер девайса
                      status:
                        type: string
                        description: Статус
                      params:
                        type: string
                        description: Параметры
    ModuleV0CreateMethodRequest:
      type: object
      properties:
        name:
         type: string
         description: Название модуля
        module_id:
          type: string
          description: Тип модуля
        status:
          type: string
          description: Статус устройства
        devices:
          type: array
          items:
            properties:
              device_id:
                type: string
                description: ID устройств
    ModuleV0CreateMethodResponse:
      type: object
      properties:
        id:
          type: string
          description: ID модуля
    ModuleV0UpdateMethodRequest:
       type: object
       properties:
                name:
                  type: string
                  description: Название модуля
                module_id:
                  type: string
                  description: Тип модуля
                status:
                  type: string
                  description: Статус устройства
                devices:
                  type: array
                  items:
                    properties:
                      device_id:
                        type: string
                        description: ID устройств
    ModuleDisableV0MethodResponse:
      type: object
      properties:
        id:
          type: string
          description: ID модуля
    RoomsV0GetMethodResponse:
      type: object
      properties:
        rooms:
          type: array
          items:
            type: object
            properties:
              id:
                type: string
                description: ID комнаты
              name:
                type: string
                description: Название комнаты
              status:
                type: string
                description: Статус
              module_ids:
                type: array
                items:
                  properties:
                    module_id:
                      type: string
                      description: ID модуля
    RoomsV0CreateMethodResponse:
      type: object
      properties:
              name:
                type: string
                description: Название комнаты
              status:
                type: string
                description: Статус
              module_ids:
                type: array
                items:
                  properties:
                    module_id:
                      type: string
                      description: ID модуля

```
