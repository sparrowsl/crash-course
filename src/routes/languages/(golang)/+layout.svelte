<script>
// biome-ignore lint/correctness/noUnusedImports:
import Aside from "$lib/components/Aside.svelte";
// biome-ignore lint/correctness/noUnusedImports:
import TableOfContents from "./toc.md";

// biome-ignore lint/correctness/noUnusedVariables:
// biome-ignore lint/correctness/noUndeclaredVariables:
const { children } = $props();
</script>

<section class="grid grid-cols-[max-content_1fr] gap-10">
  <Aside>
    <TableOfContents />
  </Aside>

  {@render children?.()}
</section>
