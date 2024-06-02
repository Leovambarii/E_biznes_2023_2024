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

:white_check_mark: 5.0 - pogrupować zapytania w gorm’owe scope'y - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/c92bb96345de9881072a7adeb9754f906f46c2e8)

## **Zadanie 5** Frontend

### **Kod**: [5](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/5)

:white_check_mark: 3.0 - W ramach projektu należy stworzyć dwa komponenty: Produkty oraz Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach z aplikacji serwerowej - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 3.5 - Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 4.0 - Dane pomiędzy wszystkimi komponentami powinny być przesyłane za pomocą React hooks - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: 4.5 - Należy dodać skrypt uruchamiający aplikację serwerową oraz kliencką na dockerze via docker-compose - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/1cb2878dd190ee6d61b9f8c8391f3ca76d9c3e01)

:white_check_mark: - 5.0 - Należy wykorzystać axios’a oraz dodać nagłówki pod CORS - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/b4c266f448ec8b6949b06f83b923395c584e5ba4)

## **Zadanie 6** Testy

### **Kod**: [6](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/6)

:white_check_mark: 3.0 - Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium (Kotlin, Python, Java, JS, Go, Scala) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/640603936e3075341076a8aa104f53efbc5b2579)

:white_check_mark: 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50 asercji - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/640603936e3075341076a8aa104f53efbc5b2579)

:x: 4.0 - Należy stworzyć testy jednostkowe do wybranego wcześniejszego projektu z minimum 50 asercjami

:x: 4.5 - Należy dodać testy API, należy pokryć wszystkie endpointy z minimum jednym scenariuszem negatywnym per endpoint

:x: 5.0 - Należy uruchomić testy funkcjonalne na Browserstacku

## **Zadanie 7** Sonar

### **Kod**: [7](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/7)

:white_check_mark: 3.0 - Należy dodać litera do odpowiedniego kodu aplikacji serwerowej w hookach gita - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/853e2b788a333521920949ba0cf520fde92b848a)

:white_check_mark: 3.5 - Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji serwerowej) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/853e2b788a333521920949ba0cf520fde92b848a)

:white_check_mark: 4.0 - Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod aplikacji serwerowej) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/853e2b788a333521920949ba0cf520fde92b848a)

:white_check_mark: 4.5 - Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa w kodzie w Sonarze (kod aplikacji serwerowej) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/853e2b788a333521920949ba0cf520fde92b848a)

:white_check_mark: 5.0 - Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/ac0e1b5dc5037be980bc093ca79d20e0f3d39332)

## **Zadanie 8** Oauth2

### **Kod**: [8](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/8)

:white_check_mark: 3.0 - Logowanie przez aplikację serwerową (bez Oauth2) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9cf7db525a176e2cb8e54c45a2a54b59dd8a6065)

:white_check_mark: 3.5 - Rejestracja przez aplikację serwerową (bez Oauth2) - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9cf7db525a176e2cb8e54c45a2a54b59dd8a6065)

:x: 4.0 - Logowanie via Google OAuth2

:x: 4.5 - Logowanie via Facebook lub Github OAuth2

:x: 5.0 - Zapisywanie danych logowania OAuth2 po stronie serwera

## **Zadanie 9** ChatGPT/LLAMA bot

### **Kod**: [9](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/9)

:white_check_mark: 3.0 - Należy stworzyć po stronie serwerowej osobny serwis do łącznia z ChatGPT/LLAMA do usługi chat - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9cac3b11eb11314d09142c35b3538003f78b5ff2)

:white_check_mark: 3.5 - Należy stworzyć interfejs frontowy dla użytkownika, który komunikuje się z serwisem; odpowiedzi powinny być wysyłane do frontendowego interfejsu - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9cac3b11eb11314d09142c35b3538003f78b5ff2)

:white_check_mark: 4.0 - Stworzyć listę 5 różnych otwarć oraz zamknięć rozmowy - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/9cac3b11eb11314d09142c35b3538003f78b5ff2)

:x: 4.5 - Filtrowanie po zagadnieniach związanych ze sklepem (np. ograniczenie się jedynie do ubrań oraz samego sklepu) do GPT/LLAMA

:x: 5.0 - Filtrowanie odpowiedzi po sentymencie

## **Zadanie 10** Chmura/CI

### **Kod**: [10](https://github.com/Leovambarii/E_biznes_2023_2024/tree/main/10)

:white_check_mark: 3.0 - Należy stworzyć odpowiednie instancje po stronie chmury na Docker - [Commit](https://github.com/Leovambarii/E_biznes_2023_2024/commit/e106c70a5041aabd04b80c31d5abe7be85566a14)

:x: 3.5 - Stworzyć odpowiedni pipeline w Github Actions do budowania aplikacji (np. via fatjar)

:x: 4.0 - Dodać notyfikację mailową o zbudowaniu aplikacji

:x: 4.5 - Dodać krok z deploymentem aplikacji serwerowej oraz klienckiej na chmurę

:x: 5.0 - Dodać uruchomienie regresyjnych testów automatycznych (funkcjonalnych) jako krok w Actions