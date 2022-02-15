# Auth app

## Running the App Locally

### Prerequisites

- [Docker](https://docs.docker.com/install/)
- [Git](https://git-scm.com/downloads)

### Building

```
# Build the auth app for development

make build-local
```

### Running

```
make run-local
```

**Note:** `.env` file should be created before all of this.

### Building specific stages

```
# Development stage.

make build-dev
```

```
# Production stage.

make build-prod
```
