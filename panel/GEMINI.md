# GEMINI.md

## Project Overview

This project is a web application with a decoupled frontend and backend.

*   **Backend:** The backend is a Django application that serves a RESTful API using the Django Rest Framework. It uses a PostgreSQL database for data storage and Redis for caching and as a message broker for Celery. Asynchronous tasks are handled by Celery workers. The backend also integrates with AI services like OpenAI and Anthropic, and with Cloudflare.

*   **Frontend:** The frontend is a single-page application built with React and Vite. It uses Material-UI for its component library, Redux for state management, and React Router for navigation.

*   **Orchestration:** The application is containerized using Docker and the services are managed with `docker-compose`.

## Building and Running

The application is designed to be run with Docker.

1.  **Prerequisites:**
    *   Docker
    *   Docker Compose

2.  **Environment Variables:**
    *   Create a `.env` file in the root of the project and populate it with the necessary environment variables. You can use the `docker-compose.yml` file as a reference for the required variables.

3.  **Build and Run:**
    *   To build and start all the services, run the following command from the root of the project:
        ```bash
        docker-compose up --build
        ```
    *   To run the services in detached mode, use the `-d` flag:
        ```bash
        docker-compose up --build -d
        ```

4.  **Accessing the Application:**
    *   The frontend will be available at `http://localhost:3000` (or the port specified in your `.env` file).
    *   The backend API will be available at `http://localhost:8000` (or the port specified in your `.env` file).

## Development Conventions

### Backend

*   The backend code is located in the `backend` directory.
*   Python dependencies are managed with `pip` and are listed in `backend/requirements.txt`.
*   The backend follows the standard Django project structure.
*   API endpoints are defined in the `views.py` file of each app and the routing is configured in `urls.py`.
*   Database models are defined in the `models.py` file of each app.
*   Serializers for the API are defined in the `serializers.py` file of each app.
*   Asynchronous tasks are defined in the `tasks.py` file of each app.

### Frontend

*   The frontend code is located in the `frontend` directory.
*   JavaScript dependencies are managed with `npm` and are listed in `frontend/package.json`.
*   The frontend is a React application built with Vite.
*   Components are located in `frontend/src/components`.
*   Pages are located in `frontend/src/pages`.
*   The Redux store is configured in `frontend/src/store`.
*   The application's routes are defined in `frontend/src/App.tsx`.

### Testing

*   **Backend:**
    *   Tests are located in the `tests.py` file of each app.
    *   To run the backend tests, you can run the following command:
        ```bash
        docker-compose exec backend pytest
        ```
*   **Frontend:**
    *   TODO: Add instructions for running frontend tests.
