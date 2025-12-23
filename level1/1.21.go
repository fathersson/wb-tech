package main

import "fmt"

// --- 1. Существующая инфраструктура ---
// jsonDocument - это наш тип данных, который мы имеем в приложении
// Его интерфейс несовместим с сервисом аналитики
// convertToXml - метод, который превращает JSON в XML строку

type jsonDocument struct {
	content string
}

func (doc *jsonDocument) convertToXml() string {
	return fmt.Sprintf("<xml>%s</xml>", doc.content)
}

// --- 2. Целевой интерфейс (что ожидает клиент) ---
// AnaliticalService - интерфейс внешнего сервиса (потребителя)
// Мы не можем его менять, и он требует наличия метода sendXmlData
// xmlDocument - стандартный объект, который уже умеет работать с сервисом

type AnaliticalService interface {
	sendXmlData()
}

type xmlDocument struct{}

func (doc xmlDocument) sendXmlData() {
	fmt.Println("Отправка стандартного XML документа")
}

// --- 3. Адаптер ---
// jsonDocumentAdapter - мост между jsonDocument и интерфейсом AnaliticalService
// Он содержит внутри себя ссылку на адаптируемый объект (композиция)
// sendXmlData - реализация метода интерфейса AnaliticalService
// Адаптер получает вызов, подготавливает данные (конвертирует) и выполняет действие
// newJsonAdapter - это конструктор, который создает адаптер, возвращая интерфейс AnaliticalService
// За счет этого клиент може даже не знать, что он работает с адаптером

type jsonDocumentAdapter struct {
	jsonDocument *jsonDocument
}

func (adapter *jsonDocumentAdapter) sendXmlData() {
	// Делегируем работу внутреннему объекту
	xmlData := adapter.jsonDocument.convertToXml()
	fmt.Printf("Адаптер: данные преобразованы и отправлены: %s\n", xmlData)
}

func newJsonAdapter(doc *jsonDocument) AnaliticalService {
	return &jsonDocumentAdapter{
		jsonDocument: doc,
	}
}

func main() {
	// Создаем наш JSON документ с какими-то данными
	myJson := &jsonDocument{content: "данные пользователя"}

	// Напрямую мы не можем передать myJson туда, где нужен AnaliticalService
	// Поэтому используем адаптер
	service := newJsonAdapter(myJson)

	// Теперь мы вызываем метод интерфейса
	// Клиент (main) думает, что просто отправляет данные в аналитику,
	// а адаптер внутри делает конвертацию
	fmt.Println("Клиент: Запускаю отправку данных...")
	service.sendXmlData()
}
