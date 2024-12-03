<script>
	import Aside from "$lib/components/Aside.svelte";
	import TableOfContents from "./toc.md";

	const { children } = $props();
</script>

<section class="grid grid-cols-[max-content_1fr] gap-10">
	<Aside>
		<TableOfContents />
	</Aside>

	{@render children?.()}
</section>
