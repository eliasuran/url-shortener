import type { Values } from '../types/Values';

export interface SlugLoad {
	url: string;
}

export const apiURL = 'http://localhost:8080';

export async function newURL(values: Values) {
	const res = await fetch(`${apiURL}/newURL`, {
		method: 'POST',
		body: JSON.stringify(values)
	});

	if (res.ok) {
		return 'Added';
	}

	return 'Error';
}
