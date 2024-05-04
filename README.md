# Ebiznes_2023_2024
## **Zadanie 1** Docker

### Należy stworzyć obraz oraz umieścić obraz na hub.docker.com, a Dockerfile na githubie

### **Kod**: [1](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/1)

:white_check_mark: 3.0 - Obraz ubuntu z Pythonem w wersji 3.8 - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9eb6aded5780b876c9a35e9a37ceab3a6c920abc)

:white_check_mark: 3.5 - Obraz ubuntu:22.04 z Javą w wersji 8 oraz Kotlinem - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9eb6aded5780b876c9a35e9a37ceab3a6c920abc)

:white_check_mark: 4.0 - Do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9eb6aded5780b876c9a35e9a37ceab3a6c920abc)

:white_check_mark: 4.5 - Należy stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9eb6aded5780b876c9a35e9a37ceab3a6c920abc)

:white_check_mark: 5.0 - Należy dodać konfigurację docker-compose - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9eb6aded5780b876c9a35e9a37ceab3a6c920abc)

## **Zadanie 2** Scala

### Należy stworzyć aplikację na frameworku Play w Scali 2

### **Kod**: [2](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/2/ScalaProject)

:white_check_mark: 3.0 - Należy stworzyć kontroler do Produktów - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/22eba476a524cb29661eab9ecbd0c5794ebc6ed9)

:white_check_mark: 3.5 - Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/22eba476a524cb29661eab9ecbd0c5794ebc6ed9)

:x: 4.0 - Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD

:x: 4.5 - Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok

:x: 5.0 - Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD

## **Zadanie 3** Kotlin

### **Kod**: [3](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/3)

:white_check_mark: 3.0 - Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/23d68b1975fd40bf206c2cb18ad62a36e0b23cc2)

:white_check_mark: 3.5 - Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/23d68b1975fd40bf206c2cb18ad62a36e0b23cc2)

:white_check_mark: 4.0 - Zwróci listę kategorii na określone żądanie użytkownika - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/23d68b1975fd40bf206c2cb18ad62a36e0b23cc2)

:white_check_mark: 4.5 - Zwróci listę produktów wg żądanej kategorii - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/23d68b1975fd40bf206c2cb18ad62a36e0b23cc2)

:x: - 5.0 - Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex, Skype

## **Zadanie 4** Go

### **Kod**: [4](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/4)

:white_check_mark: 3.0 - Należy stworzyć aplikację we frameworki echo w j. Go, która będzie miała kontroler Produktów zgodny z CRUD - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/d42cef195a7c44beb8dfe5cb1e99a18f9ffe7caf)

:white_check_mark: 3.5 - Należy stworzyć model Produktów wykorzystując gorm oraz wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast listy) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/d42cef195a7c44beb8dfe5cb1e99a18f9ffe7caf)

:white_check_mark: 4.0 - Należy dodać model Koszyka oraz dodać odpowiedni endpoint - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/d42cef195a7c44beb8dfe5cb1e99a18f9ffe7caf)

:white_check_mark: 4.5 - Należy stworzyć model kategorii i dodać relację między kategorią, a produktem - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/d42cef195a7c44beb8dfe5cb1e99a18f9ffe7caf)

:x: - 5.0 - pogrupować zapytania w gorm’owe scope'y

## **Zadanie 5** Frontend

### **Kod**: [5](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/5)

:white_check_mark: 3.0 - W ramach projektu należy stworzyć dwa komponenty: Produkty oraz Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach z aplikacji serwerowej - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 3.5 - Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 4.0 - Dane pomiędzy wszystkimi komponentami powinny być przesyłane za pomocą React hooks - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 4.5 - Należy dodać skrypt uruchamiający aplikację serwerową oraz kliencką na dockerze via docker-compose - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:x: - 5.0 - Należy wykorzystać axios’a oraz dodać nagłówki pod CORS

## **Zadanie 6** Testy

### **Kod**: [6](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/6)

:white_check_mark: 3.0 - Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium (Kotlin, Python, Java, JS, Go, Scala) - [Commit]()

:white_check_mark: 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50 asercji - [Commit]()

:x: 4.0 - Należy stworzyć testy jednostkowe do wybranego wcześniejszego projektu z minimum 50 asercjami

:x: 4.5 - Należy dodać testy API, należy pokryć wszystkie endpointy z minimum jednym scenariuszem negatywnym per endpoint

:x: 5.0 - Należy uruchomić testy funkcjonalne na Browserstacku