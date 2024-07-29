<script>
// biome-ignore lint/correctness/noUnusedImports: <explanation>
import Aside from "$lib/components/Aside.svelte";

// biome-ignore lint/correctness/noUnusedVariables: <explanation>
// biome-ignore lint/correctness/noUndeclaredVariables: <explanation>
const { children } = $props();
</script>

<section class="grid grid-cols-[max-content_1fr] gap-10">
  <Aside>
    <TableOfContents/>
  </Aside>

{@render children?.()}
</section>
