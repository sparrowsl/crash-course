<script>
import "../app.css";
import "iconify-icon";

// biome-ignore lint/correctness/noUnusedVariables:
// biome-ignore lint/correctness/noUndeclaredVariables:
const { children } = $props();
</script>

<nav class="flex justify-between items-center max-w-6xl mx-auto min-h-20">
  <a href="/" class="block font-bold text-2xl text-accent">Crash Course</a>

  <ul class="flex items-center gap-5 text-sm font-fira-sans">
    <li><a href="/">Home</a></li>
    <li><a href="/languages">Languages</a></li>
    <li>
      <a
        href="https://github.com/sparrowsl/crash-course"
        aria-label="repo"
        target="_blank"
      >
        <iconify-icon icon="simple-icons:github" class="text-2xl">
        </iconify-icon>
      </a>
    </li>
  </ul>
</nav>

<main class="min-h-[100dvh] font-fira-sans max-w-4xl mx-auto">
  {@render children?.()}
</main>

<footer class="text-center font-roboto p-3">
  <small>
    Copyright &copy; 2024 &lurdshar; {new Date().getFullYear()}
  </small>
</footer>
