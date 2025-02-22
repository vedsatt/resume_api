```markdown
# Менеджер токенов и данных пользователя

Проект позволяет получать данные пользователя из VK, GitHub и Stepik с использованием токенов доступа. Данные сохраняются в текстовый файл `user_data.txt` в удобочитаемом формате.

## Установка

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Установите зависимости:
   ```bash
   go get fyne.io/fyne/v2
   ```
3. Клонируйте репозиторий:
   ```bash
   git clone https://git.miem.hse.ru/ps-biv24x/aisavelev.git
   cd ваш-репозиторий
   ```

## Запуск

1. Запустите программу:
   ```bashВот обновленный файл `README.md` с добавленным разделом о том, как вводить токены в графическом интерфейсе:

---

```markdown
# Менеджер токенов и данных пользователя

Проект позволяет получать данные пользователя из VK, GitHub и Stepik с использованием токенов доступа. Данные сохраняются в текстовый файл `user_data.txt` в удобочитаемом формате.

## Установка

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Установите зависимости:
   ```bash
   go get fyne.io/fyne/v2
   ```
3. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/ваш-username/ваш-репозиторий.git
   cd ваш-репозиторий
   ```

## Запуск

1. Запустите программу:
   ```bash
   go run main.go
   ```
2. Введите токены для VK, GitHub и Stepik в графическом интерфейсе.
3. После успешного ввода данные сохранятся в файл `user_data.txt`.

## Как вводить токены

1. После запуска программы откроется окно с полями для ввода токенов

2. После ввода всех токенов нажмите кнопку **"Сохранить"**.

3. Закройте окно, и данные будут сохранены в файл `user_data.txt`.

## Получение токенов

### VK
1. Перейдите по ссылке для получения токена:
   ```
   https://oauth.vk.com/authorize?client_id=ВАШ_CLIENT_ID&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=users&response_type=token&v=5.131
   ```
2. Скопируйте токен из параметра `access_token` в URL после авторизации.

### GitHub
1. Создайте Personal Access Token:
   - Перейдите в [настройки GitHub](https://github.com/settings/tokens).
   - Нажмите **Generate new token**.
   - Выберите разрешения: `read:user`, `repo`.
2. Скопируйте сгенерированный токен.

### Stepik
1. Создайте приложение на [Stepik OAuth](https://stepik.org/oauth2/applications/).
2. Получите `code` через URL:
   ```
   https://stepik.org/oauth2/authorize/?response_type=code&client_id=ВАШ_CLIENT_ID&redirect_uri=ВАШ_REDIRECT_URI
   ```
3. Обменяйте `code` на токен:
   ```bash
   curl -X POST -d "grant_type=authorization_code&code=ВАШ_CODE&client_id=ВАШ_CLIENT_ID&client_secret=ВАШ_CLIENT_SECRET&redirect_uri=ВАШ_REDIRECT_URI" https://stepik.org/oauth2/token/
   ```

## Получение токенов

### VK
1. Перейдите по ссылке для получения токена:
   ```
   https://oauth.vk.com/authorize?client_id=ВАШ_CLIENT_ID&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=users&response_type=token&v=5.131
   ```
2. Скопируйте токен из параметра `access_token` в URL после авторизации.

### GitHub
1. Создайте Personal Access Token:
   - Перейдите в [настройки GitHub](https://github.com/settings/tokens).
   - Нажмите **Generate new token**.
   - Выберите разрешения: `read:user`, `repo`.
2. Скопируйте сгенерированный токен.

### Stepik
1. Создайте приложение на [Stepik OAuth](https://stepik.org/oauth2/applications/).
2. Получите `code` через URL:
   ```
   https://stepik.org/oauth2/authorize/?response_type=code&client_id=ВАШ_CLIENT_ID&redirect_uri=ВАШ_REDIRECT_URI
   ```
3. Обменяйте `code` на токен:
   ```bash
   curl -X POST -d "grant_type=authorization_code&code=ВАШ_CODE&client_id=ВАШ_CLIENT_ID&client_secret=ВАШ_CLIENT_SECRET&redirect_uri=ВАШ_REDIRECT_URI" https://stepik.org/oauth2/token/
   ```