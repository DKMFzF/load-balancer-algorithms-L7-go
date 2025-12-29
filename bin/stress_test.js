import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 20 },
        { duration: '1m', target: 100 },
        { duration: '30s', target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(95)<500'],
    },
};

export default function () {
    http.get('http://localhost:8080/');
    sleep(1);
}
