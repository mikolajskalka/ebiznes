# E-Biznes

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=mikolajskalka_ebiznes&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=mikolajskalka_ebiznes)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=mikolajskalka_ebiznes&metric=bugs)](https://sonarcloud.io/summary/new_code?id=mikolajskalka_ebiznes)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=mikolajskalka_ebiznes&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=mikolajskalka_ebiznes)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=mikolajskalka_ebiznes&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=mikolajskalka_ebiznes)


To repozytorium zawiera rozwiązania zadań, które zostały wyknane jako część kursu E-biznez na Uniwersytecie Jagiellońskim.

## Zadanie 1: Docker

### Wymagania

✅ 3.0 obraz Ubuntu z Pythonem w wersji 3.10 [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 3.5 obraz Ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 5.0 dodać konfigurację docker compose [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

**Kod:** [exercise1/](exercise1/)  
**Demo:** [Link to video](https://youtu.be/iuURFd7Obcg)

<details>
<summary>Szczegóły</summary>

- Ubuntu 24.04 obraz bazowy
- Java 8 (OpenJDK)
- Python 3.10
- Kotlin (zainstalowany za pomocą SDKMAN)
- Gradle 4.10.3

Kontener uruchamia prostą aplikację 'Hello World'.

### Uruchomienie

Aby zbudować i uruchomić aplikację:

```bash
cd exercise1
docker compose up
```

### Docker Image

Obraz Dockerowy wykonany w ramach zadania jest dostępny:
[mikolajskalka/java-hello-world-app:latest](https://hub.docker.com/repository/docker/mikolajskalka/java-hello-world-app/tags/latest/sha256-c5824510a94d5fdeedd1904e5ef0124b06fbc82af781cc287afa69949da041b3)

Obraz mozna pobrać bezpośrenio za pomocą komendy:

```bash
docker pull mikolajskalka/java-hello-world-app:latest
```
</details>

## Zadanie 2: Scala 

### Wymagania

✅ 3.0 Należy stworzyć kontroler do Produktów [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/042876bbf19eb93206001a15830ccfadb92b9614)

✅ 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane pobierane z listy [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/5d0143b636a8c70979afa045fb846453ff6bd6a1)

✅ 4.0 Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy zgodnie z CRUD [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/4ba6ce4895430f227875005a57084a6e0225d574)

✅ 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać skrypt uruchamiający aplikację via ngrok (nie podawać tokena ngroka w gotowym rozwiązaniu) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/a12303009787a739c76e3a655036dd2b87968d74)

✅ 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/a12303009787a739c76e3a655036dd2b87968d74)

**Kod:** [exercise2/](exercise2/)  
**Demo:** [Link to video](https://youtu.be/bG1TK95_NV0)

<details>
<summary>Szczegóły</summary>

Prosta aplikacja e-commerce napisana w Scali z użyciem frameworka Play. Aplikacja zawiera kontrolery do zarządzania produktami, kategoriami i koszykiem.

### Running the Solution

Aby zbudować i uruchomić aplikację:

```bash
cd exercise2
docker compose up
```

Aplikacja będzie dostępna pod adresem: http://localhost:9000

### Wystawiane endpointy API

Aplikacja wystawia następujące endpointy API:
- `GET /products` - Pobierz wszystkie produkty
- `GET /products/:id` - Pobierz produkt o danym ID
- `PUT /products/:id` - Zaktualizuj produkt o danym ID
- `POST /products` - Dodaj nowy produkt
- `DELETE /products/:id` - Usuń produkt o danym ID
- `GET /categories` - Pobierz wszystkie kategorie
- `GET /categories/:id` - Pobierz kategorię o danym ID
- `PUT /categories/:id` - Zaktualizuj kategorię o danym ID
- `POST /categories` - Dodaj nową kategorię
- `DELETE /categories/:id` - Usuń kategorię o danym ID
- `GET /cart/:id` - Pobierz zawartość koszyka od danym ID
- `PUT /cart/:id` - Zaktualizuj produkt w koszyku o danym ID
- `POST /cart` - Dodaj produkt do koszyka
- `DELETE /cart/:id` - Usuń koszyk o danym ID
</details>

## Zadanie 3: Kotlin

### Wymagania

✅ 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor, która pozwala na przesyłanie wiadomości na platformę Discord [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/948591fa81062a2fde763f87451700fc3a555b98)

✅ 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z platformy Discord skierowane do aplikacji (bota) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/948591fa81062a2fde763f87451700fc3a555b98)

✅ 4.0 Zwróci listę kategorii na określone żądanie użytkownika [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/948591fa81062a2fde763f87451700fc3a555b98)

✅ 4.5 Zwróci listę produktów wg żądanej kategorii [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/948591fa81062a2fde763f87451700fc3a555b98)

❌ 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack, Messenger, Webex 

**Kod:** [exercise3/](exercise3/)  
**Demo:** [Link to video](https://youtu.be/FQpDc7ekEng)

<details>
<summary>Szczegóły</summary>

Aplikacja została zbudowana w języku Kotlin z użyciem następujących technologii:
- Framework Ktor do tworzenia API RESTowych
- JDA (Java Discord API) do komunikacji z platformą Discord
- Docker do konteneryzacji aplikacji

### Funkcjonalności

Bot Discord obsługuje następujące komendy:
- `!categories` - Wyświetla listę wszystkich kategorii produktów
- `!products <category_id>` - Wyświetla produkty z wybranej kategorii
- `!help` - Wyświetla listę dostępnych komend

### Uruchomienie aplikacji

Aby uruchomić aplikację, należy ustawić odpowiednie zmienne środowiskowe:
- `DISCORD_TOKEN` - Token bota Discord

```bash
# Uruchomienie za pomocą Docker Compose
cd exercise3
docker compose up
```

Po uruchomieniu, aplikacja będzie dostępna pod adresem http://localhost:8080 i obsługuje następujące endpointy:
- `GET /health` - Sprawdzenie stanu aplikacji
- `GET /send-discord-message?channelId=<id>&message=<text>` - Wysłanie wiadomości na kanał Discord

</details>

## Zadanie 4: Go Echo

### Wymagania

✅ 3.0 Należy stworzyć aplikację we frameworki echo w j. Go, która będzie
miała kontroler Produktów zgodny z CRUD [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/2dd839d575203ab418d1fddc5e3752796b60c3c3)

✅ 3.5 Należy stworzyć model Produktów wykorzystując gorm oraz
wykorzystać model do obsługi produktów (CRUD) w kontrolerze (zamiast
listy) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/2dd839d575203ab418d1fddc5e3752796b60c3c3)

✅ 4.0 Należy dodać model Koszyka oraz dodać odpowiedni endpoint [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/2dd839d575203ab418d1fddc5e3752796b60c3c3)

✅ 4.5  Należy stworzyć model kategorii i dodać relację między kategorią,
a produktem [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/2dd839d575203ab418d1fddc5e3752796b60c3c3)

✅ 5.0 pogrupować zapytania w gorm’owe scope'y [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/2dd839d575203ab418d1fddc5e3752796b60c3c3)

**Kod:** [exercise4/](exercise4/)  
**Demo:** [Link to video](https://youtu.be/wOU12Hf8l-w)

<details>
<summary>Szczegóły</summary>

Aplikacja RESTowa zbudowana przy użyciu frameworka Echo w języku Go, wykorzystująca GORM do operacji bazodanowych oraz SQLite jako bazę danych.

### Struktura projektu

- **controllers/**: Zawiera handlery dla żądań HTTP
- **database/**: Zarządza połączeniem i inicjalizacją bazy danych
- **models/**: Definiuje modele danych z wykorzystaniem GORM
- **routes/**: Konfiguruje ścieżki aplikacji

### Modele

Aplikacja zawiera 5 modeli z relacjami:
1. **Product**: Główna encja produktu z funkcjonalnością CRUD
2. **Category**: Powiązana z produktami (relacja jeden-do-wielu)
3. **Cart**: Koszyk zakupowy dla użytkowników
4. **CartItem**: Elementy w koszyku (powiązane z produktami)
5. **User**: Użytkownik aplikacji, który posiada koszyki

### Endpointy API

#### Produkty
- `GET /products` - Pobierz wszystkie produkty
- `GET /products/:id` - Pobierz produkt po ID
- `POST /products` - Utwórz nowy produkt
- `PUT /products/:id` - Zaktualizuj istniejący produkt
- `DELETE /products/:id` - Usuń produkt
- `GET /products/category/:categoryId` - Pobierz produkty według kategorii
- `GET /products/price-range?min=X&max=Y` - Pobierz produkty w zakresie cenowym

#### Kategorie
- `GET /categories` - Pobierz wszystkie kategorie
- `GET /categories/with-products` - Pobierz wszystkie kategorie z ich produktami
- `GET /categories/:id` - Pobierz kategorię po ID
- `POST /categories` - Utwórz nową kategorię
- `PUT /categories/:id` - Zaktualizuj istniejącą kategorię
- `DELETE /categories/:id` - Usuń kategorię
- `GET /categories/search?name=X` - Wyszukaj kategorie po nazwie

#### Koszyki
- `GET /carts` - Pobierz wszystkie koszyki
- `GET /carts/:id` - Pobierz koszyk po ID
- `POST /carts` - Utwórz nowy koszyk
- `POST /carts/:id/items` - Dodaj element do koszyka
- `DELETE /carts/:id/items/:itemId` - Usuń element z koszyka
- `GET /carts/user/:userId` - Pobierz koszyki według ID użytkownika

### GORM Scopes

Aplikacja wykorzystuje GORM Scopes dla bardziej efektywnych zapytań do bazy danych:
- Aktywne rekordy (nie usunięte)
- Rekordy z załadowanymi powiązanymi encjami
- Filtrowanie według różnych kryteriów (zakres cenowy, kategoria, itp.)

### Uruchomienie aplikacji

#### Za pomocą Go
```bash
# Instalacja zależności
go mod download

# Uruchomienie aplikacji
go run main.go
```

#### Za pomocą Dockera
```bash
# Zbuduj i uruchom za pomocą Docker Compose
docker compose up --build
```

API będzie dostępne pod adresem http://localhost:8080
</details>

## Zadanie 5: Frontend React

### Wymagania

✅ 3.0 W ramach projektu należy stworzyć dwa komponenty: Produkty oraz Płatności; Płatności powinny wysyłać do aplikacji serwerowej dane, a w Produktach powinniśmy pobierać dane o produktach z aplikacji serwerowej [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/799eaa5c58a9b4abc9f8083ac95a0bc225ca1068)

✅ 3.5 Należy dodać Koszyk wraz z widokiem; należy wykorzystać routing [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/799eaa5c58a9b4abc9f8083ac95a0bc225ca1068)

✅ 4.0 Dane pomiędzy wszystkimi komponentami powinny być przesyłane za pomocą React hooks [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/799eaa5c58a9b4abc9f8083ac95a0bc225ca1068)

✅ 4.5 Należy dodać skrypt uruchamiający aplikację serwerową oraz kliencką na dockerze via docker compose [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/asdasdasdasd)

✅ 5.0 Należy wykorzystać axios'a oraz dodać nagłówki pod CORS [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/799eaa5c58a9b4abc9f8083ac95a0bc225ca1068)

**Kod:** [exercise5/](exercise5/)  
**Demo:** [Link to video](https://youtu.be/RnOisYnO8J0)

<details>
<summary>Szczegóły</summary>

Aplikacja kliencka zbudowana z wykorzystaniem React.js, Vite i React Router, komunikująca się z API serwerowym.

### Komponenty

1. **Produkty** - Komponent wyświetlający wszystkie produkty z możliwością filtrowania po kategoriach, pobierający dane z API serwerowego
2. **Koszyk** - Zarządzanie koszykiem zakupowym z możliwością dodawania, usuwania i aktualizowania ilości produktów
3. **Płatności** - Formularz realizacji płatności, wysyłający dane do serwera

### Technologie

- React 18 z użyciem Hooks (useState, useEffect, useContext, useReducer)
- React Router do nawigacji między komponentami
- Axios do komunikacji z API
- Vite jako narzędzie bundlujące
- Context API do zarządzania globalnym stanem aplikacji

### Uruchomienie aplikacji

#### Lokalne uruchomienie
```bash
cd exercise5
npm install
npm run dev
```

#### Za pomocą Dockera
```bash
cd exercise5
docker compose up --build
```

Aplikacja kliencka będzie dostępna pod adresem http://localhost:5173
</details>

## Zadanie 6: Testy

### Wymagania

✅ 3.0 Należy stworzyć 20 przypadków testowych w CypressJS lub Selenium [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/e26fa78f0349b774d827942fea45f51b62f612d6)

✅ 3.5 Należy rozszerzyć testy funkcjonalne, aby zawierały minimum 50 asercji [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/e26fa78f0349b774d827942fea45f51b62f612d6)

✅ 4.0 Należy stworzyć testy jednostkowe do wybranego wcześniejszego projektu z minimum 50 asercjami [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/e15490ae62d116c9845d151f6768c9b2afb2adff)

✅ 4.5 Należy dodać testy API, należy pokryć wszystkie endpointy z minimum jednym scenariuszem negatywnym per endpoint [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/e26fa78f0349b774d827942fea45f51b62f612d6)

❌ 5.0 Należy uruchomić testy funkcjonalne na Browserstacku

**Kod:** [exercise6/](exercise6/)  
**Demo:** [Link to video](https://youtu.be/8rq_GwxJX60)

<details>
<summary>Szczegóły</summary>

Kompletny zestaw testów dla aplikacji frontend i backend, wykorzystujący Cypress.

### Rodzaje testów

1. **Testy komponentów** - Testują poszczególne komponenty React w izolacji
2. **Testy E2E** - Testują cały przepływ pracy użytkownika od początku do końca
3. **Testy API** - Testują wszystkie endpointy API z uwzględnieniem scenariuszy pozytywnych i negatywnych

### Statystyki testów

- 20+ scenariuszy testowych
- 50+ asercji w testach funkcjonalnych
- 50+ asercji w testach jednostkowych
- Pełne pokrycie endpointów API z testami negatywnymi

### Uruchomienie testów

#### Testy lokalne
```bash
cd exercise6
npm install
npm run test:e2e         # Uruchomienie testów E2E
npm run test:component   # Uruchomienie testów komponentów
npm run test:api         # Uruchomienie testów API
```

</details>

## Zadanie 7: Sonar

### Wymagania

✅ 3.0 Należy dodać lintera do odpowiedniego kodu aplikacji serwerowej w hookach gita [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/cd70f1f754feecb65f176796f7b721a57b3ffaba)

✅ 3.5 Należy wyeliminować wszystkie bugi w kodzie w Sonarze (kod aplikacji serwerowej) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/ecf0b299ec422ce4ee4160109131c9bb0e2251e1)

✅ 4.0 Należy wyeliminować wszystkie zapaszki w kodzie w Sonarze (kod aplikacji serwerowej) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/ecf0b299ec422ce4ee4160109131c9bb0e2251e1)

✅ 4.5 Należy wyeliminować wszystkie podatności oraz błędy bezpieczeństwa w kodzie w Sonarze (kod aplikacji serwerowej) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/ecf0b299ec422ce4ee4160109131c9bb0e2251e1)

✅ 5.0 Należy wyeliminować wszystkie błędy oraz zapaszki w kodzie aplikacji klienckiej [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/ecf0b299ec422ce4ee4160109131c9bb0e2251e1)

**Kod:** [exercise7/](exercise7/)  
**Demo:** [Link to video](https://youtu.be/d8EQqWu7oGU)

<details>
<summary>Szczegóły</summary>

Integracja analizy kodu z SonarCloud dla aplikacji serwerowej i klienckiej, wraz z poprawkami wszystkich znalezionych problemów.

### Konfiguracja Sonara

- Dodane projekty aplikacji serwerowej i klienckiej do SonarCloud
- Skonfigurowane automatyczne skanowanie kodu przy każdym push do repozytorium

#### Wybrane reguły lintowania

- **gofmt** - Sprawdzanie formatowania kodu
- **goimports** - Sprawdzanie organizacji importów
- **govet** - Raportowanie podejrzanych konstrukcji
- **errcheck** - Wykrywanie nieobsługiwanych błędów
- **staticcheck** - Statyczna analiza kodu