<script lang="ts">
	import Section from '../components/section.svelte';
	import type { SectionType } from '../types/Sections';

	const sections: SectionType[] = [
		{
			desc: 'Enter the url you would like to shorten',
			placeholder: 'URL...',
			useCase: 'url'
		},
		{
			desc: 'What do you want the path to be?',
			placeholder: 'Path...',
			useCase: 'path'
		},
		{
			desc: 'How long would you like the url to last?',
			placeholder: 'Time in days...',
			useCase: 'days'
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
		<Section
			index={i}
			desc={section.desc}
			placeholder={section.placeholder}
			useCase={section.useCase}
			{incrementStep}
			{updateValue}
		/>
	{/if}
{/each}
