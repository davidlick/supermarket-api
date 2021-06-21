import http from 'k6/http';
import { sleep } from 'k6';
import { Rate, Trend } from 'k6/metrics';

const getProduceSuccessRate = new Rate('GetProduceOK');
const getProduceTimingTrend = new Trend('GetProduceTiming');

const addProduceSuccessRate = new Rate('AddProduceOK');
const addProduceTimingTrend = new Trend('AddProduceTiming');

export const options = {
	scenarios: {
		ramp_to_load: {
			executor: 'ramping-arrival-rate',
			startRate: 0,
			preAllocatedVUs: 200,
			maxVUs: 250,
			stages: [
				{ duration: '4m', target: 50, },
			],
		},
		maintain_load: {
			executor: 'constant-arrival-rate',
			rate: 50,
			timeUnit: '1s',
			startTime: '4m',
			duration: '1m',
			preAllocatedVUs: 200,
			maxVUs: 250,
		},
	},
	thresholds: {
		GetProduceOK: ['rate==1'],
		GetProduceTiming: ['p(90)<100','p(95)<200'],

		AddProduceOK: ['rate==1'],
		AddProduceTiming: ['p(90)<100','p(95)<200'],
	},
}

export default function() {
	const id = `${__ITER}-${__VU}`

	const addRes = http.post('http://localhost:3000/v1/produce', JSON.stringify([{
		code: `code-${id}`,
		name: `name-${id}`,
		price: {
			amount: 100,
			currency: "USD"
		}
	}]))

	const getRes = http.get('http://localhost:3000/v1/produce')

	addProduceSuccessRate.add(addRes.status < 300)
	addProduceTimingTrend.add(addRes.timings.duration)

	getProduceSuccessRate.add(getRes.status < 300)
	getProduceTimingTrend.add(getRes.timings.duration)
}
