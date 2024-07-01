<head>
    <title>README</title>
</head>

<body>
    <h1>Weather API Project</h1>
    <p>This project is a simple web server written in Go that fetches and serves weather data.</p>
    <h2>Project Structure</h2>
    <ul>
        <li><code>main.go</code>: This is the main file of the project. It sets up the HTTP server and the route handlers.</li>
        <li><code>.env</code>: This file contains environment variables, such as the API key for the weather service.</li>
    </ul>
    <h2>How to Run</h2>
    <ol>
        <li>Ensure that Go is installed on your machine.</li>
        <li>Clone the repository and navigate to the project directory.</li>
        <li>Run <code>go run main.go</code> in the terminal to start the server.</li>
        <li>Navigate to <code>http://localhost:8080/weather/{city}</code> in your web browser to fetch weather data for a specific city.</li>
    </ol>
    <h2>Dependencies</h2>
    <p>This project uses the <code>github.com/joho/godotenv</code> package to load environment variables from the <code>.env</code> file.</p>
</body>
