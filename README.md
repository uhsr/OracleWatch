# OracleWatch - Decentralized Price Oracle Monitoring and Alerting

OracleWatch provides a decentralized and automated solution for monitoring price data from Chainlink oracles and triggering alerts based on user-defined thresholds. This serverless application leverages Chainlink Keepers and custom smart contracts to ensure reliable and tamper-proof price monitoring. Alerts are delivered via WebSockets, enabling real-time notifications for critical price movements. The system is designed for developers and DeFi users who require immediate notification of significant price deviations, preventing potential losses and enabling proactive decision-making.

OracleWatch addresses the critical need for timely and reliable price data in the DeFi space. By using a decentralized architecture based on Chainlink oracles, the system eliminates single points of failure. The integration with Chainlink Keepers automates the monitoring process, removing the need for manual intervention. The threshold-based alerting system allows users to define specific price ranges and receive notifications only when these thresholds are breached, minimizing alert fatigue. The WebSockets delivery mechanism ensures that alerts are received in real-time, enabling rapid response to market changes.

This project provides a comprehensive framework for decentralized price monitoring. It includes a Go-based backend for handling price data and alert logic, custom Solidity smart contracts for defining price thresholds and interacting with Chainlink Keepers, and a WebSocket server for delivering real-time alerts. The serverless architecture makes it easy to deploy and scale the system without managing infrastructure. OracleWatch empowers developers to build robust and resilient DeFi applications that can react quickly to price fluctuations, mitigating risk and maximizing opportunities.

## Key Features

*   **Decentralized Price Monitoring:** Leverages Chainlink oracles for tamper-proof and reliable price data feeds. The system supports multiple Chainlink price feeds, allowing users to monitor a diverse range of assets.
*   **Chainlink Keeper Integration:** Uses Chainlink Keepers to automate the process of checking price thresholds. A custom smart contract is responsible for checking the price and triggering the Keeper to execute the `performUpkeep` function.
*   **Customizable Price Thresholds:** Users can define specific price thresholds (upper and lower bounds) for each asset being monitored. These thresholds are stored on-chain in a smart contract, ensuring transparency and immutability.
*   **Threshold-Triggered Alerts:** Alerts are triggered automatically when the price of an asset crosses a predefined threshold. The system checks the price against the defined thresholds on a regular basis.
*   **Real-time WebSocket Alerts:** Alerts are delivered in real-time via WebSockets, enabling immediate notification of critical price movements. Clients can subscribe to specific assets or price feeds.
*   **Serverless Architecture:** The backend is designed to be deployed on a serverless platform, eliminating the need for infrastructure management. This allows for easy scaling and reduced operational costs.
*   **Configurable Alerting Frequency:** The frequency at which the system checks price thresholds can be configured to balance accuracy and cost. This allows users to optimize the system for their specific needs.

## Technology Stack

*   **Go:** The primary programming language for the backend logic, including price data processing, alert generation, and WebSocket server.
*   **Solidity:** Used for writing the smart contracts that define price thresholds and interact with Chainlink Keepers.
*   **Chainlink:** Provides the decentralized price oracles and Keepers infrastructure.
*   **WebSockets:** Used for real-time delivery of alerts to clients. A Go library is used to implement the WebSocket server.
*   **Ethereum (or compatible EVM chain):** The blockchain where the smart contracts are deployed and executed.
*   **[Optional] Serverless Platform (e.g., AWS Lambda, Google Cloud Functions):** Used for deploying the backend in a serverless environment.

## Installation

1.  **Install Go:** Ensure you have Go installed and configured correctly. Verify by running `go version`.
2.  **Clone the Repository:** `git clone https://github.com/uhsr/OracleWatch.git`
3.  **Navigate to the Project Directory:** `cd OracleWatch`
4.  **Install Dependencies:** `go mod download && go mod tidy`
5.  **Install Hardhat:** `npm install -g hardhat`
6.  **Compile Smart Contracts:** Navigate to the `contracts` directory and run `hardhat compile`.
7.  **Deploy Smart Contracts:** Deploy the smart contracts to a test network or the mainnet. You will need to configure Hardhat with your Ethereum provider and private key. You can use a command like `npx hardhat run scripts/deploy.js --network <network-name>`. Save the contract address.

## Configuration

The application is configured using environment variables. Create a `.env` file in the root directory with the following variables:



*   `CHAINLINK_NODE_URL`: URL of the Chainlink node.
*   `CHAINLINK_CONTRACT_ADDRESS`: Address of the Chainlink contract.
*   `KEEPER_REGISTRY_ADDRESS`: Address of the Keeper registry.
*   `UPKEEP_ID`: ID of the upkeep registered on the Keeper registry.
*   `WEBSOCKET_PORT`: Port on which the WebSocket server will listen.
*   `PRICE_THRESHOLD`: The price difference threshold that triggers an alert (e.g., 0.05 for 5%).
*   `PRIVATE_KEY`: The private key of the account used to interact with the smart contracts.

## Usage

1.  **Run the Backend:** After configuring the environment variables, run the backend using the command `go run main.go`. This will start the WebSocket server and the price monitoring process.
2.  **Connect to the WebSocket Server:** Use a WebSocket client to connect to the WebSocket server running on the configured port.
3.  **Subscribe to Price Alerts:** Once connected, send a subscription message to the server to receive price alerts for specific assets. For example: `{"type": "subscribe", "asset": "ETH/USD"}`.
4.  **Receive Alerts:** The server will send alerts in JSON format when the price of the subscribed asset crosses the defined threshold. For example: `{"asset": "ETH/USD", "price": 3000, "threshold": 2900, "alert": "Price below threshold"}`.

## Contributing

We welcome contributions to OracleWatch! Please follow these guidelines:

*   Fork the repository and create a branch for your feature or bug fix.
*   Write clear and concise commit messages.
*   Submit a pull request with a detailed description of your changes.
*   Ensure your code adheres to the project's coding style.
*   Include unit tests for any new functionality.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/OracleWatch/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Chainlink team for providing the infrastructure and resources necessary to build OracleWatch.