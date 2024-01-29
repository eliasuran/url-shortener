<script lang="ts">
	import type { SlugLoad } from '$lib/backend';
	import { onMount } from 'svelte';
	export let data: SlugLoad;

	onMount(() => {
		if (data.url) {
			window.location.href = data.url;
		}
	});
</script>

{#if data.url}
	<h1 class="text-7xl">Redirecting...</h1>
{:else}
	<h1 class="text-7xl">404</h1>
{/if}
