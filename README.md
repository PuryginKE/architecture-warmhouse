# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удалённо включать/выключать отопление в своих домах
- Система поддерживает управление датчиками, регистрация/удаление/просмотр датчиков

**Мониторинг температуры:**

- Пользователи могут просматривать текущую температуру в своих домах через веб-интерфейс
- Система поддерживает просмотр комнатной температуры

### 2. Анализ архитектуры монолитного приложения

- Архитектура приложения представляет из себя монолит на Go с СУБД Postgres.
- Всё синхронно.
- Никаких асинхронных вызовов, микросервисов и реактивного взаимодействия в системе нет.
- Всё управление идёт от сервера к датчику.
- Данные о температуре также получаются через запрос от сервера к датчику.

### 3. Определение доменов и границы контекстов

- Домен: Управление умным домом
  - Поддомен: Управление датчиками системы отопления
    - Контекст: добавление(регистрация) датчика системы отопления
    - Контекст: удаление датчика системы отопления
    - Контекст: просмотр датчиков системы отопления
  - Поддомен: Управление отоплением
    - Контекст: изменение отопления всех комнат
    - Контекст: изменение отопления одной комнаты
  - Поддомен: Мониторинг температуры
    - Контекст: просмотр комнатной температуры

### **4. Проблемы монолитного решения**

- Развертывание и обновление
- Масштабируемость
- Устойчивость к отказам
- Производительность
- Риск изменений

### 5. Визуализация контекста системы — диаграмма С4

```markdown
[Диаграмма контекста](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/context/Context.puml)
```

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

```markdown
[Диаграмма контейнеров](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/container/Container.puml)
```

**Диаграмма компонентов (Components)**

```markdown
[Диаграмма компонентов Моб.приложения](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/component/Component.puml)
```

**Диаграмма кода (Code)**

```markdown
[Диаграмма кода Auth](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/code/Code-auth.puml)
```

```markdown
[Диаграмма кода Gateway](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/code/Code-gateway.puml)
```

```markdown
[Диаграмма кода Home](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/code/Code-home.puml)
```

```markdown
[Диаграмма кода Device](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/code/Code-device.puml)
```

```markdown
[Документация кода Scenario](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/code/Code-scenario.puml)
```

# Задание 3. Разработка ER-диаграммы

```markdown
[Диаграмма ER](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/diagrams/er/Code-er.puml)
```

# Задание 4. Создание и документирование API

### 1. Тип API

Используются:

- REST API для получение синхронной информации (прим. получение текущих настроек)
- AsyncAPI для чтения событий из кафки

### 2. Документация API

Здесь приложите ссылки на документацию API для микросервисов, которые вы спроектировали в первой части проектной работы. Для документирования используйте Swagger/OpenAPI или AsyncAPI.

```markdown
[Документация Scenario Api](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/docs/docs/architecture/scenario-api.md)
```

```markdown
[Документация Home Api](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/docs/docs/architecture/home-api.md)
```

```markdown
[Документация Sensor Api](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/docs/docs/architecture/sensor-api.md)
```

```markdown
[Документация Device Api](https://github.com/PuryginKE/architecture-warmhouse/blob/warmhouse/docs/docs/architecture/device-api.md)
```

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1. сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2. Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3. Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.
