# E-Biznes

To repozytorium zawiera rozwiązania zadań, które zostały wyknane jako część kursu E-biznez na Uniwersytecie Jagiellońskim.

## Zadanie 1: Docker

### Wymagania

✅ 3.0 obraz Ubuntu z Pythonem w wersji 3.10 [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 3.5 obraz Ubuntu:24.02 z Javą w wersji 8 oraz Kotlinem [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC SQLite w ramach projektu na Gradle (build.gradle) [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji przez CMD oraz gradle [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

✅ 5.0 dodać konfigurację docker-compose [Link do commita](https://github.com/mikolajskalka/ebiznes/commit/38d83e5f86e7584c5ff8656dd642ae2a4bdecda8)

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
**Demo:** [Link to video](#)

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

<details>
<summary>Szczegóły</summary>

Wkrótce.
</details>