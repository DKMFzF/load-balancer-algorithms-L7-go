import http from 'k6/http';
import { sleep, check } from 'k6';
import { randomIntBetween, randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
  scenarios: {
    realistic_load: {
      executor: 'ramping-arrival-rate',
      startRate: 10,
      timeUnit: '1s',
      preAllocatedVUs: 20,
      maxVUs: 300,
      stages: [
        { target: 50, duration: '1m' },
        { target: 100, duration: '1m' },
        
        { target: 250, duration: '1m' },
        { target: 300, duration: '1m' },
        { target: 280, duration: '1m' },
        
        { target: 200, duration: '1m' },
        { target: 180, duration: '1m' },
        { target: 160, duration: '1m' },
        
        { target: 220, duration: '1m' },
        { target: 240, duration: '1m' },
        { target: 260, duration: '1m' },
        
        { target: 200, duration: '1m' },
        { target: 150, duration: '1m' },
        { target: 100, duration: '1m' },
        
        { target: 50, duration: '1m' },
        { target: 20, duration: '1m' },
      ],
    },
  },
  
  thresholds: {
    http_req_duration: ['p(95)<500', 'p(99)<1000'],
    http_req_failed: ['rate<0.01'],
    http_reqs: ['count>50000'],
  },
};

export default function () {
  const userTypes = [
    { thinkTime: 0.5, errorChance: 0.05 },
    { thinkTime: 1.5, errorChance: 0.02 },
    { thinkTime: 3.0, errorChance: 0.01 },
  ];
  
  const userType = randomItem(userTypes);
  
  const thinkTime = userType.thinkTime * randomIntBetween(0.8, 1.2);
  
  const response = http.get('http://localhost:8080/');
  
  check(response, {
    'status is 200': (r) => r.status === 200,
    'response time OK': (r) => r.timings.duration < 1000,
  });
  
  sleep(thinkTime);
  
  if (Math.random() < 0.3) {
    const secondRequest = http.get('http://localhost:8080/');
    sleep(thinkTime * 0.5);
  }
}
