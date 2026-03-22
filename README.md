# cache-redis-config
======================

## Description
---------------

cache-redis-config is a lightweight Redis configuration manager for caching purposes. It simplifies the process of setting up and managing Redis cache configurations, allowing developers to focus on building their applications without worrying about the underlying cache infrastructure.

## Features
------------

*   **Easy Configuration**: Simplified configuration management for Redis cache setup
*   **Flexible Settings**: Comprehensive set of settings to customize cache behavior
*   **Robust Validation**: Thorough validation of configuration to prevent errors
*   **Automatic Initialization**: Initializes Redis cache connection on application startup
*   **Cache Statistics**: Exposes cache hit/miss statistics for performance analysis

## Technologies Used
---------------------

*   **Node.js**: JavaScript runtime environment
*   **Redis**: In-memory data store
*   **ESLint**: Static code analysis tool
*   **Jest**: JavaScript testing framework

## Installation
-------------

### Prerequisites

*   **Node.js**: v14.x or higher
*   **Redis**: v6.x or higher
*   **npm**: v6.x or higher

### Installation Steps

1.  Clone the repository: `git clone https://github.com/your-username/cache-redis-config.git`
2.  Navigate to the project directory: `cd cache-redis-config`
3.  Install dependencies: `npm install`
4.  Run the application: `npm start`

### Configuration

*   Create a `config.json` file in the project root with your Redis connection details
*   Update the `redis` object with your desired configuration settings

### Example Configuration
```json
{
    "redis": {
        "host": "localhost",
        "port": 6379,
        "password": "",
        "db": 0,
        "ttl": 3600,
        "maxConnections": 100,
        "maxIdleConnections": 20
    }
}
```
### API Endpoints

| Endpoint | Method | Description |
| --- | --- | --- |
| /cache/stats | GET | Fetch cache hit/miss statistics |
| /cache/init | POST | Initialize Redis cache connection |
| /cache/config | GET | Get current Redis configuration |

## Usage
-----

### Initializing Redis Cache

```javascript
const cacheRedisConfig = require('./cache-redis-config');

cacheRedisConfig.initRedisCache((err) => {
    if (err) {
        console.error(err);
    } else {
        console.log('Redis cache initialized successfully');
    }
});
```

### Fetching Cache Statistics

```javascript
const cacheRedisConfig = require('./cache-redis-config');

cacheRedisConfig.getCacheStats((err, stats) => {
    if (err) {
        console.error(err);
    } else {
        console.log('Cache statistics:', stats);
    }
});
```

## Contributing
------------

Contributions are welcome! Please submit a pull request with your changes and ensure they follow our coding standards.

## License
-------

cache-redis-config is licensed under the MIT License.

## Acknowledgments
--------------

*   [Redis](https://redis.io/)
*   [Node.js](https://nodejs.org/)
*   [ESLint](https://eslint.org/)
*   [Jest](https://jestjs.io/)