<script>
// biome-ignore lint/correctness/noUndeclaredVariables: <explanation>
// biome-ignore lint/correctness/noUnusedVariables: <explanation>
const { children } = $props();
</script>

<aside class="mt-10 prose h-screen">
    {@render children?.()}
</aside>
