<script lang="ts">
	import '../app.postcss';
	import { LightSwitch } from '@skeletonlabs/skeleton';
	import { SvelteToast } from '@zerodevx/svelte-toast';
</script>

<SvelteToast />

<nav class="z-50 fixed left-6 top-6">
	<LightSwitch />
</nav>
<div class="h-full flex flex-col items-center justify-center gap-6">
	<slot />
</div>
