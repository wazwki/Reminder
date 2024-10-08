### Архитектура на основе событий: Система управления задачами

#### Обзор
Вам нужно создать упрощенную систему управления задачами, используя архитектуру на основе событий. Система будет состоять из двух основных сервисов: Сервиса задач и Сервиса уведомлений. Сервис задач будет обрабатывать создание, обновление и удаление задач, в то время как Сервис уведомлений будет отправлять уведомления при возникновении этих событий.

#### Сервисы

1. **Сервис задач**:
    - Управляет задачами.
    - Публикует события при создании, обновлении или удалении задач.

2. **Сервис уведомлений**:
    - Слушает события задач.
    - Отправляет уведомления на основе полученных событий.

#### Требования

1. **Сервис задач**:
    - Должен предоставлять RESTful API для создания, обновления и удаления задач.
    - Должен публиковать события (`task_created`, `task_updated`, `task_deleted`) в брокер сообщений (Kafka).

2. **Сервис уведомлений**:
    - Должен слушать события задач из брокера сообщений.
    - Должен логировать сообщение при получении события ("уведомление").

#### Детали реализации

1. **Сервис задач**:
    - RESTful API эндпоинты:
        - `POST /tasks`: Создать новую задачу.
        - `PUT /tasks/{id}`: Обновить существующую задачу.
        - `DELETE /tasks/{id}`: Удалить задачу.
    - Использовать клиент брокера сообщений для публикации событий.

2. **Сервис уведомлений**:
    - Подписаться на события задач из брокера сообщений.
    - Логировать сообщение для каждого полученного события.

#### Технологии

- Язык программирования: Go, Python
- Брокер сообщений: Kafka
- HTTP сервер: встроенный пакет Go `net/http`, фреймворк Python 3 `FastAPI`
- Логирование: встроенный пакет Go `log` или `logging` для Python 3

#### Определения задач

Задача должна иметь следующие поля:

- ID (string)
- Название (string)
- Описание (string)

#### Шаги для выполнения

1. **Настройка брокера сообщений**:
    - Установить и настроить Kafka.

2. **Сервис задач**:
    - Реализовать RESTful API.
    - Реализовать публикацию событий в брокер сообщений.

3. **Сервис уведомлений**:
    - Реализовать подписку на события из брокера сообщений.
    - Реализовать логирование для полученных событий.

4. **Запуск и тестирование**:
    - Запустить оба сервиса.
    - Использовать Postman, для создания, обновления и удаления задач.
    - Проверить, что Сервис уведомлений логирует соответствующие сообщения для каждого события.

#### Пример выполнения

1. **Создание задачи**:
    - Отправить `POST /tasks` с данными задачи в Сервис задач.
    - Сервис задач создает задачу и публикует событие `task_created`.
    - Сервис уведомлений получает событие `task_created` и логирует сообщение.

2. **Обновление задачи**:
    - Отправить `PUT /tasks/{id}` с обновленными данными задачи.
    - Сервис задач обновляет задачу и публикует событие `task_updated`.
    - Сервис уведомлений получает событие `task_updated` и логирует сообщение.

3. **Удаление задачи**:
    - Отправить `DELETE /tasks/{id}`.
    - Сервис задач удаляет задачу и публикует событие `task_deleted`.
    - Сервис уведомлений получает событие `task_deleted` и логирует сообщение.