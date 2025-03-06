# Менеджер токенов и данных пользователя

Проект позволяет получать данные пользователя из VK, GitHub и Stepik с использованием токенов доступа. Данные сохраняются в текстовый файл `data.txt` в папке `./output/`.

## Структура проекта
```
cmd/
├── main.go
internal/
│ ├── APIs/
│ │ ├── github/
│ │ ├── github.go
│ │ ├── stepik/
│ │ ├── stepik.go
│ │ ├── vk/
│ │ └── vk.go
│ ├── server/
│ │ └── server.go
│ ├── public/
│ ├── templates/
│ ├── getTokens.html
│ └── resume.html
.gitignore
go.mod
README.md 
```
## Установка

**Убедитесь, что у вас установлен Go (версия 1.16 или выше).**

Клонируйте репозиторий:
   ```bash
   git clone https://git.miem.hse.ru/ps-biv24x/aisavelev.git
   cd aisavelev
   go mod tidy
   ```

## Запуск

   ```bash
   go run ./cmd/main.go
   ```
## Создание резюме

1. Перейдите по ссылке `http://localhost:8080/resume/generate`
2. Введите токены в поля на сайте
3. Вас автоматически перенаправит на сайт с готовым резюме

После этого вы сможете в любой момент перейти по ссылке `http://localhost:8080` и увидеть готовое резюме

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
