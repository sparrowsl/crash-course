import { sveltekit } from "@sveltejs/kit/vite";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";
import { NodeGlobalsPolyfillPlugin } from "@esbuild-plugins/node-globals-polyfill";

export default defineConfig({
	assetsInclude: ["**/*.md"],
	plugins: [tailwindcss(), sveltekit()],
	optimizeDeps: {
		esbuildOptions: {
			define: {
				global: "globalThis",
			},
			plugins: [NodeGlobalsPolyfillPlugin({ buffer: true })],
		},
	},
});
