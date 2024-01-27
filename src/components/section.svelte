<script lang="ts">
	import { fly } from 'svelte/transition';

	export let title: string = '';
	export let desc: string;
	export let placeholder: string;
	export let incrementStep: () => void;
	export let useCase: string;
	export let updateValue: (useCase: string, value: string) => void;

	let inputValue: string = '';
</script>

<div
	in:fly={{ duration: 300, y: -500 }}
	out:fly={{ duration: 300, y: 500 }}
	class="absolute inset-0 flex flex-col justify-center items-center gap-6"
>
	<h1 class="text-7xl">{title}</h1>
	<h2 class="text-2xl">{desc}</h2>
	<form action="" class="border-primary-400 border-2 p-3 text-xl flex">
		<input
			class="bg-transparent outline-none w-96 pr-2"
			type="text"
			{placeholder}
			bind:value={inputValue}
		/>
		<button
			on:click={() => {
				incrementStep();
				updateValue(useCase, inputValue);
			}}
			class="h-full hover:underline">Next</button
		>
	</form>
</div>
