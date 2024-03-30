const fetch = require('node-fetch');

const API_BASE_URL = "http://localhost:8080";

async function testAPICall() {
    try {
        const response = await fetch(`${API_BASE_URL}/ping`);
        const data = await response.json();
        if (data.Message === 'Pong') {
            console.log('API endpoint test passed');
        } else {
            console.error('API endpoint test failed');
            process.exit(1); // Exit with a non-zero code to indicate test failure
        }
    } catch (error) {
        console.error('Error while testing API endpoint:', error);
        process.exit(1); // Exit with a non-zero code to indicate test failure
    }
}

// Run the test
testAPICall();