```markdown
# Менеджер токенов и данных пользователя

Проект позволяет получать данные пользователя из VK, GitHub и Stepik с использованием токенов доступа. Данные сохраняются в текстовый файл `data.txt` в папке `./output/`.

## Установка

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше).
2. Установите зависимости:
   ```bash
   go get fyne.io/fyne/v2
   ```
3. Клонируйте репозиторий:
   ```bash
   git clone https://git.miem.hse.ru/ps-biv24x/aisavelev.git
   cd скопированный-репозиторий
   ```

## Запуск

1. Запустите программу:
   ```bash
   go run ./cmd/main.go
   ```
2. Введите токены для VK, GitHub и Stepik (для Stepik также понадобится айди) в графическом интерфейсе.
3. После успешного ввода данные сохранятся в файл `data.txt`.

## Как вводить токены

1. После запуска программы откроется окно с полями для ввода токенов

2. После ввода всех токенов нажмите кнопку **"Сохранить"**.

3. **Обязательно закройте окно**, и данные будут сохранены в файл `data.txt`.

## Получение токенов

### VK
1. Перейдите по ссылке:
   ```
   https://vkhost.github.io
   ```
2. Выберите поле vk.com
3. Разрешите доступ
4. Вас перенаправит на сайт:
   ```
   https://oauth.vk.com/blank.html#access_token=XXX&expires_in=XXX&user_id=XXX
   ```
5. **ХХХ** после поля **access_token=** и есть ваш персональный токен
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