import type { PageServerLoad } from './$types';
import { apiURL } from '$lib/backend';

export const load: PageServerLoad = async ({ params }) => {
	const { slug } = params;
	const res = await fetch(`${apiURL}/${slug}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json'
		}
	});

	if (res.ok) {
		const data = await res.json();
		return {
			url: data.url
		};
	}

	return { status: res.status };
};
