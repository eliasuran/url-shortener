<script lang="ts">
	import Section from '../components/section.svelte';
	import type { SectionType } from '../types/Sections';

	const sections: SectionType[] = [
		{
			desc: 'Enter the url you would like to shorten',
			placeholder: 'URL...',
			useCase: 'url',
			type: 'text'
		},
		{
			desc: 'What do you want the path to be?',
			placeholder: 'Path...',
			useCase: 'path',
			type: 'text'
		},
		{
			desc: 'How long would you like the url to last?',
			placeholder: 'Time in days...',
			useCase: 'days',
			type: 'number'
		}
	];

	let step: number = 1;
	function incrementStep() {
		step += 1;
	}

	import type { Values } from '../types/Values';
	let values: Values = {
		url: '',
		path: '',
		days: 0
	};
	function updateValue(useCase: string, value: string) {
		values[useCase] = value;
		console.log(values);
	}

	import { newURL } from '$lib/backend';
	$: if (step === 4) {
		newURL(values);
	}
</script>

{#each sections as section, i}
	{#if i + 1 === step}
		<Section index={i} sectionData={section} {incrementStep} {updateValue} />
	{/if}
{/each}
