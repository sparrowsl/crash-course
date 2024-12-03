<script>
	import TableOfContents from "./toc.md";

	const { children } = $props();
</script>

<section class="grid grid-cols-[max-content_1fr] prose gap-10">
	<aside class="mt-10 h-screen">
		<TableOfContents />
	</aside>

	{@render children?.()}
</section>
