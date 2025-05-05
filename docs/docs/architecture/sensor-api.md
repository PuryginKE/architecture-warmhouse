```
openapi: 3.0.0
info:
  title: Sensor API
  description: Сервис - управлением датчиком
  version: 1.0.0
servers:
  - url: https://localhosh:8080
    description: dev стенд
tags:
  - name: Sensor
    description: Датчик

paths:
  /api/sensor/get:
    get:
      summary: Получение события
      description: Получение события от датчика
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SensorV0GetMethodResponse'
      tags:
        - Sensor
  /api/sensor/set:
    post:
      summary: Отправка события
      description: Отправка события в датчик
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/SensorV0CreateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/SensorV0GetMethodResponse'
      tags:
        - Sensor

components:
  schemas:
    SensorV0GetMethodResponse:
      type: object
      properties:
        device_id:
          type: string
          description: ID устройства
        value:
          type: string
          description: Значение
    SensorV0CreateMethodRequest:
      type: object
      properties:
        name:
          type: string
          description: Название сценария
        status:
          type: string
          description: Статус
        modules:
          type: array
          items:
            properties:
              modules_id:
                type: string
                description: ID модуля

```
