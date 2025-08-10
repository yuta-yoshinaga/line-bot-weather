# GEMINI.md

This file provides guidance to GEMINI when working with code in this repository.

## Project Overview

This is a Go-based chatbot for the LINE messaging platform. Its purpose is to receive a location (prefecture or current location) and return the weather forecast for that area.

The project is deployed on Render, with separate environments for development and production.

## Code Architecture

The codebase is structured following the principles of **Clean Architecture**. This separates concerns into distinct layers, making the code more modular, testable, and maintainable.

*   **`entities`**: This layer contains the core domain models of the application, such as `Weather`, `Main` (for temperature), `Wind`, and `Rain`. These are plain Go structs with no dependencies on other parts of the application.
*   **`usecases`**: This layer contains the application-specific business logic. The `WeatherInteractor` is the primary component here. It orchestrates the data flow by using repository interfaces to fetch data and presenter interfaces to format it.
*   **`interfaceadapters`**: This layer contains controllers and presenters.
    *   `controllers`: Handle input from the web framework, convert it into a format the `usecases` layer can understand, and call the appropriate interactor.
    *   `presenters`: Take the output from the `usecases` and format it for the user (e.g., creating the final text message).
*   **`frameworksanddrivers`**: This layer contains concrete implementations of external dependencies. The `OpenWeatherMapRepository` is an example, which handles the specifics of making HTTP calls to the OpenWeatherMap API. It implements the repository interface defined in the `usecases` layer.

Dependency Injection is used to connect these layers. For instance, a `WeatherRepository` implementation is injected into the `WeatherInteractor`, which is then injected into the `WeatherController`.

## Development

### Environment Setup

For local development and testing, the application loads environment variables from a `.env` file in the project root using the `godotenv` package. Key variables include:
*   `WEATHER_URL`: The base URL for the OpenWeatherMap API.
*   `APPID`: Your API key for the OpenWeatherMap API.

### Common Commands

*   **Run all tests:**
    ```sh
    go test ./...
    ```

*   **Run a single test function:** Use the `-run` flag with a regular expression matching the test name.
    ```sh
    # Example for a test in the 'usecases' package
    go test -v -run ^TestWeatherInteractor_GetCurrentWeather$ ./usecases
    ```

*   **Build the application:**
    ```sh
    go build ./...
    ```

*   **Static Analysis:**
    The project uses GitHub CodeQL for security analysis, as configured in `.github/workflows/codeql-analysis.yml`. There is no separate linter currently configured.