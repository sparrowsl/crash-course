<script>
// biome-ignore lint/correctness/noUnusedImports: <explanation>
import TableOfContents from "./toc.md";

// biome-ignore lint/correctness/noUnusedVariables: <explanation>
// biome-ignore lint/correctness/noUndeclaredVariables: <explanation>
const { children } = $props();
</script>

<section class="grid grid-cols-[max-content_1fr] gap-10">
  <aside class="mt-10 prose h-screen">
    <TableOfContents/>
  </aside>

{@render children?.()}
</section>
