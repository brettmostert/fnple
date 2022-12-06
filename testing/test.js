import http from 'k6/http';
import { check, sleep } from 'k6';
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export const options = {
    stages: [
      { duration: '5s', target: 500 },
      { duration: '45s', target: 1000 },
      { duration: '1m30s', target: 2000 },
      { duration: '20s', target: 0 },
    ],
  };

  export default function () {
    const params = {
      headers: {
        'Content-Type': 'application/json',
      },
    };

    let data = {
      "CoalationId": uuidv4(),
      "Description" : "test",
      "Entries" : [
          {
              "Amount": 10,
              "AccountId": 1,
              "Type":"cr",
              "Description": "test"
          },
          {
              "Amount": 10,
              "AccountId": 2,
              "Type":"dr",
              "Description": "test"
          }
      ]
  }

    const res = http.post('http://127.0.0.1:3000/transactions', JSON.stringify(data), params);

  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
  }
