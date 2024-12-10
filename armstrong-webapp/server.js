const http = require('http');

const server = http.createServer((req, res) => {
    // Allow CORS by adding headers to every response
    res.setHeader('Access-Control-Allow-Origin', 'http://localhost:3001'); // Your React app's URL
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, DELETE');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type, Authorization');

    // Handle preflight requests (OPTIONS method)
    if (req.method === 'OPTIONS') {
        res.writeHead(204); // No Content
        res.end();
        return;
    }

    // Basic route handling (example)
    if (req.url === '/register' && req.method === 'POST') {
        let body = '';

        // Collect request data
        req.on('data', chunk => {
            body += chunk.toString();
        });

        req.on('end', () => {
            console.log('Received data:', body);

            // Respond to the client
            res.writeHead(200, { 'Content-Type': 'application/json' });
            res.end(JSON.stringify({ message: 'Data received', data: JSON.parse(body) }));
        });
    } else {
        // Handle 404 for other routes
        res.writeHead(404, { 'Content-Type': 'text/plain' });
        res.end('Not Found');
    }
});

server.listen(8080, () => {
    console.log('Server running on http://localhost:8080');
});


