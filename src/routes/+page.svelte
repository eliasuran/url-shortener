<script lang="ts">
	import Section from '../components/section.svelte';
	import { sections } from '$lib/sections';

	let step: number = 1;
	function incrementStep() {
		step += 1;
	}
	function decrementStep() {
		step -= 1;
	}

	let values: { url: string; path: string; days: number } = {
		url: '',
		path: '',
		days: 0
	};
	function updateValue(useCase: string, value: string) {
		values[useCase] = value;
		console.log(values);
	}

	import { newURL } from '$lib/backend';
	$: if (step > sections.length) {
		newURL(values);
	}
</script>

{#each sections as section, i}
	{#if i + 1 === step}
		<Section
			index={i}
			sectionData={section}
			{values}
			{incrementStep}
			{decrementStep}
			{updateValue}
		/>
	{/if}
{/each}
