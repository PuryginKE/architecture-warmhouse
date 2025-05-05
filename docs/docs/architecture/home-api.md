```
openapi: 3.0.0
info:
  title: Home API
  description: Сервис - информации по дому
  version: 1.0.0
servers:
  - url: https://localhosh:8080
    description: dev стенд
tags:
  - name: Home
    description: Дом

paths:
  /api/home/v0/get:
    get:
      summary: Получение информации о доме
      description: Получение информации о доме
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/HomeV0GetMethodResponse'
      tags:
        - Home
  /api/home/v0/get/rooms:
    get:
      summary: Получение комнат
      description: Получение информации о доме
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/HomeV0GetRoomsMethodResponse'
      tags:
        - Home
  /api/home/v0/get/room/{id}:
    get:
      summary: Получение информации по комнате
      description: Получение информации по комнате
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
                $ref: '#/components/schemas/HomeV0GetRoomMethodResponse'
      tags:
        - Home
  /api/home/v0/create/room:
    post:
      summary: Добавление комнаты
      description: Добавление комнаты в дом
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/HomeV0CreateRoomMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/HomeV0CreateRoomMethodResponse'
      tags:
        - Home
  /api/home/v0/update/room/{id}:
    post:
      summary: Обновление комнаты
      description: Обновление комнаты по id
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
              $ref: '#/components/schemas/HomeV0CreateRoomMethodRequest'
      responses:
        '200':
          description: Успешный ответ
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/HomeV0CreateRoomMethodResponse'
      tags:
        - Home
  /api/home/v0/delete/room/{id}:
    delete:
      summary: Удаление комнаты
      description: Удаление комнаты
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
                $ref: '#/components/schemas/HomeV0CreateRoomMethodResponse'
      tags:
        - Home

components:
  schemas:
    HomeV0GetMethodResponse:
      type: object
      properties:
        name:
          type: string
          description: Название дома
        rooms:
          type: array
          items:
            properties:
              id:
                type: string
                description: ID комнаты
              location:
                type: string
                description: локация
              temperature:
                type: string
                description: Температура
              lighting:
                type: string
                description: Свет
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
    HomeV0GetRoomsMethodResponse:
      type: object
      properties:
        rooms:
          type: array
          items:
            properties:
              id:
                type: string
                description: ID комнаты
              location:
                type: string
                description: локация
              temperature:
                type: string
                description: Температура
              lighting:
                type: string
                description: Свет
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
    HomeV0GetRoomMethodResponse:
      type: object
      properties:
              id:
                type: string
                description: ID комнаты
              location:
                type: string
                description: локация
              temperature:
                type: string
                description: Температура
              lighting:
                type: string
                description: Свет
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
    HomeV0CreateRoomMethodRequest:
      type: object
      properties:
              location:
                type: string
                description: локация
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
    HomeV0CreateRoomMethodResponse:
      type: object
      properties:
        id:
          type: string
          description: ID комнаты
```
