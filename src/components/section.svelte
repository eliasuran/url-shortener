<script lang="ts">
	import { toast } from '@zerodevx/svelte-toast';
	import { fly } from 'svelte/transition';
	import type { SectionType } from '../types/Sections';

	export let index: number;
	export let sectionData: SectionType;
	export let values: { url: string; path: string; days: number };
	export let incrementStep: () => void;
	export let decrementStep: () => void;
	export let updateValue: (useCase: string, value: string) => void;

	let value: string = values[sectionData.useCase];
</script>

<div
	in:fly={{ duration: 300, y: -500 }}
	out:fly={{ duration: 300, y: 500 }}
	class="absolute inset-0 flex flex-col justify-center items-center gap-6"
>
	{#if index === 0}
		<div class="text-7xl">
			<span class="text-secondary-400">URL</span>-<span class="text-primary-400">SHORTENER</span>
		</div>
	{/if}
	<h2 class="text-2xl">{sectionData.desc}</h2>
	<h2 on:click={decrementStep} class="text-2xl">GO BACK</h2>
	<form
		on:submit={() => {
			if (!value) {
				toast.push('Please fill in the field');
				return;
			}
			if (sectionData.useCase === 'url') {
				try {
					new URL(value);
				} catch (e) {
					toast.push('Please enter a valid URL');
					return;
				}
			}
			if (sectionData.type === 'number' && isNaN(Number(value))) {
				toast.push('Please enter a valid number');
				return;
			}
			incrementStep();
			updateValue(sectionData.useCase, value);
		}}
		class={`${index === 0 ? 'border-secondary-400' : index === 1 ? 'border-tertiary-300' : 'border-primary-400'} border-2 text-xl flex`}
	>
		<input
			class="bg-transparent outline-none w-96 pr-2 p-3"
			type="text"
			placeholder={sectionData.placeholder}
			bind:value
		/>
		<button
			disabled={value ? false : true}
			class={`${value ? 'hover:bg-surface-600 text-gray-300' : 'text-gray-400 cursor-not-allowed'} h-full bg-surface-800 duration-300 p-3`}
			>Next</button
		>
	</form>
</div>
