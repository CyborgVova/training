### Идиоматичные паттерны многопоточности

- For-Select-Done Позволяет осуществить выход из горутины для исключения ее утечки
- Fan-In Слияние произвольного количества каналов в один того же типа
- Take-First-N Позволяет обрабатывать количество данных ограниченное
  количеством переданным в качестве параметра
- Subscription Подписчик, паттерн получает с определенной периодичностью
  какие-то данные
- Mapping Позволяет применять к данным какое-то действие через вызов
  функции переданной в качестве параметра
- Filter Позволяет отфильтровать поступающие данные по признаку
  применяемому в функции переданной в качестве параметра
