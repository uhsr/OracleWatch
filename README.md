# OracleWatch: Proactive Oracle Database Monitoring

OracleWatch is a robust and lightweight monitoring tool written in Go, designed to provide real-time insights into the performance and health of your Oracle database instances. It allows developers and database administrators to proactively identify and address potential issues before they impact critical applications. This tool focuses on collecting vital database metrics and exposing them through a simple and configurable interface.

OracleWatch fills a crucial gap by offering an easily deployable and highly customizable monitoring solution. Unlike monolithic monitoring suites that can be resource-intensive and complex to configure, OracleWatch is designed for efficiency and simplicity. It leverages Go's concurrency capabilities to efficiently collect data from multiple Oracle instances without significant overhead. The collected metrics can be easily integrated with existing monitoring systems, such as Prometheus, Grafana, or similar time-series databases, allowing users to build comprehensive dashboards and alerting rules.

The primary goal of OracleWatch is to provide actionable intelligence. By monitoring key database metrics, users can gain a deep understanding of database performance trends, identify bottlenecks, and optimize resource allocation. This proactive approach minimizes downtime, improves application performance, and reduces the risk of critical database failures. OracleWatch promotes a data-driven approach to database administration, enabling informed decision-making and continuous improvement.

## Key Features

*   **SQL*Net Connection Pooling:** OracleWatch utilizes robust SQL*Net connection pooling to minimize connection overhead and efficiently manage database connections. This ensures consistent performance even under heavy load. The connection pool size is configurable to optimize resource utilization based on the number of monitored instances.
*   **Configurable Metrics:** Users can define specific SQL queries to retrieve custom metrics tailored to their individual monitoring needs. This allows for maximum flexibility and enables monitoring of application-specific performance indicators. The query results are automatically converted into numerical values for easy integration with monitoring systems.
*   **Multiple Instance Support:** OracleWatch supports monitoring multiple Oracle database instances simultaneously. The configuration file specifies the connection details for each instance, allowing for centralized monitoring of an entire database environment.
*   **REST API Endpoint:** A simple and well-defined REST API endpoint exposes the collected metrics in JSON format. This allows for easy integration with monitoring dashboards and alerting systems. The API endpoint can be secured using basic authentication or other security mechanisms.
*   **Prometheus Integration:** OracleWatch can be configured to expose metrics in Prometheus format, allowing seamless integration with the popular Prometheus monitoring system. This simplifies the process of building dashboards and setting up alerts based on database performance metrics.
*   **Minimal Dependencies:** OracleWatch is built with minimal external dependencies, ensuring easy deployment and reducing the risk of dependency conflicts. The core functionality relies primarily on the standard Go library and the go-oci8 driver.
*   **Asynchronous Data Collection:** Data collection from each Oracle instance is performed asynchronously using Go routines. This ensures that monitoring performance does not degrade even when monitoring multiple instances with varying response times.

## Technology Stack

*   **Go (Golang):** The core programming language used to build OracleWatch. Go's concurrency capabilities and performance make it ideal for building efficient and scalable monitoring tools.
*   **go-oci8:** The primary Go driver for connecting to Oracle databases. This driver provides a robust and reliable interface for executing SQL queries and retrieving data from Oracle instances.
*   **net/http:** The standard Go library package used for building the REST API endpoint and handling HTTP requests.
*   **encoding/json:** The standard Go library package used for encoding and decoding JSON data for the REST API.
*   **github.com/joho/godotenv:** Used for loading environment variables from a `.env` file.

## Installation

1.  Ensure you have Go installed and configured on your system (version 1.18 or later).
2.  Install the go-oci8 driver. This typically requires installing the Oracle Instant Client SDK. Refer to the go-oci8 documentation for specific installation instructions based on your operating system. export ORACLE_HOME=/path/to/your/oracle/instantclient
3.  Clone the OracleWatch repository from GitHub: git clone https://github.com/uhsr/OracleWatch.git
4.  Navigate to the OracleWatch directory: cd OracleWatch
5.  Build the OracleWatch executable: go build -o oraclewatch main.go

## Configuration

OracleWatch is configured using environment variables. You can define these variables directly in your shell environment or in a `.env` file in the OracleWatch directory.

*   `ORACLE_CONNECTION_STRING`: The Oracle connection string in the format `username/password@host:port/service_name`. Example: `system/password@localhost:1521/ORCL`. You can specify multiple connection strings separated by commas. Each connection string corresponds to a different Oracle instance.
*   `PORT`: The port on which the REST API endpoint will listen. Defaults to 8080.
*   `METRICS_QUERY`: The SQL query used to retrieve metrics from the Oracle database. The query must return a single row with numerical values for each metric. Example: `SELECT value1, value2 FROM my_metrics_table`.
*   `PROMETHEUS_ENABLED`: Set to `true` to enable Prometheus endpoint at /metrics. Defaults to `false`.

Example `.env` file:

ORACLE_CONNECTION_STRING=system/password@localhost:1521/ORCL
PORT=8080
METRICS_QUERY=SELECT SUM(bytes) FROM dba_segments
PROMETHEUS_ENABLED=true

## Usage

1.  Run the OracleWatch executable: ./oraclewatch
2.  Access the REST API endpoint to retrieve metrics in JSON format: `curl http://localhost:8080/metrics` or `curl http://localhost:8080/prometheus` if prometheus is enabled.
3.  The JSON response will contain the metrics collected from the Oracle database(s).

Example JSON response:

{
"value1": 12345,
"value2": 67890
}

If Prometheus is enabled, scrape the /metrics endpoint with your Prometheus instance. The exposed metrics will be in Prometheus exposition format.

## Contributing

We welcome contributions to OracleWatch! Please follow these guidelines:

1.  Fork the repository on GitHub.
2.  Create a new branch for your feature or bug fix.
3.  Implement your changes and write tests to ensure they are working correctly.
4.  Submit a pull request to the main branch.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/OracleWatch/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the developers of the go-oci8 driver for their excellent work in providing a reliable and efficient Oracle database driver for Go.