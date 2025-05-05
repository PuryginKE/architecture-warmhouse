```
openapi: 3.0.0
info:
  title: Device API
  description: Сервис - управлением устройствами
  version: 1.0.0
servers:
  - url: https://localhosh:8080
    description: dev стенд
tags:
  - name: Device
    description: Устройства

paths:
  /api/device/v0/get:
    get:
      summary: Получение списка устройств
      description: Получение списка устройств
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/DeviceV0GetMethodResponse'
      tags:
        - Device
  /api/device/v0/create:
    post:
      summary: Добавление устройства
      description: Добавление устройства дома.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/DeviceV0CreateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/DeviceV0CreateMethodResponse'
      tags:
        - Device
  /api/device/v0/update/{id}:
    post:
      summary: Обновление устройства
      description: Обновление устройства по id
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
              $ref: '#/components/schemas/DeviceV0CreateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/DeviceV0CreateMethodResponse'
      tags:
        - Device
  /api/device/v0/delete/{id}:
    delete:
      summary: Удаление устройства
      description: Удаление устройства дома
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
                $ref: '#/components/schemas/DeviceV0CreateMethodResponse'
      tags:
        - Device

components:
  schemas:
    DeviceV0GetMethodResponse:
      type: object
      properties:
        devices:
          type: array
          items:
            properties:
              id:
                type: string
                description: ID устройства
              name:
                type: string
                description: Название устройства
              type_id:
                type: string
                description: Тип
              serial_number:
                type: string
                description: GUID Номер
              params:
                type: string
                description: Значение
              status:
                type: string
                description: Статус
    DeviceV0CreateMethodRequest:
      type: object
      properties:
        name:
          type: string
          description: Название устройства
        type_id:
          type: string
          description: Тип
        serial_number:
          type: string
          description: GUID Номер
        params:
          type: string
          description: Значение
        status:
          type: string
          description: Статус

    DeviceV0CreateMethodResponse:
      type: object
      properties:
        id:
          type: string
          description: ID устройства
```
