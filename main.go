// Основной пакет Go-программы
package main

// Подключаем необходимые пакеты
import (
	"net/http" // стандартный пакет для работы с HTTP-статусами
	"time"     // стандартный пакет для работы с датами и временем

	"github.com/gin-gonic/gin" // внешний фреймворк Gin для упрощённой работы с HTTP-сервером
)

func main() {
	// Создаём новый роутер Gin с настройками по умолчанию (логирование, recovery и т.д.)
	r := gin.Default()

	// Определяем обработчик GET-запроса по маршруту "/info"
	r.GET("/info", func(c *gin.Context) {
		// Получаем текущую дату и время
		now := time.Now()

		// Создаём объект даты следующего 1 января в текущем временном поясе
		newYear := time.Date(
			now.Year()+1, // следующий год
			1,            // январь
			1,            // первый день месяца
			0, 0, 0, 0,   // часы, минуты, секунды, наносекунды
			now.Location(), // временная зона текущей даты
		)

		// Вычисляем количество дней до нового года
		// newYear.Sub(now) возвращает разницу во времени (Duration)
		// Hours() — количество часов в Duration
		// Делим на 24, чтобы получить количество дней и приводим к int
		days := int(newYear.Sub(now).Hours() / 24)

		// Отправляем JSON-ответ клиенту со статусом 200 OK
		c.JSON(http.StatusOK, gin.H{
			"days_before_new_year": days, // ключ: "days_before_new_year", значение: days
		})
	})

	// Новый маршрут для подсчёта дней до произвольной даты 
	r.GET("/countdown", countdownHandler)
	

	// Запускаем HTTP-сервер на порту 4200
	// Сервер будет слушать все интерфейсы (0.0.0.0:4200)
	r.Run(":4200")
}

// Handler for GET /countdown?date=YYYY-MM-DD
func countdownHandler(c *gin.Context) {
    // Получаем query-параметр
    targetDateStr := c.Query("date")
    
    // Проверяем, передан ли параметр
    if targetDateStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing required query parameter: date",
        })
        return
    }
    
    // Парсим дату в формате YYYY-MM-DD
    targetDate, err := time.Parse("2006-01-02", targetDateStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid date format. Use YYYY-MM-DD",
        })
        return
    }
    
    // Получаем текущую дату (без времени)
    now := time.Now()
    today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
    
    // Считаем разницу в днях
    daysRemaining := int(targetDate.Sub(today).Hours() / 24)
    
    // Возвращаем успешный ответ
    c.JSON(http.StatusOK, gin.H{
        "target_date":    targetDateStr,
        "days_remaining": daysRemaining,
    })
}
