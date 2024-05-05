<script>
// biome-ignore lint/correctness/noUnusedImports: <explanation>
import Navbar from "$lib/components/Navbar.svelte";
import "../app.css";

// biome-ignore lint/correctness/noUnusedVariables: <explanation>
// biome-ignore lint/correctness/noUndeclaredVariables: <explanation>
const { children } = $props();
</script>

<Navbar/>

<main class="min-h-[100dvh] font-fira-sans max-w-4xl mx-auto">
{@render children?.()}
</main>

<footer class="text-center font-roboto p-3 bg-[#fafafb]">
<small class="block">Copyright &copy; 2024</small>
</footer>
