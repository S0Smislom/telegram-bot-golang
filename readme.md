# Telegram Bot (Golang)

Небольшой Telegram-бот, написанный на **Golang**, с использованием:
- [`cobra-cli`](https://github.com/spf13/cobra) — для организации CLI команд
- [`telegram-bot-api`](https://github.com/go-telegram-bot-api/telegram-bot-api) — для работы с Telegram Bot API

## 🚀 Запуск проекта

Перед запуском:
1. Создайте бота в Telegram через [@BotFather](https://t.me/BotFather)
2. Получите токен и сохраните его в переменной окружения или в файле `.env`

### Локальный запуск (Go)
```bash
go build -o bot .
./bot
```

### Запуск с hot-reload (Air)
Если у вас установлен [`air`](https://github.com/air-verse/air):
```bash
air
```

### Запуск в docker
```bash
docker compose up
```