/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/**/*.{html,js,ts,svelte}",
  ],
  theme: {
    extend: {
      fontFamily: {
        roboto: ["Roboto", "sans-serif"],
        "fira-sans": ["Fira Sans", "sans-serif"]
      },
      colors: {
        accent: "#FAA300",
        primary: "#FFC700"
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
