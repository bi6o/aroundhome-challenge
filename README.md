# Partner Service API

This service is designed to match customers with partners based on specific criteria such as floor material and location.

## Documentation

For a detailed API specification, refer to the Swagger documentation provided at `/docs/swagger.yaml` or use Redoc to view the documentation by running:

```sh
npx @redocly/cli preview-docs docs/swagger.yaml
```

Make sure to install dependencies with `make install` before running the redoc tool, to ensure that the CLI is installed locally.

## Requirements

- Go (for service development and execution)
- PostgreSQL (for database storage)
- Make (for managing commands like migrations and dependencies installation)

## Using the Postgres earth_distance Extension

### Why earth_distance?

In the partner matching service, accurately calculating the geographical distance between a customer's location and a partner's office is crucial for determining if a partner is within the operating radius for a given customer request. To achieve this, we've utilized the PostgreSQL earth_distance extension.

The `earth_distance` extension offers functions to compute distances on the Earth's surface, assuming the Earth is a perfect sphere. It simplifies the process of calculating the "beeline" or great-circle distance between two points given their latitudes and longitudes. This method is both efficient and sufficiently accurate for our application's purposes, where exact precision to the meter is not critical, but a reliable approximation of distance is necessary for filtering and sorting partners.

### Overview of the earth_distance Extension

The `earth_distance` extension builds upon the `cube` extension in PostgreSQL, which adds support for multidimensional cubes or arrays. `earth_distance` uses this functionality to represent points on the Earth's surface for distance calculations. Key functions include:

**ll_to_earth(lat, long)**: Converts latitude and longitude values into a point on the Earth's surface, represented as a `cube`.

**earth_distance(ll_to_earth(lat1, long1), ll_to_earth(lat2, long2))**: Calculates the great-circle distance between two points specified by their latitudes and longitudes.

For more detailed information, [please refer to the official documentation of the extension](https://www.postgresql.org/docs/current/earthdistance.html)

## Getting Started

### Install Dependencies and Tools

Before running the service, ensure you have the above requirements installed on your system. Then, install the necessary Go dependencies and tools needed to run the service and manage the database:

```sh
make install
```

This command installs required Go packages, including `goose` for database migrations and `@redocly/cli` for viewing the API documentation.

### Prepare the Database

Run database migrations to set up the necessary database schema before starting the service:

```sh
make migrate
```

### Seeding the Database

Execute the following SQL query to seed your database with partner data:

```sql
INSERT INTO partners (flooring_materials, address_lat, address_long, operating_radius, rating)
VALUES
('{wood, tile}', 40.7128, -74.0060, 50, 5),
('{carpet}', 34.0522, -118.2437, 30, 4);
```

Adjust the values as needed for your testing purposes.

### Running the Service

Before starting the service, copy `.env.example` file and rename the copy to `.env`. Ensure you have set up the necessary environment variables or configuration files needed by `main.go` to specify the service's port, database connection details, etc.

To start the service, run the following command:

```sh
go run main.go
```

### Making Requests

#### Match Endpoint

To find matching partners for a customer, send a POST request to `/partners/match` with the following example body:

```json
{
  "floor_material": "wood",
  "address_long": -74.006,
  "address_lat": 40.7128,
  "floor_area": 120.5,
  "phone_number": "+1234567890"
}
```

This request matches partners experienced with wood flooring within the specified location radius.

#### Get Partner Endpoint

To retrieve details of a specific partner, send a GET request to `/partners/{id}`, where `{id}` is the UUID of the partner you wish to retrieve.

Example request:

```sh
curl http://localhost:8080/partners/uuid-here
```

Replace `uuid-here` with the actual UUID of the partner.
