package main
 
import (
    "net/http"
    "strconv"
    "fmt"
	"html/template"

)
// Создается функция-обработчик "home", которая записывает байтовый слайс, содержащий
// текст "Привет из Snippetbox" как тело ответа.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
	// функция http.NotFound() для возвращения клиенту ошибки 404.
	// Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
	// продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
	if r.URL.Path != "/" {
		app.notFound(w)
		http.NotFound(w, r)
		return
	}
	// Инициализируем срез содержащий пути к двум файлам. Обратите внимание, что
	// файл home.page.tmpl должен быть *первым* файлом в срезе.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// Используем функцию template.ParseFiles() для чтения файлов шаблона.
	// Если возникла ошибка, мы запишем детальное сообщение ошибки и
	// используя функцию http.Error() мы отправим пользователю
	// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
	}
    //w.Write([]byte("Привет из Snippetbox"))
}
// Обработчик для отображения содержимого заметки.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
    // Извлекаем значение параметра id из URL и попытаемся
	// конвертировать строку в integer используя функцию strconv.Atoi(). Если его нельзя
	// конвертировать в integer, или значение меньше 1, возвращаем ответ
	// 404 - страница не найдена!
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Использование помощника notFound()
		return
	}
 
	// Используем функцию fmt.Fprintf() для вставки значения из id в строку ответа
	// и записываем его в http.ResponseWriter.
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

// Обработчик для создания новой заметки.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        // Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
        // карту HTTP-заголовков. Первый параметр - название заголовка, а
        // второй параметр - значение заголовка.
        w.Header().Set("Allow", http.MethodPost)
        
        // Если это не так, то вызывается метод w.WriteHeader() для возвращения статус-кода 405
        // и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
        // Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
        // w.WriteHeader(405)
        // w.Write([]byte("GET-Метод запрещен!"))

		app.clientError(w, http.StatusMethodNotAllowed) // Используем помощник clientError()
        return
    }
	w.Write([]byte("Форма для создания новой заметки..."))
}
 