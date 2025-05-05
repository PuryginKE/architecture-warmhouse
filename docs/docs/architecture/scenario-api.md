```
openapi: 3.0.0
info:
  title: Scenario API
  description: Сервис - управлением сценариями
  version: 1.0.0
servers:
  - url: https://localhosh:8080
    description: dev стенд
tags:
  - name: Scenario
    description: Сценарий

paths:
  /api/scenario/v0/get:
    get:
      summary: Получение списка сценариев
      description: Получение списка сценариев
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ScenarioV0GetMethodResponse'
      tags:
        - Scenario
  /api/scenario/v0/create:
    post:
      summary: Добавление сценария
      description: Получение сценария дома.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/ScenarioV0CreateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ScenarioV0CreateMethodResponse'
      tags:
        - Scenario
  /api/scenario/v0/update/{id}:
    post:
      summary: Обновление сценария
      description: Обновление сценария по id
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
              $ref: '#/components/schemas/ScenarioV0CreateMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ScenarioV0CreateMethodResponse'
      tags:
        - Scenario
  /api/scenario/v0/delete/{id}:
    delete:
      summary: Удаление сценария
      description: Удаление сценария
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
                $ref: '#/components/schemas/ScenarioV0CreateMethodResponse'
      tags:
        - Scenario
  /api/scenario/v0/get/module/{id}:
    get:
      summary: Получение события по модулю
      description: Получение события по модулю
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
                $ref: '#/components/schemas/ScenarioV0GetModuleMethodResponse'
      tags:
        - Scenario
  /api/scenario/v0/post/module/{id}:
    post:
      summary: Отправка события в модуль
      description: Отправка события в модуль
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
              $ref: '#/components/schemas/ScenarioV0CreateModuleMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/ScenarioV0CreateMethodResponse'
      tags:
        - Scenario

components:
  schemas:
    ScenarioV0GetMethodResponse:
      type: object
      properties:
        scenarios:
          type: array
          items:
            properties:
              id:
                type: string
                description: ID сценария
              name:
                type: string
                description: Название сценария
              status:
                type: string
                description: Статус
              dt_create:
                type: string
                description: Дата создания
              dt_update:
                type: string
                description: Дата обновления
              modules:
                  type: array
                  items:
                    properties:
                      modulesInfo:
                        type: string
                        description: Объект модуля
    ScenarioV0CreateMethodRequest:
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
    ScenarioV0CreateMethodResponse:
      type: object
      properties:
        id:
          type: string
          description: ID модуля
    ScenarioV0GetModuleMethodResponse:
      type: object
      properties:
        modules_id:
          type: string
          description: ID модуля
        value:
          type: string
          description: Значение
        status:
          type: string
          description: Статус
    ScenarioV0CreateModuleMethodRequest:
      type: object
      properties:
        modules:
          type: array
          items:
            properties:
              modules_id:
                type: string
                description: ID модуля
              value:
                type: string
                description: Значение
              status:
                type: string
                description: Статус

```
